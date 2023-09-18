package structs

import "fmt"

type DoubleWaterCollisionHandler struct{}

func NewDoubleWaterCollisionHandler() *DoubleWaterCollisionHandler {
	return &DoubleWaterCollisionHandler{}
}

func (d *DoubleWaterCollisionHandler) Match(s1, s2 *Sprite) bool {
	return s1.String() == "W" && s2.String() == "W"
}

func (d *DoubleWaterCollisionHandler) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	fmt.Println("Double Water 移動失敗")
}
