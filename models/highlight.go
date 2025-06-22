package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type HighlightCreate struct {
	Text     string    `json:"text,omitempty"`
	Note     string    `json:"note,omitempty"`
	Color    string    `json:"color,omitempty"`
	Position int       `json:"position,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

type HighlightUpdate struct {
	Text     string    `json:"text,omitempty"`
	Note     string    `json:"note,omitempty"`
	Color    string    `json:"color,omitempty"`
	Position int       `json:"position,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Id       string    `json:"_id,omitempty"`
}

type HighlightResponse struct {
	Text       string            `json:"text,omitempty"`
	Note       string            `json:"note,omitempty"`
	Color      string            `json:"color,omitempty"`
	Position   int               `json:"position,omitempty"`
	Created    time.Time         `json:"created,omitempty"`
	LastUpdate time.Time         `json:"lastUpdate,omitempty"`
	CreatorRef CreatorRefWrapper `json:"creatorRef,omitempty"`
	Id         string            `json:"_id,omitempty"`
}

type CreatorRef struct {
	Id     int    `json:"_id,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

type CreatorRefWrapper struct {
	Full *CreatorRef
	ID   *int
}

func (w *CreatorRefWrapper) UnmarshalJSON(data []byte) error {
	// Try if creatorRef is an object
	var full CreatorRef
	if err := json.Unmarshal(data, &full); err == nil && (full.Id != 0 || full.Name != "") {
		w.Full = &full
		return nil
	}

	// Try if CreatorRef is a number
	var id int
	if err := json.Unmarshal(data, &id); err == nil {
		w.ID = &id
		return nil
	}

	return fmt.Errorf("unexpected creatorRef format: %s", string(data))
}
