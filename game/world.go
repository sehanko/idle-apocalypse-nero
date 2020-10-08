package game

import (
	"fmt"
	"time"
)

const globalTick = 1 * time.Second

type World struct {
	mobs   []*Mob
	player Player
}

func NewWorld() *World {
	world := World{}

	world.AddMob(mob1)
	world.AddMob(mob2)

	world.player = NewPlayer()

	return &world
}

func (w *World) Run() {
	//go func() {
	for range time.Tick(globalTick) {
		for _, mob := range w.mobs {
			resources := mob.Tick()

			if resources.IsEmpty() {
				continue
			}

			w.player.Res.Append(resources)
		}

		fmt.Println(w.player)
	}
	//}()
}

func (w *World) AddMob(mob *Mob) {
	w.mobs = append(w.mobs, mob)
}
