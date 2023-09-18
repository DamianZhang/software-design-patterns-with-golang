package structs

type CollisionHandler struct {
	polytype ICollisionHandler
	next     *CollisionHandler
}

type ICollisionHandler interface {
	Match(s1, s2 *Sprite) bool
	DoHandlingCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite)
}

func NewCollisionHandler(polytype ICollisionHandler, next *CollisionHandler) *CollisionHandler {
	return &CollisionHandler{
		polytype: polytype,
		next:     next,
	}
}

func (c *CollisionHandler) HandleCollision(s1, s2 *Sprite, spritesOfWorld *[LENGTH_OF_WORLD]*Sprite) {
	if c.polytype.Match(s1, s2) {
		c.polytype.DoHandlingCollision(s1, s2, spritesOfWorld)
	} else if c.next != nil {
		c.next.HandleCollision(s1, s2, spritesOfWorld)
	}
}
