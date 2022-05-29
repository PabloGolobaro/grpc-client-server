package cmd

import pokemonpc "grpc-client-server/proto_app"

type PokemonItem struct {
	Pid         string
	Name        string
	Power       string
	Description string
}

func getPokemonData(data *PokemonItem) *pokemonpc.Pokemon {
	return &pokemonpc.Pokemon{
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}

}
func DataToPokemonItem(data []string) *PokemonItem {
	return &PokemonItem{
		Pid:         data[0],
		Name:        data[1],
		Power:       data[2],
		Description: data[3],
	}
}
