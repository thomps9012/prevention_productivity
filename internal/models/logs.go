package models

type NewLog struct {
	DailyActivity string `json:"daily_activity"`
	Positives     string `json:"positives"`
	Improvements  string `json:"improvements"`
	NextSteps     string `json:"next_steps"`
}

type UpdateLog struct {
	ID            string `json:"id"`
	DailyActivity string `json:"daily_activity"`
	Positives     string `json:"positives"`
	Improvements  string `json:"improvements"`
	NextSteps     string `json:"next_steps"`
	Status        string `json:"status"`
}

type Log struct {
	ID            string `json:"id" bson:"_id"`
	UserID        string `json:"user_id" bson:"user_id"`
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives" bson:"positives"`
	Improvements  string `json:"improvements" bson:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
	Status        string `json:"status" bson:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
}

type LogOverview struct {
	ID        string          `json:"id"`
	LogAuthor []*UserOverview `json:"log_author"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Status    string          `json:"status"`
	NoteCount int             `json:"note_count"`
}

type LogRes struct {
	ID        string        `json:"id"`
	LogAuthor *UserOverview `json:"log_author"`
	Status    string        `json:"status"`
	CreatedAt string        `json:"created_at"`
}

type LogWithNotes struct {
	ID            string          `json:"id"`
	LogAuthor     []*UserOverview `json:"log_author"`
	DailyActivity string          `json:"daily_activity"`
	Positives     string          `json:"positives"`
	Improvements  string          `json:"improvements"`
	NextSteps     string          `json:"next_steps"`
	Status        string          `json:"status"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	Notes         []*Note         `json:"notes"`
}
