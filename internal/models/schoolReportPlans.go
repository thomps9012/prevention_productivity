package models

type NewSchoolReportPlan struct {
	Date           string    `json:"date"`
	CoFacilitators []*string `json:"co_facilitators"`
	Curriculum     string    `json:"curriculum"`
	School         string    `json:"school"`
	LessonTopics   string    `json:"lesson_topics"`
}

type UpdateSchoolReportPlan struct {
	ID             string    `json:"id"`
	CoFacilitators []*string `json:"co_facilitators"`
	Curriculum     string    `json:"curriculum"`
	LessonTopics   string    `json:"lesson_topics"`
	School         string    `json:"school"`
	Status         string    `json:"status"`
}

type SchoolReportPlan struct {
	ID              string    `json:"id" bson:"_id"`
	UserID          *string   `json:"user_id" bson:"user_id"`
	Date            string    `json:"date" bson:"date"`
	Co_Facilitators []*string `json:"co_facilitators" bson:"co_facilitators"`
	Curriculum      string    `json:"curriculum" bson:"curriculum"`
	School          string    `json:"school" bson:"school"`
	LessonTopics    string    `json:"lesson_topics" bson:"lesson_topics"`
	Status          string    `json:"status" bson:"status"`
	CreatedAt       string    `json:"created_at" bson:"created_at"`
	UpdatedAt       string    `json:"updated_at" bson:"updated_at"`
}

type SchoolReportPlanDescription struct {
	ID     string `json:"id"`
	School string `json:"school"`
	Date   string `json:"date"`
}

type SchoolReportPlanOverview struct {
	ID           string          `json:"id"`
	Date         string          `json:"date"`
	ReportAuthor []*UserOverview `json:"report_author"`
	School       string          `json:"school"`
	Status       string          `json:"status"`
	CreatedAt    string          `json:"created_at"`
	NoteCount    int             `json:"note_count"`
}

type SchoolReportPlanRes struct {
	ID         string        `json:"id"`
	PlanAuthor *UserOverview `json:"plan_author"`
	Date       string        `json:"date"`
	School     string        `json:"school"`
	Status     string        `json:"status"`
	CreatedAt  string        `json:"created_at"`
}

type SchoolReportPlanWithNotes struct {
	ID             string          `json:"id"`
	Date           string          `json:"date"`
	ReportAuthor   []*UserOverview `json:"report_author"`
	CoFacilitators []*UserOverview `json:"co_facilitators"`
	Curriculum     string          `json:"curriculum"`
	School         string          `json:"school"`
	LessonTopics   string          `json:"lesson_topics"`
	Status         string          `json:"status"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
}
