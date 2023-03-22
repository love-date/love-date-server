package httpsserver

import (
	"fmt"
	"love-date/delivery/httpsserver/route"
	"love-date/pkg/oauth"
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

	partnerService := service.NewPartnerService(repo)
	profileService := service.NewProfileService(repo)
	userService := service.NewUserService(repo, partnerService, profileService)
	oauthService := oauth.NewOauthProvider()
	authService := service.NewAuthService(oauthService, userService)
	route.SetProfileRoute(mux, &profileService)
	route.SetPartnerRoute(mux, &partnerService)
	route.SetAuthRoute(mux, &authService)
	route.SetUserRoute(mux, &userService)
	route.SetAppRoute(mux)

	fmt.Printf("server is runing on %s:%d\n", s.host, s.port)

	err := http.ListenAndServe(fmt.Sprintf(`%s:%d`, s.host, s.port), mux)
	if err != nil {

		panic("server stopped " + err.Error())
	}
}
