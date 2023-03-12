package service

import (
	"fmt"
	"love-date/entity"
	"time"
)

type PartnerServiceRepository interface {
	CreatePartner(partner entity.Partner) (entity.Partner, error)
	UpdatePartner(partnerID int, partner entity.Partner) (entity.Partner, error)
	DoesUserHaveActivePartner(userID int) (bool, entity.Partner, error)
	AppendNameWithPartner(userID int) (string, error)
}
type Partner struct {
	repo PartnerServiceRepository
}

type CreatePartnerRequest struct {
	Name                string
	Birthday            time.Time
	FirstDate           time.Time
	AuthenticatedUserID int
}

type CreatePartnerResponse struct {
	Partner entity.Partner
}

func (p Partner) Create(req CreatePartnerRequest) (CreatePartnerResponse, error) {
	partnerExist, _, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return CreatePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if partnerExist {

		return CreatePartnerResponse{}, fmt.Errorf("this user has active partner")
	}

	if createdPartner, cErr := p.repo.CreatePartner(entity.Partner{
		UserID:    req.AuthenticatedUserID,
		Name:      req.Name,
		Birthday:  req.Birthday,
		FirstDate: req.FirstDate,
	}); cErr != nil {

		return CreatePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	} else {

		return CreatePartnerResponse{createdPartner}, nil
	}
}

type UpdatePartnerRequest struct {
	AuthenticatedUserID int
	ProfileID           int
	Name                string
	Birthday            time.Time
	FirstDate           time.Time
}

type UpdatePartnerResponse struct {
	Partner entity.Partner
}

func (p Partner) Update(req UpdatePartnerRequest) (UpdatePartnerResponse, error) {
	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return UpdatePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return UpdatePartnerResponse{}, fmt.Errorf("the partner not found")
	}

	//if partner.UserID != req.AuthenticatedUserID {
	//
	//	return UpdatePartnerResponse{}, fmt.Errorf("this user doesn't have permission to update this partner")
	//}

	if updatedPartner, uErr := p.repo.UpdatePartner(partner.ID, entity.Partner{
		Name:      req.Name,
		Birthday:  req.Birthday,
		FirstDate: req.FirstDate,
	}); uErr != nil {

		return UpdatePartnerResponse{}, fmt.Errorf("unexpected error : %w", uErr)
	} else {

		return UpdatePartnerResponse{updatedPartner}, nil
	}

}

type RemovePartnerRequest struct {
	AuthenticatedUserID int
}

type RemovePartnerResponse struct {
	Partner entity.Partner
}

func (p Partner) Remove(req RemovePartnerRequest) (RemovePartnerResponse, error) {
	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return RemovePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return RemovePartnerResponse{}, fmt.Errorf("the partner not found")
	}

	//if partner.UserID != req.AuthenticatedUserID {
	//
	//	return RemovePartnerResponse{}, fmt.Errorf("this user doesn't have permission to update this partner")
	//}

	partner.Delete()

	if updatedPartner, uErr := p.repo.UpdatePartner(partner.ID, partner); uErr != nil {

		return RemovePartnerResponse{}, fmt.Errorf("unexpected error : %w", uErr)
	} else {

		return RemovePartnerResponse{updatedPartner}, nil
	}

}

type GetUserActivePartnerRequest struct {
	AuthenticatedUserID int
}

type GetUserActivePartnerResponse struct {
	Partner entity.Partner
}

func (p Partner) GetUserActivePartner(req GetUserActivePartnerRequest) (GetUserActivePartnerResponse, error) {
	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return GetUserActivePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return GetUserActivePartnerResponse{}, fmt.Errorf("this user has't any active partner")
	}

	return GetUserActivePartnerResponse{partner}, err
}

type AppendPartnerNameRequest struct {
	AuthenticatedUserID int
}

type AppendPartnerNameResponse struct {
	AppendNames string
}

func (p Partner) AppendNames(req AppendPartnerNameRequest) (AppendPartnerNameResponse, error) {
	partnerExist, _, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return AppendPartnerNameResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return AppendPartnerNameResponse{}, fmt.Errorf("this user has't any active partner")
	}

	if appendNames, aErr := p.repo.AppendNameWithPartner(req.AuthenticatedUserID); aErr != nil {

		return AppendPartnerNameResponse{}, fmt.Errorf("unexpected error : %w", err)
	} else {

		return AppendPartnerNameResponse{appendNames}, nil
	}
}

// CalculateTimeHasBeenTogether TODO: List of how time has been together
