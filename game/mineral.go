package game

var (
	ZeroMinerals = NewMinerals(0, 0, 0,0)

	BlueMineral  = Mineral{"Blueium Rod", "blue"}
	GreenMineral = Mineral{"Greenium Rod", "green"}
	RedMineral   = Mineral{"Redium Rod", "red"}
	GoldMineral  = Mineral{"Goldium Rod", "gold"}
)

type Minerals map[Mineral]int

type Mineral struct {
	Name string
	Color string
}

// NewMinerals создаем сет из минералов
func NewMinerals(blue, green, red, gold int) Minerals {
	m := make(Minerals, 0)

	if blue > 0 {
		m[BlueMineral] = blue
	}

	if green > 0 {
		m[GreenMineral] = green
	}

	if red > 0 {
		m[RedMineral] = red
	}

	if gold > 0 {
		m[GoldMineral] = gold
	}

	return m
}

func (m *Minerals) IsEmpty() bool {
	return len(*m) <= 0
}

func (m *Minerals) Append(mm *Minerals) {
	for mineral, v := range *mm {
		m.Add(mineral, v)
	}
}

func (m Minerals) Add (mineral Mineral, value int) {
	_, ok := m[mineral]

	// Есил в списке небыло - добавляем
	if !ok {
		m[mineral] = value
		return
	}

	m[mineral] += value
}