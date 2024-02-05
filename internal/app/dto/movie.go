package dto

type Movie struct {
	ID         int    `json:"ID"`
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Director   string `json:"Director"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Actors     string `json:"Actors"`
	Writer     string `json:"writer"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	Dvd        string `json:"Dvd"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   bool   `json:"Response"`
}

type MoviesResponse struct {
	Movies   []Movie `json:"Search"`
	Total    string  `json:"totalResults"`
	Response string  `json:"Response"`
}

type CartMovie struct {
	ID     int    `json:"ID"`
	ImdbID string `json:"imdbID"`
}

type AddtoCartRequestBody struct {
	ImdbID string `json:"imdbID"`
}
