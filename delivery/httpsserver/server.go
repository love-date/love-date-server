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

	fmt.Printf("server is runing on %s:%d\n", s.host, s.port)

	mux.Handle("/profile/create", http.HandlerFunc(profileHandler.CreateNewProfile))
	fmt.Println(http.MethodPost + " /profile/create --> create profile route")

	mux.Handle("/profile/", http.HandlerFunc(profileHandler.GetUserProfile))
	fmt.Println(http.MethodGet + " /profile --> get user profile route")

	mux.Handle("/partner/create", http.HandlerFunc(partnerHandler.CreateNewPartner))
	fmt.Println(http.MethodPost + " /partner/create --> create partner route")

	mux.Handle("/partner", http.HandlerFunc(partnerHandler.GetUserPartner))
	fmt.Println(http.MethodGet + " /partner --> get user active partner route")

	err := http.ListenAndServe(fmt.Sprintf(`%s:%d`, s.host, s.port), mux)
	if err != nil {

		panic("server stopped " + err.Error())
	}
}
