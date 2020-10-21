package game

import (
	"fmt"
	"time"
)

const globalTick = 1 * time.Second

type World struct {
	mobs  []*Mob
	res   Resources
	boost *Booster
}

func NewWorld() *World {
	w := World{}

	w.AddMob(mob1)
	w.AddMob(mob2)

	w.boost = NewBooster()

	return &w
}

func (w *World) Run() {
	//go func() {
	for range time.Tick(globalTick) {
		for _, mob := range w.mobs {
			resources := mob.Tick()

			if resources.IsEmpty() {
				continue
			}

			w.res.Append(&resources)

			success := mob.LevelUp(&w.res)

			if success {
				fmt.Println("mob ", mob.Name, " updated, level ", mob.Level.Value)
			}
		}

		fmt.Println(w.res)
	}
	//}()
}

func (w *World) AddMob(mob *Mob) {
	w.mobs = append(w.mobs, mob)
}

func (w *World) AddBoost(b 
	Boost) {
	w.boost.Add(b)
}
