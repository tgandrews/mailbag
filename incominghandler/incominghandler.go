package incominghandler

import "net/http"
import "log"
import "time"
import "github.com/gorilla/mux"
import s "github.com/tgandrews/mailbag/subscriber"

// An IncomingHandler handles the incoming HTTP requests and performs appropriate
// actions. The top leve controller if you will.
type IncomingHandler struct {
	router *mux.Router
}

// Run starts the server
func (h IncomingHandler) Run() {
	log.Print("Starting server on port 8080")
	h.router = h.createRouter()
	log.Fatal(http.ListenAndServe(":8080", h.router))
}

func (h IncomingHandler) createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.formPostHandler).Methods("POST")
	router.HandleFunc("/admin", h.adminGetHandler).Methods("GET")
	return router
}

func (h IncomingHandler) formPostHandler(w http.ResponseWriter, r *http.Request) {
	subscriber := s.Subscriber{}
	subscriber.SubscribeTime = time.Now()
	subscriber.EmailAddress = r.PostFormValue("email")
	subscriber.Referer = r.Referer()
	subscriber.IPAddress = r.RemoteAddr
	subscriber.UserAgent = r.UserAgent()

	log.Printf("URI: %s", r.RequestURI)
	log.Printf("Subscriber: %v", subscriber)
}

func (h IncomingHandler) adminGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world")
}
