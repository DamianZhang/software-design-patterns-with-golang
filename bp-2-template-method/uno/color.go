package uno

type Color int

const (
	Blue Color = iota
	Red
	Yellow
	Green
)

var colorStr = [4]string{"Blue", "Red", "Yellow", "Green"}

func (c Color) String() string {
	return colorStr[c]
}
