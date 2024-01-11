package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"server-1/pkg/messages"
)

func main() {
	file, _ := os.Create("DataBase/test.json")
	file.Close()

	store := messages.NewFileStore("DataBase/test.json")
	messageHandler := NewMessageHandler(store)
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/message", messageHandler)
	mux.Handle("/message/", messageHandler)
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type messageStore interface {
	Add(message messages.Message) error
	List() ([]messages.Message, error)
	Remove(id string) error
}
type messageHandler struct {
	store messageStore
}

func NewMessageHandler(s messageStore) *messageHandler {
	return &messageHandler{
		store: s,
	}
}

var (
	MessageRe       = regexp.MustCompile(`^/message/*$`)
	MessageReWithId = regexp.MustCompile(`^/message/([0-9]+(?:[0-9]+)+)&`)
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func (h *messageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage messages.Message
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		fmt.Println(err)
		InternalServerError(w, r)
		return
	}
	err = h.store.Add(newMessage)
	if err != nil {
		fmt.Println(err)
		InternalServerError(w, r)
		return
	}

}
func (h *messageHandler) ListMessage(w http.ResponseWriter, r *http.Request) {
	messages, err := h.store.List()
	if err != nil {
		fmt.Println(err)
		InternalServerError(w, r)
		return
	}
	text, _ := json.MarshalIndent(messages, "", "	")
	w.Write(text)
}
func (h *messageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {}

func (h *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && MessageRe.MatchString(r.URL.Path):
		h.CreateMessage(w, r)
		return
	case r.Method == http.MethodGet && MessageRe.MatchString(r.URL.Path):
		h.ListMessage(w, r)
		return
	case r.Method == http.MethodDelete && MessageReWithId.MatchString(r.URL.Path):
		h.DeleteMessage(w, r)
		return
	default:
		return

	}
}
