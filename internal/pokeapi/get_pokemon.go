package pokeapi

import(
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonInfo, error) {
	var response PokemonInfo

	URL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	httpResp, err := c.httpClient.Get(URL)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer httpResp.Body.Close()

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return PokemonInfo{}, err
	}
	return response, nil

}