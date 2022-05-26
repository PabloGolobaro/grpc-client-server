package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	pokemonpc "grpc-client-server/proto_app"
	"log"
	"net"
	"os"
	"os/signal"
)

const defaultPort = "4041"

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
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	mongo_uri := os.Getenv("MONGODB_URL")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Connecting to MongoDB...")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pokemon Service started")
	collection = client.Database("pokemonDB").Collection("pokemon")
	listen, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	newServer := grpc.NewServer(opts...)
	pokemonpc.RegisterPokemonServiceServer(newServer, &server{})
	reflection.Register(newServer)

	go func() {
		fmt.Println("Starting Server...")
		if err := newServer.Serve(listen); err != nil {
			log.Fatal("Failed to serve : %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Closing MongoDB Collection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error on disconection with MongoDB : %V", err)
	}
	fmt.Println("Stopping the server")
	newServer.Stop()
	log.Println("End of programm")
}
