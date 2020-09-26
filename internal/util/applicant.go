package util

import (
	"epgu-generator/internal/model"
	"strings"
)

func toApplicantType(value string) model.ApplicantType {
	switch value {
	case "ЮЛ":
		return model.ApplicantLE
	case "ФЛ":
		return model.ApplicantNP
	case "ИП":
		return model.ApplicantIE
	}
	return model.ApplicantNone
}

// ParseApplicant ...
func ParseApplicant(value string) []model.ApplicantType {

	list := strings.Split(value, ",")
	result := make([]model.ApplicantType, len(list))

	for i, v := range list {
		result[i] = toApplicantType(v)
	}

	return result
}

// ParseUseSignature ...
func ParseUseSignature(value string) bool {
	switch value {
	case "да":
	case "1":
	case "y":
	case "yes":
	case "+":
		return true
	}
	return false
}
