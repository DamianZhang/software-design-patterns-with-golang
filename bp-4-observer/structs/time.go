package structs

type Time int

func NewTime(second int) Time {
	return Time(second)
}

func (t Time) Second() int {
	return int(t)
}

func (t Time) Minute() int {
	return int(t) / 60
}
