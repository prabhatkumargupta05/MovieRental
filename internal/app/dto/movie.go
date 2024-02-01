package dto

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MoviesResponse struct {
	Movies   []Movie `json:"Search"`
	Total    string  `json:"totalResults"`
	Response string  `json:"Response"`
}
