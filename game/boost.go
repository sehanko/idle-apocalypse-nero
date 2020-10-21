package game

import "sync"

var (
	blueBoost  = Boost{"Blue boost", "blue", blueLevels[1], blueLevels}
	greenBoost = Boost{"Green boost", "green", greenLevels[1], greenLevels}

	blueLevels = map[int]*BoostLevel{
		1: &BoostLevel{1, BlueResource(1), BlueResource(5)},
		2: &BoostLevel{2, BlueResource(2), BlueResource(5)},
		3: &BoostLevel{3, BlueResource(3), BlueResource(5)},
		4: &BoostLevel{4, BlueResource(4), BlueResource(5)},
		5: &BoostLevel{5, BlueResource(5), BlueResource(5)},
	}

	greenLevels = map[int]*BoostLevel{
		1: &BoostLevel{1, GreenResource(1), GreenResource(5)},
		2: &BoostLevel{2, GreenResource(2), GreenResource(5)},
		3: &BoostLevel{3, GreenResource(3), GreenResource(5)},
		4: &BoostLevel{4, GreenResource(4), GreenResource(5)},
		5: &BoostLevel{5, GreenResource(5), GreenResource(5)},
	}
)

type Booster struct {
	List []Boost
	mx   sync.Mutex
}

type Boost struct {
	Name  string
	Color string
	Bonus *BoostLevel
	bList map[int]*BoostLevel
}

type BoostLevel struct {
	Value   int
	Cost    ResourceValue
	Produce ResourceValue
}

func NewBooster() *Booster {
	l := make([]Boost, 0, 8)

	return &Booster{l, sync.Mutex{}}
}

func (br *Booster) Add(b Boost) {
	br.mx.Lock()
	defer br.mx.Unlock()

	br.List = append(br.List, b)
}

func (b *Boost) MaxLevel() int {
	return len(b.bList)
}

func (b *Boost) LevelUp(res *Resources) bool {
	curLvl := b.Bonus.Value

	if curLvl == b.MaxLevel() {
		return false
	}

	nextLvl := b.bList[curLvl+1]

	success := res.Widthdraw(&nextLvl.Cost)

	if success {
		b.Bonus = nextLvl
		return true
	}

	return false
}
