package tiel

import (
	"fmt"
	"os"
)

// Song contains the config for a Tiel instance
type Song struct {
	ListenAddress string
}

// UsableListenAddress returns the configured listen address or tries to guess one based on environment variables and set defaults
func (s Song) UsableListenAddress() string {
	if s.ListenAddress != "" {
		return s.ListenAddress
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf("%s:%s", host, port)
}
