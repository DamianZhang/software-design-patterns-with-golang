package showdown

import "fmt"

type Rank int

func (r Rank) String() string {
	switch r {
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	default:
		return fmt.Sprintf("%d", r)
	}
}
