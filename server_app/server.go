package server_app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pokemonpc "grpc-client-server/proto_app"
	"log"
)

var collection *mongo.Collection

type server struct {
	pokemonpc.PokemonServiceServer
}
type pokemonItem struct {
	ID          primitive.ObjectID `bson:"id"`
	Pid         string             `bson:"pid"`
	Name        string             `bson:"name"`
	Power       string             `bson:"power"`
	Description string             `bson:"description"`
}

func getPokemonData(data *pokemonItem) *pokemonpc.Pokemon {
	return &pokemonpc.Pokemon{
		Id:          data.ID.Hex(),
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}

}

func (s *server) CreatePokemon(ctx context.Context, request *pokemonpc.CreatePokemonRequest) (*pokemonpc.CreatePokemonResponse, error) {
	log.Println("Create Pokemon...")
	pokemon := request.GetPokemon()
	data := pokemonItem{
		Pid:         pokemon.GetPid(),
		Name:        pokemon.GetName(),
		Power:       pokemon.GetPower(),
		Description: pokemon.GetDescription(),
	}
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID"))
	}
	return &pokemonpc.CreatePokemonResponse{
		Pokemon: &pokemonpc.Pokemon{
			Id:          oid.Hex(),
			Pid:         pokemon.GetPid(),
			Name:        pokemon.GetName(),
			Power:       pokemon.GetPower(),
			Description: pokemon.GetDescription(),
		},
	}, nil
}
