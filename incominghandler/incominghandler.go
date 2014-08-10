package incominghandler

import "net/http"
import "log"
import "fmt"

type IncomingHandler struct{}

func (h IncomingHandler) Run() {
	log.Fatal(http.ListenAndServe(":8080", h))
}

func (h IncomingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.UserAgent())
}
