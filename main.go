package main

import (
	"time"

	"github.com/hackastak/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
