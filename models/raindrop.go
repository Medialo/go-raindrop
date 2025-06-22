package models

import "time"

// RaindropRequest represents the request body for creating a raindrop
type RaindropRequest struct {
	PleaseParse bool        `json:"pleaseParse,omitempty"`
	Created     string      `json:"created,omitempty"`
	LastUpdate  string      `json:"lastUpdate,omitempty"`
	Order       int         `json:"order,omitempty"`
	Important   bool        `json:"important,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	Media       []string    `json:"media,omitempty"`
	Cover       string      `json:"cover,omitempty"`
	Collection  *Collection `json:"collection,omitempty"`
	Type        string      `json:"type,omitempty"`
	Excerpt     string      `json:"excerpt,omitempty"`
	Title       string      `json:"title,omitempty"`
	Link        string      `json:"link"`
	Highlights  []string    `json:"highlights,omitempty"`
	//Reminder    *Reminder   `json:"reminder,omitempty"`
}

// RaindropResponse représente un favori dans RaindropResponse.io.
type RaindropResponse struct {
	Id      int    `json:"_id"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Note    string `json:"note"`
	Type    string `json:"type"`
	User    struct {
		Ref string `json:"$ref"`
		Id  int    `json:"$id"`
	} `json:"user"`
	Cover string `json:"cover"`
	Media []struct {
		Link string `json:"link"`
		Type string `json:"type"`
	} `json:"media"`
	Tags      []interface{} `json:"tags"`
	Important bool          `json:"important"`
	Reminder  struct {
		Date interface{} `json:"date"`
	} `json:"reminder"`
	Removed    bool      `json:"removed"`
	Created    time.Time `json:"created"`
	Collection struct {
		Ref string `json:"$ref"`
		Id  int    `json:"$id"`
		Oid int    `json:"oid"`
	} `json:"collection"`
	Highlights []interface{} `json:"highlights"`
	LastUpdate time.Time     `json:"lastUpdate"`
	Domain     string        `json:"domain"`
	CreatorRef struct {
		Id     int    `json:"_id"`
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	} `json:"creatorRef"`
	Sort         int `json:"sort"`
	CollectionId int `json:"collectionId"`
}
