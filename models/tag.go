package models

// Tag représente un tag dans RaindropResponse.io.
type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
