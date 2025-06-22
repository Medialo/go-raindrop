package models

type MediaType string

const (
	MediaTypeImage MediaType = "image"
)

type Media struct {
	Link string    `json:"link,omitempty"`
	Type MediaType `json:"type,omitempty"`
}
