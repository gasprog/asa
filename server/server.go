//Package server is http server.
package server

import "net/http"

type Server struct {
	http.Server
	mux *http.ServeMux
}

func NewServer() *Server {
	s := Server{}
	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type Config struct {
	StaticDir string
	LogDir    string
}
