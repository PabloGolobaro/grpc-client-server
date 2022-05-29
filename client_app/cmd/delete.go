package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pokemonpc "grpc-client-server/proto_app"
)

var deleteCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a pokemon",
	Long:  "Reads a pokemon from MongoDB by given pid",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) > 1 {
		//	log.Fatal("subcommand read only take one argument as pid")
		//}
		err := DeletePokemon(args[0], client)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func DeletePokemon(pid string, client pokemonpc.PokemonServiceClient) error {
	fmt.Println("Deleting Pokemon...")
	deletePokemonRequest := &pokemonpc.DeletePokemonRequest{Pid: pid}
	deletePokemonResponse, err := client.DeletePokemon(context.Background(), deletePokemonRequest)
	if err != nil {
		return err
	}
	fmt.Printf("Succesfuly delete pokemon: %v", deletePokemonResponse)
	return nil
}
