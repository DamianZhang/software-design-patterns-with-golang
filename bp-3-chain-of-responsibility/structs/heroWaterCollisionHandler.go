package structs

import "fmt"

type HeroWaterCollisionHandler struct{}

func NewHeroWaterCollisionHandler() *HeroWaterCollisionHandler {
	return &HeroWaterCollisionHandler{}
}

func (h *HeroWaterCollisionHandler) Match(s1, s2 *Sprite) bool {
	return (s1.String() == "H" && s2.String() == "W") || (s1.String() == "W" && s2.String() == "H")
}

func (h *HeroWaterCollisionHandler) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	fmt.Println("Hero +10 HP")
	fmt.Println("Water 從世界中被移除")

	if s1.String() == "H" {
		spritesOfWorld[s1.Coordinate().X()] = nil

		s1.SetHp(s1.Hp() + 10)
		s1.SetCoordinate(s2.Coordinate())
		spritesOfWorld[s2.Coordinate().X()] = s1

		fmt.Println("Hero 移動成功")
		fmt.Println("Hero HP:", s1.Hp())
	} else if s2.String() == "H" {
		s2.SetHp(s2.Hp() + 10)
		spritesOfWorld[s1.Coordinate().X()] = nil

		fmt.Println("Hero HP:", s2.Hp())
	}
}
