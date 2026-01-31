package pokeapi

type LocationAreaResponse struct {
	Count int
	Next string
	Previous string
	Results []struct {
		Name string
		url string
	}
}