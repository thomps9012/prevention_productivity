package models

type NewContact struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Notes *string `json:"notes"`
}

type UpdateContact struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Notes *string `json:"notes"`
}

type Contact struct {
	ID        string  `json:"id" bson:"_id"`
	Type      string  `json:"type" bson:"type"`
	Name      string  `json:"name" bson:"name"`
	Email     *string `json:"email" bson:"email"`
	Phone     *string `json:"phone" bson:"phone"`
	Notes     *string `json:"notes" bson:"notes"`
	Active    bool    `json:"active" bson:"active"`
	CreatedBy string  `json:"created_by" bson:"created_by"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt string  `json:"deleted_at" bson:"deleted_at"`
}

type ContactDetail struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Email     *string         `json:"email"`
	Phone     *string         `json:"phone"`
	Notes     *string         `json:"notes"`
	Active    bool            `json:"active"`
	CreatedBy []*UserOverview `json:"created_by"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	DeletedAt string          `json:"deleted_at"`
}

type ContactOverview struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
