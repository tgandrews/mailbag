package incominghandler

import "net/http"
import "log"
import "github.com/gorilla/mux"

type IncomingHandler struct {
	router *mux.Router
}

func (h IncomingHandler) Run() {
	log.Print("Starting server on port 8080")
	h.router = h.createRouter()
	log.Fatal(http.ListenAndServe(":8080", h.router))
}

func (h IncomingHandler) createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.FormPostHandler).Methods("POST")
	router.HandleFunc("/admin", h.AdminGetHandler).Methods("GET")
	return router
}

func (h IncomingHandler) FormPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("User Agent: %s | URI: %s", r.UserAgent(), r.RequestURI)
}

func (h IncomingHandler) AdminGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world")
}
