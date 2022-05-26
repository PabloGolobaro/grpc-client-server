package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pokemonpc "grpc-client-server/proto_app"
	"log"
	"os"
)

const defaultPort = "4041"

func main() {
	fmt.Println("Pokemon Client")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn, err := grpc.Dial("localhost:4041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pokemonpc.NewPokemonServiceClient(conn)
	//Создаем покемана
	fmt.Println("Creating Pokemon")
	p := &pokemonpc.Pokemon{
		Pid:         "Poke01",
		Name:        "Pikachu",
		Power:       "Fire",
		Description: "Fluffy",
	}
	createPokemonResponse, err := client.CreatePokemon(context.Background(), &pokemonpc.CreatePokemonRequest{Pokemon: p})
	if err != nil {
		log.Fatal("Unexpected error: %v", err)
	}
	fmt.Printf("Pokemon has been created: %v", createPokemonResponse)

}
