package schema

type GetJokeResponseValue struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

type GetJokeResponse struct {
	Type  string               `json:"type"`
	Value GetJokeResponseValue `json:"value"`
}
