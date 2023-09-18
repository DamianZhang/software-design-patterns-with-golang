package structs

import "fmt"

type DoubleFireCollisionHandler struct{}

func NewDoubleFireCollisionHandler() *DoubleFireCollisionHandler {
	return &DoubleFireCollisionHandler{}
}

func (d *DoubleFireCollisionHandler) Match(s1, s2 *Sprite) bool {
	return s1.String() == "F" && s2.String() == "F"
}

func (d *DoubleFireCollisionHandler) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	fmt.Println("Double Fire 移動失敗")
}
