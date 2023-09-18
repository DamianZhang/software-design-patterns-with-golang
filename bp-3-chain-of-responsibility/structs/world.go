package structs

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	LENGTH_OF_WORLD        int = 30
	NUM_OF_INITIAL_SPRITES int = 10
	NUM_OF_POLYTYPE        int = 3
)

type World struct {
	sprites          *[LENGTH_OF_WORLD]*Sprite
	collisionHandler *CollisionHandler
}

func NewWorld(collisionHandler *CollisionHandler) (*World, error) {
	w := &World{sprites: new([LENGTH_OF_WORLD]*Sprite)}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < NUM_OF_INITIAL_SPRITES; {
		randX := r.Intn(LENGTH_OF_WORLD)

		if w.sprites[randX] == nil {
			randPolytype := getPolytype(r.Intn(NUM_OF_POLYTYPE))
			sprite, err := NewSprite(randX, randPolytype)
			if err != nil {
				return w, err
			}

			w.sprites[randX] = sprite
			i++
		}
	}

	w.collisionHandler = collisionHandler
	return w, nil
}

func getPolytype(number int) ISprite {
	switch number {
	case 0:
		return NewHero()
	case 1:
		return NewFire()
	case 2:
		return NewWater()
	default:
		return NewHero()
	}
}

func (w *World) Start() {
	fmt.Printf("collision detection of game is starting...\n")

	for {
		fmt.Println()
		w.printSprites()

		var input1, input2 string
		fmt.Printf("請輸入 x1 和 x2:\n")
		fmt.Scanf("%s %s", &input1, &input2)

		if input1 == input2 {
			panic(errors.New("x1 and x2 should NOT be same"))
		}

		x1, err := strconv.Atoi(input1)
		if err != nil {
			panic(err)
		}

		x2, err := strconv.Atoi(input2)
		if err != nil {
			panic(err)
		}

		c1, err := NewCoordinate(x1)
		if err != nil {
			panic(err)
		}

		c2, err := NewCoordinate(x2)
		if err != nil {
			panic(err)
		}

		err = w.moveSprite(c1, c2)
		if err != nil {
			panic(err)
		}
	}
}

func (w *World) moveSprite(c1, c2 *Coordinate) error {
	x1, x2 := c1.X(), c2.X()
	spritesOfWorld := w.sprites

	s1 := spritesOfWorld[x1]
	if s1 == nil {
		return errors.New("s1 should NOT be nil")
	}

	s2 := spritesOfWorld[x2]
	if s2 == nil {
		s1.SetCoordinate(c2)
		spritesOfWorld[x1] = nil
		spritesOfWorld[x2] = s1
		fmt.Println("移動成功")
	} else {
		w.handleCollision(s1, s2, spritesOfWorld)
	}

	return nil
}

func (w *World) handleCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	w.collisionHandler.HandleCollision(s1, s2, spritesOfWorld)
}

func (w *World) printSprites() {
	fmt.Printf("ALL sprites: %v\n", w.sprites)

	fmt.Printf("NOT NIL sprites: ")
	for i, sprite := range w.sprites {
		if sprite != nil {
			fmt.Printf("%v(%d), ", sprite, i)
		}
	}
	fmt.Println()
}
