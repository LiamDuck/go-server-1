package main

import (
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/message", &messageHandler{})
	mux.Handle("/message/", &messageHandler{})
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type messageHandler struct{}

var (
	MessageRe       = regexp.MustCompile(`^/message/*$`)
	MessageReWithId = regexp.MustCompile(`^/message/([0-9]+(?:[0-9]+)+)&`)
)

func (h *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of messages"))
}

func (h *messageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {}
func (h *messageHandler) ListMessage(w http.ResponseWriter, r *http.Request)   {}
func (h *messageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {}
