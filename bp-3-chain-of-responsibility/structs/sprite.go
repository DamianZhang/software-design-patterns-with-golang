package structs

type Sprite struct {
	coordinate *Coordinate
	polytype   ISprite
}

type ISprite interface {
	String() string
	Hp() int
	SetHp(hp int)
	IsDead() bool
}

func NewSprite(x int, polytype ISprite) (*Sprite, error) {
	s := &Sprite{}

	c, err := NewCoordinate(x)
	if err != nil {
		return s, err
	}

	s.coordinate = c
	s.polytype = polytype
	return s, nil
}

func (s Sprite) String() string {
	return s.polytype.String()
}

func (s *Sprite) Hp() int {
	return s.polytype.Hp()
}

func (s *Sprite) SetHp(hp int) {
	s.polytype.SetHp(hp)
}

func (s *Sprite) IsDead() bool {
	return s.polytype.IsDead()
}

func (s *Sprite) Coordinate() *Coordinate {
	return s.coordinate
}

func (s *Sprite) SetCoordinate(coordinate *Coordinate) {
	s.coordinate = coordinate
}
