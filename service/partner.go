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
}
type PartnerService struct {
	repo PartnerServiceRepository
}

func NewPartnerService(repo PartnerServiceRepository) PartnerService {

	return PartnerService{repo}
}

type CreatePartnerRequest struct {
	Name                string    `json:"name"`
	Birthday            time.Time `json:"birthday"`
	FirstDate           time.Time `json:"first_date"`
	AuthenticatedUserID int
}

type CreatePartnerResponse struct {
	Partner entity.Partner
}

func (p PartnerService) Create(req CreatePartnerRequest) (CreatePartnerResponse, error) {
	if len(req.Name) < 2 {

		return CreatePartnerResponse{}, fmt.Errorf("the name's len must be longer than 1")
	}

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
	Name                string    `json:"name"`
	Birthday            time.Time `json:"birthday"`
	FirstDate           time.Time `json:"first_date"`
}

type UpdatePartnerResponse struct {
	Partner entity.Partner
}

func (p PartnerService) Update(req UpdatePartnerRequest) (UpdatePartnerResponse, error) {
	if len(req.Name) < 2 {

		return UpdatePartnerResponse{}, fmt.Errorf("the name's len must be longer than 1")
	}

	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return UpdatePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return UpdatePartnerResponse{}, fmt.Errorf("the partner not found")
	}

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

func (p PartnerService) Remove(req RemovePartnerRequest) (bool, error) {
	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return false, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return false, fmt.Errorf("the partner not found")
	}

	partner.Delete()

	if _, uErr := p.repo.UpdatePartner(partner.ID, partner); uErr != nil {

		return false, fmt.Errorf("unexpected error : %w", uErr)
	} else {

		return true, nil
	}

}

type GetUserActivePartnerRequest struct {
	AuthenticatedUserID int
}

type GetUserActivePartnerResponse struct {
	Partner entity.Partner
}

func (p PartnerService) GetUserActivePartner(req GetUserActivePartnerRequest) (GetUserActivePartnerResponse, error) {
	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return GetUserActivePartnerResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !partnerExist {

		return GetUserActivePartnerResponse{}, fmt.Errorf("this user has't any active partner")
	}

	return GetUserActivePartnerResponse{partner}, err
}

// CalculateTimeHasBeenTogether TODO: List of how time has been together
