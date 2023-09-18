package structs

import "fmt"

type WaterFireCollisionHandlers struct{}

func NewWaterFireCollisionHandlers() *WaterFireCollisionHandlers {
	return &WaterFireCollisionHandlers{}
}

func (w *WaterFireCollisionHandlers) Match(s1, s2 *Sprite) bool {
	return (s1.String() == "W" && s2.String() == "F") || (s1.String() == "F" && s2.String() == "W")
}

func (w *WaterFireCollisionHandlers) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	spritesOfWorld[s1.Coordinate().X()] = nil
	spritesOfWorld[s2.Coordinate().X()] = nil
	fmt.Println("Water, Fire 從世界中被移除")
}
