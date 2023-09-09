package models

type Note struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Body       string `json:"body"`
	Category   string `json:"category"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Notes struct {
	Notes []Note `json:"notes"`
}
