package monolith

import (
	"time"
)

type AppConfig struct {
	Environment     string
	LogLevel        string
	PGUrl           string
	RpcAddress      string
	HttpAddress     string
	ShutdownTimeout time.Duration
}
