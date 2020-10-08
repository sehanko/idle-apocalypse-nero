package game

var (
	ZeroResources = NewResources()
	ZeroValue     = ResourceValue{Resource: nil, Value: 0}

	blueResource  = Resource{"Blueium Rod", "blue"}
	GreenResource = Resource{"Greenium Rod", "green"}
	RedResource   = Resource{"Redium Rod", "red"}
	GoldResource  = Resource{"Goldium Rod", "gold"}
)

type Resources struct {
	Values map[*Resource]int
}

type Resource struct {
	Name  string
	Color string
}

type ResourceValue struct {
	*Resource
	Value int
}

func BlueResource(val int) ResourceValue {
	return ResourceValue{
		Resource: &blueResource,
		Value:    val,
	}
}

func NewResources(values ...ResourceValue) Resources {
	val := make(map[*Resource]int, len(values))
	res := Resources{val}

	for _, v := range values {
		res.Values[v.Resource] = v.Value
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

func (rr *Resources) Add(res *Resource, value int) {
	_, ok := rr.Values[res]

	// Есил в списке небыло - добавляем
	if !ok {
		rr.Values[res] = value
		return
	}

	rr.Values[res] += value
}
