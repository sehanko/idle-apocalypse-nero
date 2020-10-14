package game

import "time"

var (
	mob1 = NewMob("Trinity", time.Second*5, mob1Rewards)
	mob2 = NewMob("Bronte", time.Second*10, mob2Rewards)

	mob1Rewards = map[int]*Level{
		1: &Level{1, ZeroResources, NewResources(BlueResource(1), GoalResource(5))},
		2: &Level{2, NewResources(BlueResource(5)), NewResources(BlueResource(2), GoalResource(10))},
		3: &Level{3, NewResources(BlueResource(50)), NewResources(BlueResource(3), GoalResource(25))},
		4: &Level{4, NewResources(GreenResource(50)), NewResources(BlueResource(4), GoalResource(50))},
	}

	mob2Rewards = map[int]*Level{
		1: &Level{1, NewResources(BlueResource(250)), NewResources(GreenResource(1), GoalResource(25))},
		2: &Level{2, NewResources(BlueResource(250), GreenResource(50)), NewResources(GreenResource(2), GoalResource(50))},
		3: &Level{3, NewResources(GreenResource(300)), NewResources(GreenResource(3), GoalResource(125))},
		4: &Level{4, NewResources(GreenResource(1250)), NewResources(GreenResource(4), GoalResource(250))},
	}
)

type Mob struct {
	Name      string
	SpawnTime time.Duration
	LastSpawn time.Time
	Level     *Level
	Levels    map[int]*Level
}

type Level struct {
	Value   int
	Cost    Resources
	Produce Resources
}

func NewMob(name string, spawnTime time.Duration, levels map[int]*Level) *Mob {
	return &Mob{
		name,
		spawnTime,
		time.Now(),
		levels[1],
		levels,
	}
}

func (m *Mob) Tick() Resources {
	if time.Since(m.LastSpawn) < m.SpawnTime {
		return ZeroResources
	}

	m.LastSpawn = time.Now()

	return m.Level.Produce
}

func (m *Mob) MaxLevel() int {
	return len(m.Levels)
}

func (m *Mob) LevelUp(res *Resources) bool {
	curLvl := m.Level.Value

	if curLvl == m.MaxLevel() {
		return false
	}

	nextLvl := m.Levels[curLvl+1]

	success := res.Widthdraw(&nextLvl.Cost)

	if success {
		m.Level = nextLvl
		return true
	}

	return false
}
