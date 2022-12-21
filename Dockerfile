FROM golang:1.19 as builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -tags netgo -ldflags '-s -w -extldflags "-static"' -o app ./cmd/subscribe

FROM gcr.io/distroless/static-debian11:nonroot as runner

WORKDIR /

COPY --from=builder /go/src/app/app /app

ENTRYPOINT ["/app"]
