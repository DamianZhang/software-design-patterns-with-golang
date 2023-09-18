package structs

const (
	NUM_OF_KEYS int = 26
)

type Keyboard struct {
	keys [NUM_OF_KEYS]Key
}

func NewKeyboard() *Keyboard {
	k := &Keyboard{}

	for key := 0; key < NUM_OF_KEYS; key++ {
		k.keys[key] = Key(key)
	}

	return k
}

func (k *Keyboard) Keys() [NUM_OF_KEYS]Key {
	return k.keys
}
