package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client)GetAreaLocation(inputUrl string) (LocationAreaResponse, error) {
	var response LocationAreaResponse
	url = baseUrl + "/location-area"
	if inputUrl != nil {
		url = inputUrl
	}
	
	resp, ok := cache.Get(url)
	if ok {
		
		if err := json.Unmarshal(resp, &response); err != nil {
			return LocationAreaResponse{}, err
		}
		return response, nil
		
	} else {

		httpResp, err := http.Get(url)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		defer httpResp.Body.Close()

		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		
		cache.Add(url, bodyBytes)
		
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			return LocationAreaResponse{}, err
		}

		return response, nil
	
	}

}