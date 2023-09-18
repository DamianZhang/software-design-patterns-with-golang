package structs

import "errors"

type Direction int

const (
	W Direction = iota
	S
	A
	D
)

const (
	LenOfDirection int = 4
)

var (
	symbolsOfDirection   = [4]string{SymbolOfCharacterUp, SymbolOfCharacterDown, SymbolOfCharacterLeft, SymbolOfCharacterRight}
	alphabetsOfDirection = [4]string{"W", "S", "A", "D"}
)

func (d Direction) String() string {
	return symbolsOfDirection[d]
}

func (d Direction) Alphabet() string {
	return alphabetsOfDirection[d]
}

func ConvertAlphabetToDirection(alphabet string) (Direction, error) {
	for direction, alphabetOfDirection := range alphabetsOfDirection {
		if alphabet == alphabetOfDirection {
			return Direction(direction), nil
		}
	}

	return Direction(-1), errors.New("alphabet NOT in alphabetsOfDirection")
}
