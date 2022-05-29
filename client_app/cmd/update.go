package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pokemonpc "grpc-client-server/proto_app"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates a pokemon",
	Long:  "Updates the Pokemon in MongoDB with given data",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		pokemonItem := DataToPokemonItem(args)
		err := UpdatePokemon(pokemonItem, client)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func UpdatePokemon(pokemon *PokemonItem, client pokemonpc.PokemonServiceClient) error {
	fmt.Println("Updating Pokemon...")
	p := getPokemonData(pokemon)
	UpdatePokemonResponce, err := client.UpdatePokemon(context.Background(), &pokemonpc.UpdatePokemonRequest{Pokemon: p})
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon has been updated: %v", UpdatePokemonResponce)
	return nil
}
