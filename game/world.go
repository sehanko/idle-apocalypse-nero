package game

import (
	"fmt"
	"time"
)

const globalTick = 1 * time.Second

type World struct {
	mobs []*Mob
	player Player
}


func NewWorld() *World {
	world := World{}

	world.AddMob(mob1)
	world.AddMob(mob2)
	world.AddMob(mob3)
	world.AddMob(mob4)

	world.player = NewPlayer()

	return &world
}

func (w *World) Run() {
	//go func() {
		for range time.Tick(globalTick) {
			for _, mob := range w.mobs {
				minerals := mob.Tick()

				if minerals.IsEmpty() {
					continue
				}

				w.player.Minerals.Append(&minerals)
			}


			fmt.Println(w.player)
		}
	//}()
}


func (w *World) AddMob(mob *Mob) {
	w.mobs = append(w.mobs, mob)
}