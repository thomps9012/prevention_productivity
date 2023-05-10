package models

type NewGrant struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Goals       []string `json:"goals"`
	Objectives  []string `json:"objectives"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Budget      float64  `json:"budget"`
	AwardNumber string   `json:"award_number"`
	AwardDate   string   `json:"award_date"`
}

type UpdateGrant struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Goals       []string `json:"goals"`
	Objectives  []string `json:"objectives"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Budget      float64  `json:"budget"`
	AwardNumber string   `json:"award_number"`
	AwardDate   string   `json:"award_date"`
	Active      bool     `json:"active"`
}

type Grant struct {
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Goals       []string `json:"goals" bson:"goals"`
	Objectives  []string `json:"objectives" bson:"objectives"`
	StartDate   string   `json:"start_date" bson:"start_date"`
	AwardDate   string   `json:"award_date" bson:"award_date"`
	EndDate     string   `json:"end_date" bson:"end_date"`
	AwardNumber string   `json:"award_number" bson:"award_number"`
	Budget      float64  `json:"budget" bson:"budget"`
	CreatedBy   string   `json:"created_by" bson:"created_by"`
	CreatedAt   string   `json:"created_at" bson:"created_at"`
	UpdatedAt   string   `json:"updated_at" bson:"updated_at"`
	Active      bool
}

type GrantDetail struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Goals       []string        `json:"goals"`
	Objectives  []string        `json:"objectives"`
	StartDate   string          `json:"start_date"`
	AwardDate   string          `json:"award_date"`
	EndDate     string          `json:"end_date"`
	AwardNumber string          `json:"award_number"`
	Budget      float64         `json:"budget"`
	Active      bool            `json:"active"`
	CreatedBy   []*UserOverview `json:"created_by"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

type GrantOverview struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	StartDate   string  `json:"start_date"`
	AwardDate   string  `json:"award_date"`
	EndDate     string  `json:"end_date"`
	AwardNumber string  `json:"award_number"`
	Budget      float64 `json:"budget"`
	Active      bool    `json:"active"`
}
