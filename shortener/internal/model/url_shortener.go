package model

import "time"

type URLShortener struct {
	Destination  string    `json:"destination"`
	ShortURL     string    `json:"short_url"`
	ClickCounter int       `json:"click_counter"`
	ExpireAt     time.Time `json:"expire_at"`
}
