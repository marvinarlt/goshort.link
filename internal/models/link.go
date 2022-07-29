package models

import "time"

type Link struct {
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"-"`
	LatestUse time.Time `json:"-"`
}
