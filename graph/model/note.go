package model

type Note struct {
	Author *User `json:"author"`
	Date string `json:"date"`
	Text string `json:"text"`
	ItemID string `json:"item_id"`
	Section float64 `json:"section"`
}