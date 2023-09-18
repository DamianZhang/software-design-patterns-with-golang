package structs

import "fmt"

type HeroFireCollisionHandler struct{}

func NewHeroFireCollisionHandler() *HeroFireCollisionHandler {
	return &HeroFireCollisionHandler{}
}

func (h *HeroFireCollisionHandler) Match(s1, s2 *Sprite) bool {
	return (s1.String() == "H" && s2.String() == "F") || (s1.String() == "F" && s2.String() == "H")
}

func (h *HeroFireCollisionHandler) DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	fmt.Println("Hero -10 HP")
	fmt.Println("Fire 從世界中被移除")

	if s1.String() == "H" {
		spritesOfWorld[s1.Coordinate().X()] = nil

		s1.SetHp(s1.Hp() - 10)
		if s1.IsDead() {
			spritesOfWorld[s2.Coordinate().X()] = nil
			fmt.Println("Hero is dead, Hero 從世界中被移除")
		} else {
			s1.SetCoordinate(s2.Coordinate())
			spritesOfWorld[s2.Coordinate().X()] = s1
			fmt.Println("Hero 移動成功")
		}

		fmt.Println("Hero HP:", s1.Hp())
	} else if s2.String() == "H" {
		s2.SetHp(s2.Hp() - 10)
		if s2.IsDead() {
			spritesOfWorld[s2.Coordinate().X()] = nil
		}
		spritesOfWorld[s1.Coordinate().X()] = nil

		fmt.Println("Hero HP:", s2.Hp())
	}
}
