package worldservices

type WorldHistory struct {
	timeline []Grid
}

func (wh WorldHistory) Push(grid Grid) {
	wh.timeline = append(wh.timeline, grid)
}

func (wh WorldHistory) Get(cycle int) Grid {
	return wh.timeline[cycle]
}
