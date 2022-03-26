package worldservices

import "github.com/jinzhu/copier"

type WorldHistory struct {
	timeline []GridHistory
}

func (wh *WorldHistory) Push(grid Grid) {
	gridCopy := GridHistory{}
	copier.Copy(&gridCopy, &grid)
	gridCopy.objects = []Fossil{}

	for _, creature := range grid.objects {
		X, Y := creature.GetCoordsXY()
		gridCopy.objects = append(gridCopy.objects, Fossil{X, Y})
	}
	wh.timeline = append(wh.timeline, gridCopy)
}

func (wh *WorldHistory) Get(cycle int) GridHistory {
	return wh.timeline[cycle]
}

type Fossil struct {
	X, Y int
}

type GridHistory struct {
	objects   []Fossil
	locations map[string]Fossil
	Size      int
	cycle     int
}
