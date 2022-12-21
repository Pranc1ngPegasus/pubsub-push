package configuration

type Configuration interface {
	Common() *Common
	GCP() *GCP
	Server() *Server
}

type (
	Config struct {
		Common Common
		GCP    GCP
		Server Server
	}

	Common struct {
		Debug bool `env:"DEBUG"`
	}

	GCP struct {
		GCPProjectID string `env:"GCP_PROJECT_ID"`
		Topic        string `env:"GCP_PUBSUB_TOPIC"`
	}

	Server struct {
		Port string `env:"PORT"`
	}
)
