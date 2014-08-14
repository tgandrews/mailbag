package incominghandler

import "net/http"
import "log"

type IncomingHandler struct{}

func (h IncomingHandler) Run() {
	log.Print("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func (h IncomingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("User Agent: %s | URI: %s", r.UserAgent(), r.RequestURI)
}
