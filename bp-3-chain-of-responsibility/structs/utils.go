package structs

func ShouldBeInRange(number, minNumber, maxNumber int) bool {
	return number >= minNumber && number <= maxNumber
}
