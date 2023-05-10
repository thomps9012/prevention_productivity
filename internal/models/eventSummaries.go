package models

type NewEventSummary struct {
	EventID       string    `json:"event_id"`
	CoPlanners    []*string `json:"co_planners"`
	AttendeeCount int       `json:"attendee_count"`
	Challenges    string    `json:"challenges"`
	Successes     string    `json:"successes"`
	Improvements  string    `json:"improvements"`
}

type UpdateEventSummary struct {
	ID            string    `json:"id"`
	EventID       string    `json:"event_id"`
	CoPlanners    []*string `json:"co_planners"`
	AttendeeCount int       `json:"attendee_count"`
	Challenges    string    `json:"challenges"`
	Successes     string    `json:"successes"`
	Improvements  string    `json:"improvements"`
	Status        string    `json:"status"`
}

type EventSummary struct {
	ID            string    `json:"id" bson:"_id"`
	EventID       string    `json:"event_id" bson:"event_id"`
	UserID        string    `json:"user_id" bson:"user_id"`
	Co_Planners   []*string `json:"co_planners" bson:"co_planners"`
	AttendeeCount int       `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges" bson:"challenges"`
	Successes     string    `json:"successes" bson:"successes"`
	Improvements  string    `json:"improvements" bson:"improvements"`
	Status        string    `json:"status" bson:"status"`
	CreatedAt     string    `json:"created_at" bson:"created_at"`
	UpdatedAt     string    `json:"updated_at" bson:"updated_at"`
}

type EventSummaryOverview struct {
	ID               string              `json:"id"`
	EventDescription []*EventDescription `json:"event_description"`
	SummaryAuthor    []*UserOverview     `json:"summary_author"`
	AttendeeCount    int                 `json:"attendee_count"`
	Status           string              `json:"status"`
	CreatedAt        string              `json:"created_at"`
	NoteCount        int                 `json:"note_count"`
}

type EventSummaryRes struct {
	ID            string            `json:"id"`
	Event         *EventDescription `json:"event"`
	SummaryAuthor *UserOverview     `json:"summary_author"`
	Status        string            `json:"status"`
	CreatedAt     string            `json:"created_at"`
}

type EventSummaryWithNotes struct {
	ID               string              `json:"id"`
	EventDescription []*EventDescription `json:"event_description"`
	SummaryAuthor    []*UserOverview     `json:"summary_author"`
	CoPlanners       []*UserOverview     `json:"co_planners"`
	AttendeeCount    int                 `json:"attendee_count"`
	Challenges       string              `json:"challenges"`
	Successes        string              `json:"successes"`
	Improvements     string              `json:"improvements"`
	Status           string              `json:"status"`
	CreatedAt        string              `json:"created_at"`
	UpdatedAt        string              `json:"updated_at"`
	Notes            []*Note             `json:"notes"`
}
