package structs

import (
	"math"
	"regexp"
)

func IntShouldBeInRange(number, minNumber, maxNumber int) bool {
	return number >= minNumber && number <= maxNumber
}

func Float64ShouldBeInRange(number, minNumber, maxNumber float64) bool {
	return number >= minNumber && number <= maxNumber
}

func LenOfStringShouldBeInRange(s string, minLen, maxLen int) bool {
	return len(s) >= minLen && len(s) <= maxLen
}

func CalculateBMI(height, weight float64) int {
	return int(weight / math.Pow(height/100, 2))
}

func IsLegalPatientId(patientId string) bool {
	match, _ := regexp.MatchString("\\b[A-Z]{1}[0-9]{9}\\b", patientId)
	return match
}
