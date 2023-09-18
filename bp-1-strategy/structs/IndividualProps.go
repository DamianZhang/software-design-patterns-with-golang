package structs

type IndividualProps struct {
	gender string
	age    int
	intro  string
	habits []string
	coord  *Coord
}

func NewIndividualProps(gender string, age int, intro string, habits []string, coord *Coord) *IndividualProps {
	return &IndividualProps{
		gender: gender,
		age:    age,
		intro:  intro,
		habits: habits,
		coord:  coord,
	}
}
