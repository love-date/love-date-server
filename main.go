package main

import (
	"fmt"
	"love-date/repository/sqldb"
	"love-date/service"
)

func main() {

	repo := sqldb.New()
	profileService := service.NewProfileService(repo)
	partnerService := service.NewPartnerService(repo)
	userService := service.NewUserService(repo, partnerService, profileService)
	res, err := userService.AppendNames(service.AppendPartnerNameRequest{AuthenticatedUserID: 4})
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("user:", res.AppendNames)
	//res, err := userService.Create(service.UserCreateRequest{Email: "masooodsk@gmail.com"})
	//if err != nil {
	//	fmt.Println("error", err)
	//}
	//fmt.Println("user:", res.User)
	//res2, err := profileService.Create(service.CreateProfileRequest{
	//	Name:                    "elaa",
	//	BirthdayNotifyActive:    true,
	//	SpecialDaysNotifyActive: true,
	//	AuthenticatedUserID:     3,
	//})
	//if err != nil {
	//	fmt.Println("error", err)
	//}
	//fmt.Println(res2)
	//res2, err := profileService.GetUserProfile(service.GetProfileByIDRequest{AuthenticatedUserID: 4})
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

	//res2, err := partnerService.Create(service.CreatePartnerRequest{
	//	Name:                "elaaaa",
	//	Birthday:            time.Now(),
	//	FirstDate:           time.Now(),
	//	AuthenticatedUserID: 4,
	//})
	//res2, err := partnerService.Update(service.UpdatePartnerRequest{
	//	AuthenticatedUserID: 3,
	//	Name:                "e",
	//	Birthday:            time.Now(),
	//	FirstDate:           time.Now(),
	//})
	//if err != nil {
	//	fmt.Println("error", err)
	//}
	//fmt.Println("profile:", res2)
}
