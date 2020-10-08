package game

import (
	"sync"
)

var (
	ZeroResources = NewResources()
	ZeroValue     = ResourceValue{Resource: nil, Value: 0}

	blueResource  = Resource{"Blueium Rod", "blue"}
	greenResource = Resource{"Greenium Rod", "green"}
	redResource   = Resource{"Redium Rod", "red"}
	goldResource  = Resource{"Goldium Rod", "gold"}

	goalResource = Resource{"Goal", "goal"}
)

type Resources struct {
	Values map[*Resource]int
	mx     sync.Mutex
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

func GreenResource(val int) ResourceValue {
	return ResourceValue{
		Resource: &greenResource,
		Value:    val,
	}
}

func GoalResource(val int) ResourceValue {
	return ResourceValue{
		Resource: &goalResource,
		Value:    val,
	}
}

func NewResources(values ...ResourceValue) Resources {
	val := make(map[*Resource]int, len(values))
	res := Resources{val, sync.Mutex{}}

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

func (rr *Resources) Available(val *ResourceValue) bool {
	res, ok := rr.Values[val.Resource]

	if !ok {
		return false
	}

	if res >= val.Value {
		return true
	}

	return false
}

func (rr *Resources) Widthdraw(res *Resources) bool {
	rr.mx.Lock()
	defer rr.mx.Unlock()

	for r, val := range res.Values {
		if !rr.Available(r) {
			return false
		}
	}

	return true
}
