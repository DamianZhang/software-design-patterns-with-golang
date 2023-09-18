package structs

import "fmt"

type DoubleHeroesCollisionHandler struct{}

func NewDoubleHeroesCollisionHandler() *DoubleHeroesCollisionHandler {
	return &DoubleHeroesCollisionHandler{}
}

func (d *DoubleHeroesCollisionHandler) Match(s1, s2 *Sprite) bool {
	return s1.String() == "H" && s2.String() == "H"
}

func (d *DoubleHeroesCollisionHandler) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	fmt.Println("Double Heroes 移動失敗")
}
