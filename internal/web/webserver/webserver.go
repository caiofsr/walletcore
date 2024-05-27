package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	WebServerPort string
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		WebServerPort: webServerPort,
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)

	for path, handler := range ws.Handlers {
		ws.Router.Post(path, handler)
	}

	http.ListenAndServe(ws.WebServerPort, ws.Router)
}
