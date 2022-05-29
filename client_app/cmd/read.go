package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pokemonpc "grpc-client-server/proto_app"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a pokemon",
	Long:  "Reads a pokemon from MongoDB by given pid",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := ReadPokemon(args[0], client)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func ReadPokemon(pid string, client pokemonpc.PokemonServiceClient) error {
	fmt.Println("Reading Pokemon...")
	readPokemonRequest := &pokemonpc.ReadPokemonRequest{Pid: pid}
	readPokemonResponse, err := client.ReadPokemon(context.Background(), readPokemonRequest)
	if err != nil {
		return err
	}
	fmt.Printf("Succesfuly read pokemon: %v", readPokemonResponse)
	return nil
}
