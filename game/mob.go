package game

import "time"

var (
	Level0 = Level{0, ZeroMinerals, ZeroMinerals}

	mob1 = NewMob("SD-DX", time.Second * 5, ZeroMinerals)
	mob2 = NewMob("Alien", time.Second * 10, NewMinerals(3,0 ,0 ,0))
	mob3 = NewMob("Mudrox", time.Second * 15, ZeroMinerals)
	mob4 = NewMob("Xertorifiro Xato", time.Second * 5, NewMinerals(0, 6,0 ,0))

	mob1Rewards = map[int]Level{
		1: Level{Cost: NewMinerals()}
	}
)

type Mob struct {
	Name string
	Cost Minerals
	SpawnTime time.Duration
	LastSpawn time.Time
	Level Level
}

type Level struct {
	Value int
	Cost Minerals
	Produce Minerals
}

func NewMob(name string, spawnTime time.Duration, cost map[Mineral]int) *Mob {
	return &Mob{
		name,
		cost,
		spawnTime,
		time.Now(),
		Level0,
	}
}

func (m *Mob) Tick() Minerals {
	// Если время тика не настало
	if (time.Since(m.LastSpawn) < m.SpawnTime) {
		return ZeroMinerals
	}

	m.LastSpawn = time.Now()

	return m.Level.Produce
}