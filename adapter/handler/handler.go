package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Pranc1ngPegasus/middlechain"
	"github.com/Pranc1ngPegasus/pubsub-push/adapter/handler/middleware"
	"github.com/Pranc1ngPegasus/pubsub-push/domain/logger"
	"github.com/google/wire"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	logger logger.Logger
	mux    http.Handler
}

func NewHandler(
	logger logger.Logger,
) *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		logger: logger,
		mux:    mux,
	}
	h.mux = middlechain.Chain(h.mux,
		middleware.Logger(logger),
	)

	mux.HandleFunc("/ping", h.ping)
	mux.HandleFunc("/pubsub", h.pubsub)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type PubSubRequest struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func (h *Handler) pubsub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(ctx, "ioutil.ReadAll", h.logger.Field("err", err))
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	var m PubSubRequest
	if err := json.Unmarshal(body, &m); err != nil {
		h.logger.Error(ctx, "json.Unmarshal", h.logger.Field("err", err))
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	data := bytes.NewBuffer(m.Message.Data).String()

	h.logger.Info(ctx, "request",
		h.logger.Field("id", m.Message.ID),
		h.logger.Field("body", data),
	)

	w.WriteHeader(http.StatusOK)
}
