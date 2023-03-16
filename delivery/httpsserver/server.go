package httpsserver

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/repository/sqldb"
	"love-date/service"
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
	profileService := service.NewProfileService(repo)
	partnerService := service.NewPartnerService(repo)
	profileHandler := handlre.NewProfileHandler(profileService)
	partnerHandler := handlre.NewPartnerHandler(partnerService)
	mux.Handle("/profile/create", http.HandlerFunc(profileHandler.CreateNewProfile))
	mux.Handle("/profile", http.HandlerFunc(profileHandler.GetUserProfile))
	mux.Handle("/partner/create", http.HandlerFunc(partnerHandler.CreateNewPartner))
	mux.Handle("/partner", http.HandlerFunc(partnerHandler.GetUserPartner))
	err := http.ListenAndServe(fmt.Sprintf(`%s:%d`, s.host, s.port), mux)
	if err != nil {

		panic("server stopped " + err.Error())
	}
}
