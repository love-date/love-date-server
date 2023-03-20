package route

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

func SetUserRoute(mux *http.ServeMux) {
	repo := sqldb.New()
	profileService := service.NewProfileService(repo)
	partnerService := service.NewPartnerService(repo)
	userService := service.NewUserService(repo, partnerService, profileService)
	userHandler := handlre.NewUserHandler(userService)

	mux.Handle("/user/append-name", middleware.AuthMiddleware(http.HandlerFunc(userHandler.AppendNames)))
	fmt.Println(http.MethodGet + " /user/append-name --> get append user and partner names route")

	mux.Handle("/create-user", http.HandlerFunc(createUser))
	mux.Handle("/jwt", http.HandlerFunc(generateJWT))
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
		response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)

		return

	}
	token, err := jwttoken.GenerateJWT(user.User.ID, user.User.Email)
	if err != nil {

		response.Fail(dErr.Error(), http.StatusUnauthorized).ToJSON(w)

		return
	}
	response.OK("user created", token).ToJSON(w)
}

func generateJWT(w http.ResponseWriter, r *http.Request) {
	u := &entity.User{}
	handlre.DecodeJSON(r.Body, u)
	jwt, _ := jwttoken.GenerateJWT(u.ID, u.Email)

	response.OK("", jwt).ToJSON(w)
}
