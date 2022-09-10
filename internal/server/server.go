package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/lukegriffith/worldservices/internal/world"
)

var (
	WorldKeyName      = "world"
	CycleKeyName      = "cycle"
	World1KeyName     = "world1"
	World2KeyName     = "world2"
	WorldSizeKeyName  = "size"
	PopulationKeyName = "pop"

	SimLength = 512
)

func findGrid(keyName string, r *http.Request) (world.GridHistory, string, error) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	worldName := m[WorldKeyName][0]
	cycle, err := strconv.Atoi(m[CycleKeyName][0])
	if err != nil {
		return world.GridHistory{}, "", err
	}
	return world.GetWorldBoard(worldName, cycle), worldName, nil
}

func WorldServer(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		getWorld(w, r)
	}

	if r.Method == http.MethodPost {
		addToWorldService(w, r)
	}

}

func getWorld(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(world.Worlds))
	for k := range world.Worlds {
		keys = append(keys, k)
	}
	jsonResp, err := json.Marshal(keys)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func addToWorldService(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	worldName := m[WorldKeyName][0]
	fmt.Println(m)
	size, err := strconv.Atoi(m[WorldSizeKeyName][0])
	if err != nil {
		log.Println("unable to parse size")
	}
	pop, err := strconv.Atoi(m[PopulationKeyName][0])
	if err != nil {
		log.Println("unable to parse pop")
	}
	simworld := world.NewWorld(size, pop)
	simworld.Run(SimLength)
	world.RegisterWorld(worldName, simworld)
}

func BoardServer(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		getBoard(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func getBoard(w http.ResponseWriter, r *http.Request) {
	// TODO make world singleton a service with parameters recieved.
	// have that find the world, and cycle number of the board.
	grid, _, err := findGrid(WorldKeyName, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	objects := grid.Objects()
	jsonResp, err := json.Marshal(objects)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

/*

TODO:  This dont work no more.

func breedWorld(w http.ResponseWriter, r *http.Request) {
	grid1, g1Name, err := findGrid(World1KeyName, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	grid2, g2Name, err := findGrid(World2KeyName, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// TODO
	RegisterWorld(fmt.Sprintf("%s-%s", g1Name, g2Name), WorldFromDebugOfWorlds(grid1, grid2))
}

func CreaturesServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getCreatureAtCoords(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

TODO: This also needs re thinking
func getCreatureAtCoords(w http.ResponseWriter, r *http.Request) {
	grid, _, err := findGrid(WorldKeyName, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	m, _ := url.ParseQuery(r.URL.RawQuery)
	X, err := strconv.Atoi(m["X"][0])
	if err != nil {
		log.Println("unable to parse x")
	}
	Y, err := strconv.Atoi(m["Y"][0])
	if err != nil {
		log.Println("unable to parse y")
	}
	objects := grid.GetObjectSenseData(X, Y, 20)
	for _, obj := range objects {
		obj.SetDebug()
	}
	jsonResp, err := json.Marshal(objects)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
*/

func SetupServer(port string, staticPath string) {
	world.NewWorldService()
	http.HandleFunc("/board", BoardServer)
	http.HandleFunc("/world", WorldServer)
	//http.HandleFunc("/creatures", CreaturesServer)
	//http.HandleFunc("/creatures", getCreatureAtCoords)

	//http.HandleFunc("/breed", breedWorld)

	world.Worlds = map[string]*world.World{}

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
