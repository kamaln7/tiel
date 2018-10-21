package tiel

import "net/http"

// Any sets a route handle that handles all HTTP verbs
func (t *Tiel) Any(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler)
}

// Get sets a route handle that handles GET requests only
func (t *Tiel) Get(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler).Methods("GET")
}

// Post sets a route handle that handles POST requests only
func (t *Tiel) Post(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler).Methods("POST")
}

// Put sets a route handle that handles PUT requests only
func (t *Tiel) Put(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler).Methods("PUT")
}

// Patch sets a route handle that handles PATCH requests only
func (t *Tiel) Patch(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler).Methods("PATCH")
}

// Delete sets a route handle that handles DELETE requests only
func (t *Tiel) Delete(path string, handler http.HandlerFunc) {
	t.Router.HandleFunc(path, handler).Methods("DELETE")
}
