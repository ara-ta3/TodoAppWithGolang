package models

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
