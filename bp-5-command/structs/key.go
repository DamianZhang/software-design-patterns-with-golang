package structs

import "errors"

type Key int

const (
	A Key = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

var keyStr = [NUM_OF_KEYS]string{"A", "B", "C", "D", "E",
	"F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z"}

func (k Key) String() string {
	return keyStr[k]
}

func ConvertAlphabetToKey(alphabet string) (Key, error) {
	for i, key := range keyStr {
		if alphabet == key {
			return Key(i), nil
		}
	}

	return Key(-1), errors.New("alphabet NOT in keyStr")
}
