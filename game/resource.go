package game

var (
	ZeroResources = NewResources(0, 0, 0, 0)

	BlueResource  = Resource{"Blueium Rod", "blue"}
	GreenResource = Resource{"Greenium Rod", "green"}
	RedResource   = Resource{"Redium Rod", "red"}
	GoldResource  = Resource{"Goldium Rod", "gold"}
)

type Resources struct {
	Values map[Resource]int
}

type Resource struct {
	Name  string
	Color string
}

// NewResources creating set of resources
func NewResources(blue, green, red, gold int) Resources {
	values := make(map[Resource]int, 0)
	res := Resources{values}

	if blue > 0 {
		res.Values[BlueResource] = blue
	}

	if green > 0 {
		res.Values[GreenResource] = green
	}

	if red > 0 {
		res.Values[RedResource] = red
	}

	if gold > 0 {
		res.Values[GoldResource] = gold
	}

	return res
}

func (rr *Resources) IsEmpty() bool {
	return len(rr.Values) <= 0
}

func (rr *Resources) Append(r Resources) {
	for res, v := range rr.Values {
		rr.Add(res, v)
	}
}

func (rr *Resources) Add(res Resource, value int) {
	_, ok := rr.Values[res]

	// Есил в списке небыло - добавляем
	if !ok {
		rr.Values[res] = value
		return
	}

	rr.Values[res] += value
}
