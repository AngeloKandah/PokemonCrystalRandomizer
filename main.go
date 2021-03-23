package main

import (
	"fmt"

	pokemon "github.com/AngeloKandah/pokemon/pokemoninfo"
	"github.com/AngeloKandah/pokemon/utils"
	goxymemmory "github.com/Xustyx/goxymemory"
)

var search_addr string
var final_address uint
var prefix uint

func init() {
	search_addr = utils.GetSearchAddress()
	dm := goxymemmory.DataManager("mGBA.exe")
	if !dm.IsOpen {
		fmt.Printf("Failed opening process.\n")
		return
	}
	start := 0x0000147D //Offset for OT_Name in memory
	for i := 0x0; ; i += 0x00010000 {
		current_address := start + i
		data, err := dm.Read(uint(current_address), goxymemmory.UINT)
		if err != nil {
			fmt.Printf("Failed reading memory. %s", err)
		}
		temp_address := fmt.Sprintf("%X", data.Value)
		if temp_address == search_addr { //Comparing values of name to values at the temp address to find the 4 bytes to address
			fmt.Printf("ADDRESS IS %X \n", current_address)
			final_address = uint(current_address)
			break
		} else {
			final_address = 0x00000000
		}
	}
	prefix = final_address ^ uint(start)
}

func main() {
	dm := goxymemmory.DataManager("mGBA.exe")
	if !dm.IsOpen {
		fmt.Printf("Failed opening process.\n")
		return
	}
	party_size_address := prefix ^ 0x00001CD7
	data, err := dm.Read(uint(party_size_address), goxymemmory.BYTE)
	if err != nil {
		fmt.Printf("Failed reading memory. %s", err)
	}
	party_size := data.Value.(uint8)

	party_start := uint(0x00001CDF)
	const sizeOfPokemon int = 48
	party_start = prefix ^ party_start
	party_pokemon := make([]pokemon.PokemonInfo, party_size)
	//var pokeClone [68]uint8

	for party := 0; party < int(party_size); party++ {
		pokeData := make([]byte, 68)
		party_species := 0x00001CD8 + uint(party)
		OT_address := 0x00001DFF + uint(party*11)
		nickname_address := 0x00001E41 + uint(party*11)
		for i := 0; i < 48; i++ {
			slot := party_start + uint(i*0x01) + uint(sizeOfPokemon*party)
			data, err := dm.Read(slot, goxymemmory.BYTE)
			if err != nil {
				fmt.Printf("Failed reading memory. %s", err)
			}
			pokeData[i] = data.Value.(byte)
			//pokeClone[i] = data.Value.(uint8)
		}
		data, err := dm.Read(prefix^party_species, goxymemmory.BYTE)
		if err != nil {
			fmt.Printf("Failed reading memory. %s", err)
		}
		pokeData[48] = data.Value.(byte)
		for j := 0; j < 8; j++ {
			data, err := dm.Read((prefix^OT_address)+uint(j), goxymemmory.BYTE)
			if err != nil {
				fmt.Printf("Failed reading memory. %s", err)
			}
			pokeData[j+49] = data.Value.(byte)
		}
		for j := 0; j < 11; j++ {
			data, err := dm.Read((prefix^nickname_address)+uint(j), goxymemmory.BYTE)
			if err != nil {
				fmt.Printf("Failed reading memory. %s", err)
			}
			pokeData[j+57] = data.Value.(byte)
		}
		pokemon, err := pokemon.CreatePokemonInfo(pokeData)
		if err != nil {
			continue
		}
		pokemon.Gender = utils.GetGender(pokemon.DEX_NUM, pokemon.IV_ATK)
		pokemon.Types = utils.GetTypes(pokemon.DEX_NUM)
		party_pokemon[party] = pokemon
	}
	for i := 0; i < len(party_pokemon); i++ {
		fmt.Println(party_pokemon[i])
	}
}
