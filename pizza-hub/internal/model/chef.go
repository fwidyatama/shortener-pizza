package model

type Chef struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsWorking bool   `json:"is_working"`
}
