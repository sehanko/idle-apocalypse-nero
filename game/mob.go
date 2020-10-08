package game

import "time"

var (
	mob1 = NewMob("Trinity", time.Second*5)
	mob2 = NewMob("Alien", time.Second*10)

	mob1Rewards = map[int]Level{
		1: Level{Cost: ZeroResources, Produce: NewResources(BlueResource(1))},
		2: Level{Cost: ZeroResources, Produce: NewResources(BlueResource(2))},
		3: Level{Cost: ZeroResources, Produce: NewResources(BlueResource(3))},
		4: Level{Cost: ZeroResources, Produce: NewResources(BlueResource(4))},
	}
)

type Mob struct {
	Name      string
	SpawnTime time.Duration
	LastSpawn time.Time
	Level     Level
}

type Level struct {
	Value   int
	Cost    Resources
	Produce Resources
}

func NewMob(name string, spawnTime time.Duration) *Mob {
	return &Mob{
		name,
		spawnTime,
		time.Now(),
		mob1Rewards[0],
	}
}

func (m *Mob) Tick() Resources {
	if time.Since(m.LastSpawn) < m.SpawnTime {
		return ZeroResources
	}

	m.LastSpawn = time.Now()

	return m.Level.Produce
}
