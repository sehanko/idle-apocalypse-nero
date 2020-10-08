package main

import (
	"github.com/sehanko/idle-apocalypse-nero/game"
)

func main()  {
	world := game.NewWorld()

	world.Run()
}