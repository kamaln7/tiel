package tiel

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Tiel is an instance of the Tiel HTTP API
type Tiel struct {
	Server *http.Server
	Mux    *http.ServeMux
	Router *mux.Router

	song *Song
}

// New creates a new Tiel instance
func New(song *Song, server *http.Server) (*Tiel, error) {
	// take care of an empty song val
	if song == nil {
		song = &Song{}
	}

	t := &Tiel{
		Mux:    http.NewServeMux(),
		Router: mux.NewRouter(),

		song: song,
	}

	// take care of an empty server val
	if server == nil {
		server = &http.Server{}
	}

	// set server options
	server.Addr = song.UsableListenAddress()
	server.Handler = t.Mux

	t.Server = server
	t.Mux.Handle("/", t.Router)

	return t, nil
}

// Sing starts the HTTP Server
func (t *Tiel) Sing() error {
	log.Printf("tiel ⌁ singing on %s\n\n", t.Server.Addr)

	return t.Server.ListenAndServe()
}

// SingTLS starts the HTTP Server with TLS
func (t *Tiel) SingTLS(certFile, keyFile string) error {
	log.Printf("tiel ⌁ singing on %s\n\n", t.Server.Addr)

	return t.Server.ListenAndServeTLS(certFile, keyFile)
}
