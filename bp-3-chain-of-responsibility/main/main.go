package main

import (
	"bp-3-chain-of-responsibility/structs"
)

func main() {
	world, err := structs.NewWorld(
		structs.NewCollisionHandler(structs.NewDoubleHeroesCollisionHandler(),
			structs.NewCollisionHandler(structs.NewDoubleWaterCollisionHandler(),
				structs.NewCollisionHandler(structs.NewDoubleFireCollisionHandler(),
					structs.NewCollisionHandler(structs.NewWaterFireCollisionHandlers(),
						structs.NewCollisionHandler(structs.NewHeroFireCollisionHandler(),
							structs.NewCollisionHandler(structs.NewHeroWaterCollisionHandler(), nil)))))))
	if err != nil {
		panic(err)
	}

	world.Start()
}
