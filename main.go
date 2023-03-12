package main

import (
	"fmt"
	"love-date/repository/sqldb"
	"love-date/service"
)

func main() {
	repo := sqldb.New()
	//userService := service.NewUserService(repo)
	//res, err := userService.Create(service.UserCreateRequest{Email: "masooodsk@gmail.com"})
	//if err != nil {
	//	fmt.Println("error", err)
	//}
	//fmt.Println("user:", res.User)
	profileService := service.NewProfileService(repo)
	res2, err := profileService.Create(service.CreateProfileRequest{
		Name:                    "elaa",
		BirthdayNotifyActive:    true,
		SpecialDaysNotifyActive: true,
		AuthenticatedUserID:     3,
	})
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(res2)
	//res2, err := profileService.GetUserProfileByID(service.GetProfileByIDRequest{AuthenticatedUserID: 4})
	//if err != nil {
	//	fmt.Println("error", err)
	//}

	//res2, err := profileService.Update(service.UpdateProfileRequest{
	//	Name:                    "elaxe",
	//	BirthdayNotifyActive:    false,
	//	SpecialDaysNotifyActive: true,
	//	AuthenticatedUserID:     4,
	//})
	//if err != nil {
	//	fmt.Println("error", err)
	//}
	//fmt.Println("profile:", res2.Profile)
}
