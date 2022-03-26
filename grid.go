package worldservices

import (
	"errors"
	"fmt"
	"sort"

	cartesian "github.com/schwarmco/go-cartesian-product"
)

type Grid struct {
	objects   []WorldObject
	locations map[string]WorldObject
	Size      int
	cycle     int
}

func NewGrid(obj []WorldObject, loc map[string]WorldObject, size int) Grid {
	return Grid{obj, loc, size, 0}
}

func (g *Grid) GetOrderedObjectListByFitness() []WorldObject {
	ordered_objects := []WorldObject{}
	grouped_objects := map[float64][]WorldObject{}
	for _, obj := range g.objects {
		fitness := obj.Fitness()
		slice := grouped_objects[fitness]
		grouped_objects[fitness] = append(slice, obj)
	}
	keys := make([]float64, 0, len(grouped_objects))
	for k := range grouped_objects {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	for _, v := range keys {
		slice_objs := grouped_objects[v]
		for _, o := range slice_objs {
			ordered_objects = append(ordered_objects, o)
		}
	}
	return ordered_objects
}

func formatCoords(x int, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}

func (g *Grid) GetObjectAtCoords(x int, y int) (WorldObject, error) {
	loc := formatCoords(x, y)
	if val, ok := g.locations[loc]; ok {
		return val, nil
	}
	return nil, errors.New("Unable to find object at location")

}

// This basically means if any creatures overlap they get removed from the map.
// objects will still exists though.
// this map is used for object selection - probably not good things get removed.
// TODO improve this
func (g *Grid) UpdateLocationsCoords() {
	locations := map[string]WorldObject{}
	objects := g.GetOrderedObjectListByFitness()
	/*
		for i, j := 0, len(objects)-1; i < j; i, j = i+1, j-2 {
			obj := objects[i]
			x, y := obj.GetCoordsXY()
			locations[formatCoords(x, y)] = obj
		}*/
	for _, obj := range objects {
		x, y := obj.GetCoordsXY()
		locations[formatCoords(x, y)] = obj
	}
	g.locations = locations
}

/*
GetObjectSenseData

returns a list of objects that are in the sense area of the
given coordinates
*/
func (g *Grid) GetObjectSenseData(x int, y int, vision int) []WorldObject {
	xstart := x - vision
	ystart := y - vision
	xrange := make([]interface{}, vision*2)
	yrange := make([]interface{}, vision*2)
	objects := []WorldObject{}
	for i := range xrange {
		xrange[i] = xstart
		yrange[i] = ystart
		xstart = xstart + 1
		ystart = ystart + 1
	}
	c := cartesian.Iter(xrange, yrange)
	for coords := range c {
		cx := coords[0].(int)
		cy := coords[1].(int)
		// This check ensures the self is not in sense data.
		if cx == x && cy == y {
			continue
		}
		obj, err := g.GetObjectAtCoords(cx, cy)
		if err != nil {
			continue
		}
		objects = append(objects, obj)
	}
	return objects
}
