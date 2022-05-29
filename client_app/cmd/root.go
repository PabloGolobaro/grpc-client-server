package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pokemonpc "grpc-client-server/proto_app"
	"log"
	"os"
)

const defaultPort = "4041"

var client pokemonpc.PokemonServiceClient

var rootCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Pokemon App for CRUD",
	Long:  "Pokemon app for implementing CRUD operations on MongoDB",
}

func init() {
	rootCmd.AddCommand(createCmd, readCmd, updateCmd, deleteCmd)
}
func Execute() {
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
	client = pokemonpc.NewPokemonServiceClient(conn)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
