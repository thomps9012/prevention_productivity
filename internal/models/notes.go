package models

type NewNote struct {
	ItemID  string `json:"item_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNote struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Note struct {
	ID        string `json:"id" bson:"_id"`
	ItemID    string `json:"item_id" bson:"item_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	Title     string `json:"title" bson:"title"`
	Content   string `json:"content" bson:"content"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

type ItemOverview struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type NoteDetail struct {
	ID        string        `json:"id"`
	ItemInfo  *ItemOverview `json:"item_info"`
	Author    *UserOverview `json:"author"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}
