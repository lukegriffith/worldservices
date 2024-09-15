package server

import (
	"context"
	"errors"

	"github.com/lukegriffith/worldservices/internal/world"
	pb "github.com/lukegriffith/worldservices/pkg/proto"
)

type WorldServer struct {
	pb.UnimplementedWorldServiceServer
	Worlds world.WorldsContainer
}

func NewServer() *WorldServer {
	w := WorldServer{
		Worlds: world.WorldsContainer{
			Worlds:      []world.WorldsObject{},
			ActiveWorld: 0,
		},
	}
	return &w
}

func (s *WorldServer) CreateWorld(ctx context.Context, w *pb.World) (*pb.WorldResponse, error) {
	id := s.Worlds.NewWorld(w.Name)
	resp := &pb.WorldResponse{
		World: &pb.World{
			Id:   id,
			Name: w.Name,
		},
		Error: "",
	}
	return resp, nil
}

// Sets a singletons value to the selected world
func (s *WorldServer) SelectWorld(ctx context.Context, w *pb.WorldSelectionRequest) (*pb.WorldResponse, error) {
	if len(s.Worlds.Worlds) > int(w.Id) {
		s.Worlds.ActiveWorld = w.Id
		world := s.Worlds.Worlds[w.Id]
		resp := &pb.WorldResponse{
			World: &pb.World{
				Id:   world.Id,
				Name: world.Name,
			},
			Error: "",
		}
		return resp, nil
	}
	return &pb.WorldResponse{
		World: nil,
		Error: "Unable to find index",
	}, errors.New("Unable to find index")

}

// Returns a list of all created worlds
// If an ID is provided, it returns a single world
func (s *WorldServer) GetWorld(ctx context.Context, w *pb.WorldRequest) (*pb.WorldResponse, error) {
	if len(s.Worlds.Worlds) > int(w.Id) {
		world := s.Worlds.Worlds[w.Id]
		resp := &pb.WorldResponse{
			World: &pb.World{
				Id:   world.Id,
				Name: world.Name,
			},
			Error: "",
		}
		return resp, nil
	}
	return &pb.WorldResponse{
		World: nil,
		Error: "Unable to find index",
	}, errors.New("Unable to find index")

}

// Sets a singletons value to the selected world

// Returns a list of all created worlds
// If an ID is provided, it returns a single world
