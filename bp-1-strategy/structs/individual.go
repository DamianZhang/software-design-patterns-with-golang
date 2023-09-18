package structs

import "fmt"

var (
	incrementCounter int    = 1
	MALE             string = "MALE"
	FEMALE           string = "FEMALE"
)

type Individual struct {
	id     int
	gender string
	age    int
	intro  string
	habits []string
	coord  *Coord
}

func NewIndividual(individualProps *IndividualProps) *Individual {
	individual := &Individual{}
	individual.SetId()
	individual.SetGender(individualProps.gender)
	individual.SetAge(individualProps.age)
	individual.SetIntro(individualProps.intro)
	individual.SetHabits(individualProps.habits)
	individual.SetCoord(individualProps.coord)
	return individual
}

func (i *Individual) Id() int {
	return i.id
}

func (i *Individual) SetId() {
	i.id = incrementCounter
	incrementCounter++
}

func (i *Individual) SetGender(gender string) {
	if gender == MALE || gender == FEMALE {
		i.gender = gender
	}
}

func (i *Individual) SetAge(age int) {
	if age >= 18 {
		i.age = age
	}
}

func (i *Individual) SetIntro(intro string) {
	lenOfIntro := len(intro)
	if lenOfIntro >= 0 && lenOfIntro <= 200 {
		i.intro = intro
	}
}

func (i *Individual) Habits() []string {
	return i.habits
}

func (i *Individual) SetHabits(habits []string) {
	for _, habit := range habits {
		if habitIsExisting(habit, i.habits) {
			continue
		}

		lenOfHabit := len(habit)
		if lenOfHabit >= 1 && lenOfHabit <= 10 {
			i.habits = append(i.habits, habit)
		}
	}
}

func habitIsExisting(habit string, habits []string) bool {
	for _, h := range habits {
		if habit == h {
			return true
		}
	}
	return false
}

func (i *Individual) Coord() *Coord {
	return i.coord
}

func (i *Individual) SetCoord(coord *Coord) {
	i.coord = coord
}

func (c *Individual) String() string {
	return fmt.Sprintf("id: %d, gender: %s, age: %d, intro: %s, habits: %v, coord: %v", c.id, c.gender, c.age, c.intro, c.habits, c.coord)
}
