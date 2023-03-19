package httpsserver

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"love-date/delivery/httpsserver/response"
	"love-date/entity"
	"love-date/pkg/jwttoken"
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
	userService := service.NewUserService(repo, partnerService, profileService)
	userHandler := handlre.NewUserHandler(userService)

	fmt.Printf("server is runing on %s:%d\n", s.host, s.port)

	mux.Handle("/user/append-name", middleware.AuthMiddleware(http.HandlerFunc(userHandler.AppendNames)))
	fmt.Println(http.MethodGet + " /user/append-name --> get append user and partner names route")

	mux.Handle("/profile/create", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.CreateNewProfile)))
	fmt.Println(http.MethodPost + " /profile/create --> create profile route")

	mux.Handle("/profile/get-one", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.GetUserProfile)))
	fmt.Println(http.MethodGet + " /profile --> get user profile route")

	mux.Handle("/profile/update", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.UpdateProfile)))
	fmt.Println(http.MethodPut + " /profile/update --> update profile route")

	mux.Handle("/partner/create", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.CreateNewPartner)))
	fmt.Println(http.MethodPost + " /partner/create --> create partner route")

	mux.Handle("/partner/get-active", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.GetUserPartner)))
	fmt.Println(http.MethodGet + " /partner --> get user active partner route")

	mux.Handle("/partner/update", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.UpdatePartner)))
	fmt.Println(http.MethodPut + " /partner/update --> update partner route")

	mux.Handle("/partner/delete", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.DeleteActivePartner)))
	fmt.Println(http.MethodGet + " /partner --> delete user active partner route")

	mux.Handle("/create-user", http.HandlerFunc(createUser))

	mux.Handle("/jwt", http.HandlerFunc(generateJWT))

	err := http.ListenAndServe(fmt.Sprintf(`%s:%d`, s.host, s.port), mux)
	if err != nil {

		panic("server stopped " + err.Error())
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

	repo := sqldb.New()
	profileService := service.NewProfileService(repo)
	partnerService := service.NewPartnerService(repo)
	userService := service.NewUserService(repo, partnerService, profileService)

	createUserRequest := &service.UserCreateRequest{}

	dErr := handlre.DecodeJSON(r.Body, createUserRequest)
	if dErr != nil {
		response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

		return
	}

	user, err := userService.Create(service.UserCreateRequest{Email: createUserRequest.Email})
	if err != nil {
		fmt.Printf("user error ", err)
		response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)

		return

	}
	token, err := jwttoken.GenerateJWT(user.User.ID, user.User.Email)
	if err != nil {
		fmt.Printf(err.Error())

		response.Fail(dErr.Error(), http.StatusUnauthorized).ToJSON(w)

		return
	}
	fmt.Printf(token)
	response.OK("profile loaded", token).ToJSON(w)
}

func generateJWT(w http.ResponseWriter, r *http.Request) {
	u := &entity.User{}
	handlre.DecodeJSON(r.Body, u)
	jwt, _ := jwttoken.GenerateJWT(u.ID, u.Email)

	response.OK("", jwt).ToJSON(w)
}
