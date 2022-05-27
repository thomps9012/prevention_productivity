package model

type Note struct {
	Author *User `json:"author"`
	Date time.Time `bson:"date"`
	Text string `bson:"text"`
	Section float64 `bson:"section"`
}