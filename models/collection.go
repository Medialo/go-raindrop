package models

type Collection struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}
