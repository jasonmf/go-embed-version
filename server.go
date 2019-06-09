package versserv

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AgentZombie/go-embed-version/cmd"
)

type Server struct {
	// server components
}

func NewServer() (*Server, error) {
	// make your own Server, Mux
	s := &Server{}
	http.HandleFunc("/", s.root)
	return s, nil
}

func (s *Server) ListenAndServe() error {
	log.Print("starting server version ", cmd.Version)
	return http.ListenAndServe("localhost:8000", nil)
}

func (s *Server) root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Version", cmd.Version)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<html><body><h1>Hello World!</h1></body></html>")
}
