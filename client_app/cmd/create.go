package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pokemonpc "grpc-client-server/proto_app"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a pokemon",
	Long:  "Creates a new Pokemon in MongoDB from given data",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) <= 4 {
		//	log.Fatal("subcommand create-pokemon only take four arguments\"")
		//}
		pokemonItem := DataToPokemonItem(args)
		err := CreatePokemon(pokemonItem, client)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func CreatePokemon(pokemon *PokemonItem, client pokemonpc.PokemonServiceClient) error {
	fmt.Println("Creating Pokemon...")
	p := getPokemonData(pokemon)
	createPokemonResponse, err := client.CreatePokemon(context.Background(), &pokemonpc.CreatePokemonRequest{Pokemon: p})
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon has been created: %v", createPokemonResponse)
	return nil
}
