package models

import "time"

// RaindropCreate represents the request body for creating a raindrop
type RaindropCreate struct {
	PleaseParse bool   `json:"pleaseParse,omitempty"`
	Created     string `json:"created,omitempty"`
	//LastUpdate  string            `json:"lastUpdate,omitempty"` // DO NOT WORK ON API
	//Order       int               `json:"order,omitempty"`      // DO NOT WORK ON API
	Important    bool        `json:"important,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	Media        []Media     `json:"media,omitempty"`
	Cover        string      `json:"cover,omitempty"`
	Collection   *Collection `json:"collection,omitempty"`
	CollectionId int         `json:"collectionId,omitempty"`
	//Type         string            `json:"type,omitempty"`
	Excerpt    string            `json:"excerpt,omitempty"`
	Title      string            `json:"title,omitempty"`
	Link       string            `json:"link"`
	Highlights []HighlightCreate `json:"highlights,omitempty"`
	//Reminder    *Reminder   `json:"reminder,omitempty"`
}

type RaindropUpdate struct {
	Created string `json:"created,omitempty"`
	//LastUpdate   string   `json:"lastUpdate,omitempty"` // DO NOT WORK ON API
	//Order        int      `json:"order,omitempty"`      // DO NOT WORK ON API
	Important    bool     `json:"important,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Media        []Media  `json:"media,omitempty"`
	Cover        string   `json:"cover,omitempty"`
	CollectionId int      `json:"collectionId,omitempty"`
	//Type         string   `json:"type,omitempty"`       // USELESS
	Title      string            `json:"title,omitempty"`
	Excerpt    string            `json:"excerpt,omitempty"`
	Note       string            `json:"note,omitempty"`
	Link       string            `json:"link,omitempty"`
	Highlights []HighlightUpdate `json:"highlights,omitempty"`
	//Reminder	*Reminder   `json:"reminder,omitempty"`   // PRO only
}

// RaindropResponse represent api response of raindrop.io
type RaindropResponse struct {
	Id      int    `json:"_id,omitempty"`
	Link    string `json:"link,omitempty"`
	Title   string `json:"title,omitempty"`
	Excerpt string `json:"excerpt,omitempty"`
	Note    string `json:"note,omitempty"`
	Type    string `json:"type,omitempty"`
	User    struct {
		Ref string `json:"$ref,omitempty"`
		Id  int    `json:"$id,omitempty"`
	} `json:"user,omitempty"`
	Cover string `json:"cover,omitempty"`
	Media []struct {
		Link string `json:"link,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"media,omitempty"`
	Tags      []interface{} `json:"tags,omitempty"`
	Important bool          `json:"important,omitempty"`
	Reminder  struct {
		Date interface{} `json:"date,omitempty"`
	} `json:"reminder,omitempty"`
	Removed    bool      `json:"removed,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	Collection struct {
		Ref string `json:"$ref,omitempty"`
		Id  int    `json:"$id,omitempty"`
		Oid int    `json:"oid,omitempty"`
	} `json:"collection,omitempty"`
	Highlights *[]HighlightResponse `json:"highlights,omitempty"`
	LastUpdate time.Time            `json:"lastUpdate,omitempty"`
	Domain     string               `json:"domain,omitempty"`
	CreatorRef struct {
		Id     int    `json:"_id,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Name   string `json:"name,omitempty"`
		Email  string `json:"email,omitempty"`
	} `json:"creatorRef,omitempty"`
	Sort         int `json:"sort,omitempty"`
	CollectionId int `json:"collectionId,omitempty"`
}
