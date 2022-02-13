package worldservices_test

import (
	"testing"

	"github.com/lukegriffith/worldservices"
)

type TestObject struct {
	x, y int
}

func (t TestObject) Fitness() float64 {
	return float64(1)
}

func (t TestObject) GetCoordsXY() (int, int) {
	return t.x, t.y
}

func (t TestObject) Process(g worldservices.Grid, osc float64) {
}

func (b TestObject) SetDebug() {
}

func TestGetObjectSenseData(t *testing.T) {
	objects := []worldservices.WorldObject{
		TestObject{3, 3},
		TestObject{2, 2},
	}
	g := SetupTestGrid(objects)
	obj := g.GetObjectSenseData(2, 3, 10)
	t.Log(len(obj))
	t.Log(obj)
	if len(obj) < 2 {
		t.Fail()
	}

}

func SetupTestGrid(objects []worldservices.WorldObject) worldservices.Grid {

	g := worldservices.NewGrid(objects, map[string]worldservices.WorldObject{}, 10)
	g.UpdateLocationsCoords()
	return g
}
