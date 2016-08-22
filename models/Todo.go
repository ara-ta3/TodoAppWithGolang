package models

import "encoding/json"

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

func (t *Todo) ToJson() ([]byte, error) {
	return json.Marshal(t)
}
