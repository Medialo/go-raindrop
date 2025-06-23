package models

import "time"

type Backups struct {
	Result bool `json:"result"`
	Items  []struct {
		Id      string    `json:"_id"`
		Created time.Time `json:"created"`
	} `json:"items"`
}
