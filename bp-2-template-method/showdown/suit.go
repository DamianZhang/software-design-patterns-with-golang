package showdown

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

var suitStr = [4]string{"Club", "Diamond", "Heart", "Spade"}

func (s Suit) String() string {
	return suitStr[s]
}
