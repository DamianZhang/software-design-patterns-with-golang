package structs

type Gender int

const (
	Male Gender = iota
	Female
)

var genderStr = [2]string{"Male", "Female"}

func (g Gender) String() string {
	return genderStr[g]
}
