package worldservices

var (
	Worlds         map[string]World
	WorldSingleton *World
)

func GetWorldBoard(name string, cycle int) Grid {
	grid := Worlds[name].History.Get(cycle)
	return grid
}

func GetWorld(name string) World {
	return Worlds[name]
}
