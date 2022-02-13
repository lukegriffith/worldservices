package worldservices

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func getBoard(w http.ResponseWriter, r *http.Request) {
	world, err := GetWorldSingleton()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	objects := world.Grid.GetOrderedObjectListByFitness()
	jsonResp, err := json.Marshal(objects)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func cycleWorld(w http.ResponseWriter, r *http.Request) {
	world, err := GetWorldSingleton()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	world.Cycle()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

func getCreatureAtCoords(w http.ResponseWriter, r *http.Request) {
	world, err := GetWorldSingleton()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	m, _ := url.ParseQuery(r.URL.RawQuery)
	X, err := strconv.Atoi(m["X"][0])
	if err != nil {
		log.Println("unable to parse x")
	}
	Y, err := strconv.Atoi(m["Y"][0])
	objects := world.Grid.GetObjectSenseData(X, Y, 20)
	fmt.Println(objects, X, Y)
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

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
