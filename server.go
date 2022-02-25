package worldservices

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func findGrid(keyName string, r *http.Request) (Grid, error) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	worldName := m[keyName][0]
	cycle, err := strconv.Atoi(m["cycle"][0])
	if err != nil {
		return Grid{}, err
	}
	return GetWorldBoard(worldName, cycle), nil
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	// TODO make world singleton a service with parameters recieved.
	// have that find the world, and cycle number of the board.
	grid, err := findGrid("world", r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	objects := grid.GetOrderedObjectListByFitness()
	jsonResp, err := json.Marshal(objects)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func resetWorld(w http.ResponseWriter, r *http.Request) {
	var size int
	m, _ := url.ParseQuery(r.URL.RawQuery)
	size, err := strconv.Atoi(m["worldsize"][0])
	if err != nil {
		log.Println("unable to parse size")
	}
	pop, err := strconv.Atoi(m["pop"][0])
	if err != nil {
		log.Println("unable to parse pop")
	}
	world := NewWorld(size, pop)
	WorldSingleton = &world
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func breedWorld(w http.ResponseWriter, r *http.Request) {
	grid1, err := findGrid("world1", r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	grid2, err := findGrid("world2", r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// TODO
	WorldSingleton = world.NewWorldFromDebug()
}

func getCreatureAtCoords(w http.ResponseWriter, r *http.Request) {
	grid := findGrid(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	m, _ := url.ParseQuery(r.URL.RawQuery)
	X, err := strconv.Atoi(m["X"][0])
	if err != nil {
		log.Println("unable to parse x")
	}
	Y, err := strconv.Atoi(m["Y"][0])
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

func SetupServer(port string, staticPath string) {
	http.HandleFunc("/board", getBoard)
	http.HandleFunc("/cycle", cycleWorld)
	http.HandleFunc("/reset", resetWorld)
	http.HandleFunc("/creatures", getCreatureAtCoords)
	http.HandleFunc("/breed", breedWorld)

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
