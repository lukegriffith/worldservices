package worldservices

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func getWorldSize(w http.ResponseWriter, r *http.Request) {
	world, err := GetWorldSingleton()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonResp, err := json.Marshal(world.Grid.Size)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func resetWorld(w http.ResponseWriter, r *http.Request) {
	world := NewWorld(100, 30)
	WorldSingleton = &world
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SetupServer(port string, staticPath string) {
	http.HandleFunc("/board", getBoard)
	http.HandleFunc("/cycle", cycleWorld)
	http.HandleFunc("/worldsize", getWorldSize)
	http.HandleFunc("/reset", resetWorld)

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
