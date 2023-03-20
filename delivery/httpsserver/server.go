package httpsserver

import (
	"fmt"
	"love-date/delivery/httpsserver/route"
	"love-date/repository/sqldb"
	"net/http"
)

type Server struct {
	host string
	port int
}

func NewHttpServer(host string, port int) *Server {
	return &Server{host, port}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	repo := sqldb.New()
	route.SetProfileRoute(mux, repo)
	route.SetPartnerRoute(mux, repo)
	route.SetUserRoute(mux)

	fmt.Printf("server is runing on %s:%d\n", s.host, s.port)

	err := http.ListenAndServe(fmt.Sprintf(`%s:%d`, s.host, s.port), mux)
	if err != nil {

		panic("server stopped " + err.Error())
	}
}
