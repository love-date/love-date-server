package specialday

import (
	"encoding/json"
	"love-date/constant"
	"love-date/pkg/errhandling/richerror"
	"os"
)

type SpecialDays = map[int]string

func GetSpecialDays() (*SpecialDays, error) {
	const op = "specialday-pkg.GetSpecialDays"

	var specialDays = new(SpecialDays)

	file, rErr := os.ReadFile(constant.SpecialDaysFilePath)
	if rErr != nil {

		return nil, richerror.New(op).WithWrapError(rErr).WithMessage(rErr.Error()).
			WithKind(richerror.KindUnexpected)
	}

	if uErr := json.Unmarshal(file, specialDays); uErr != nil {

		return nil, richerror.New(op).WithWrapError(rErr).WithMessage(rErr.Error()).
			WithKind(richerror.KindUnexpected)
	}

	return specialDays, nil
}
