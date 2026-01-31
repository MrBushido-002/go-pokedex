package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
	
)

func (c *Client)GetAreaLocation(inputURL *string) (LocationAreaResponse, error) {
	var response LocationAreaResponse
	URL := baseURL + "/location-area"
	if inputURL != nil {
		URL = *inputURL
	}
	
	resp, ok := c.cache.Get(URL)
	if ok {
		
		if err := json.Unmarshal(resp, &response); err != nil {
			return LocationAreaResponse{}, err
		}
		return response, nil
		
	} else {

		httpResp, err := http.Get(URL)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		defer httpResp.Body.Close()

		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
	
		c.cache.Add(URL, bodyBytes)
		
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			return LocationAreaResponse{}, err
		}

		return response, nil
	
	}

}