package pokeapi

import(
	"encoding/json"
	"io"
)

func (c *Client)GetLocation(locationName string) (Location, error) {
	var response Location
	URL := "https://pokeapi.co/api/v2/location-area/" + locationName
	
	resp, ok := c.cache.Get(URL)
	if ok {
		if err := json.Unmarshal(resp, &response); err != nil {
			return Location{}, err
		}
		return response, nil
		
	}

	httpResp, err := c.httpClient.Get(URL)
	if err != nil {
		return Location{}, err
	}

	defer httpResp.Body.Close()

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(URL, bodyBytes)
	
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return Location{}, err
	}

	return response, nil



}