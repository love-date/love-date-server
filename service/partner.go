package service

import (
	"love-date/entity"
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
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
	const op = "partner-service.Create"

	if len(req.Name) < 2 {

		return CreatePartnerResponse{}, richerror.New(op).WithMessage("the name's len must be longer than 1").
			WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
			"name": req.Name,
		})
	}

	partnerExist, _, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return CreatePartnerResponse{}, richerror.New(op).WithWrapError(err)
	}

	if partnerExist {

		return CreatePartnerResponse{}, richerror.New(op).WithMessage("this user has active partner").
			WithKind(richerror.KindForbidden)
	}

	if createdPartner, cErr := p.repo.CreatePartner(entity.Partner{
		UserID:    req.AuthenticatedUserID,
		Name:      req.Name,
		Birthday:  req.Birthday,
		FirstDate: req.FirstDate,
	}); cErr != nil {

		return CreatePartnerResponse{}, richerror.New(op).WithWrapError(cErr)
	} else {

		return CreatePartnerResponse{createdPartner}, nil
	}
}

type UpdatePartnerRequest struct {
	AuthenticatedUserID int
	Name                *string    `json:"name"`
	Birthday            *time.Time `json:"birthday"`
	FirstDate           *time.Time `json:"first_date"`
}

type UpdatePartnerResponse struct {
	Partner entity.Partner
}

func (p PartnerService) Update(req UpdatePartnerRequest) (UpdatePartnerResponse, error) {
	const op = "partner-service.Update"

	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return UpdatePartnerResponse{}, richerror.New(op).WithWrapError(err)
	}
	if !partnerExist {

		return UpdatePartnerResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound)
	}

	if req.Name != nil {
		if len(*req.Name) < 2 {

			return UpdatePartnerResponse{}, richerror.New(op).WithMessage("the name's len must be longer than 1 char").
				WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
				"name": req.Name,
			})
		}

		partner.Name = *req.Name
	}
	if req.FirstDate != nil {
		partner.FirstDate = *req.FirstDate
	}
	if req.Birthday != nil {
		partner.Birthday = *req.Birthday
	}

	if updatedPartner, uErr := p.repo.UpdatePartner(partner.ID, entity.Partner{
		Name:      partner.Name,
		Birthday:  partner.Birthday,
		FirstDate: partner.FirstDate,
	}); uErr != nil {

		return UpdatePartnerResponse{}, richerror.New(op).WithWrapError(uErr)
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
	const op = "partner-service.Remove"

	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return false, richerror.New(op).WithWrapError(err)
	}
	if !partnerExist {

		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound)
	}

	partner.Delete()

	if _, uErr := p.repo.UpdatePartner(partner.ID, partner); uErr != nil {

		return false, richerror.New(op).WithWrapError(uErr)
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
	const op = "partner-service.GetUserActivePartner"

	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
	if err != nil {

		return GetUserActivePartnerResponse{}, richerror.New(op).WithWrapError(err)
	}

	if !partnerExist {

		return GetUserActivePartnerResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound)
	}

	return GetUserActivePartnerResponse{partner}, err
}
