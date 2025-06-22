package models

// Collection représente une collection dans RaindropResponse.io.
type Collection struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}
