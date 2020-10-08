package game

import "time"

var (
	Level0 = Level{0, ZeroResources, ZeroResources}

	mob1 = NewMob("SD-DX", time.Second*5, ZeroResources)
	mob2 = NewMob("Alien", time.Second*10, NewResources(3, 0, 0, 0))
	mob3 = NewMob("Mudrox", time.Second*15, ZeroResources)
	mob4 = NewMob("Xertorifiro Xato", time.Second*5, NewResources(0, 6, 0, 0))

	mob1Rewards = map[int]Level{
		1: Level{Cost: NewResources(0, 0, 0, 0)},
	}
)

type Mob struct {
	Name      string
	Cost      Resources
	SpawnTime time.Duration
	LastSpawn time.Time
	Level     Level
}

type Level struct {
	Value   int
	Cost    Resources
	Produce Resources
}

func NewMob(name string, spawnTime time.Duration, cost Resources) *Mob {
	return &Mob{
		name,
		cost,
		spawnTime,
		time.Now(),
		Level0,
	}
}

func (m *Mob) Tick() Resources {
	if time.Since(m.LastSpawn) < m.SpawnTime {
		return ZeroResources
	}

	m.LastSpawn = time.Now()

	return m.Level.Produce
}
