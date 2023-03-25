package specialday

import (
	"encoding/json"
	"fmt"
	"love-date/constant"
	"os"
)

type SpecialDays = map[int]string

func GetSpecialDays() (*SpecialDays, error) {
	var specialDays = new(SpecialDays)

	file, rErr := os.ReadFile(constant.SpecialDaysFilePath)
	if rErr != nil {

		return nil, fmt.Errorf("unexpected error : can't read file: %w", rErr)
	}

	if uErr := json.Unmarshal(file, specialDays); uErr != nil {

		return nil, fmt.Errorf("unexpected error : can't unmarshal from file: %w", uErr)
	}

	return specialDays, nil
}
