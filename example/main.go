package main

import (
	"log"
	"net/http"

	"github.com/kamaln7/tiel"
)

func main() {
	t, err := tiel.New(nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	t.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("chirp chirp"))
	})

	err = t.Sing()
	if err != nil {
		log.Fatalln(err)
	}
}
