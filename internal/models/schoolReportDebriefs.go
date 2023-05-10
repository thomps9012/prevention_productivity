package models

type NewSchoolReportDebrief struct {
	LessonPlanID           string   `json:"lesson_plan_id"`
	StudentCount           int      `json:"student_count"`
	StudentList            []string `json:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
}

type UpdateSchoolReportDebrief struct {
	ID                     string   `json:"id"`
	LessonPlanID           string   `json:"lesson_plan_id"`
	StudentCount           int      `json:"student_count"`
	StudentList            []string `json:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
	Status                 string   `json:"status"`
}

type SchoolReportDebrief struct {
	ID                     string   `json:"id" bson:"_id"`
	UserID                 string   `json:"user_id" bson:"user_id"`
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           int      `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives" bson:"positives"`
	Discussion             string   `json:"discussion" bson:"discussion"`
	Status                 string   `json:"status" bson:"status"`
	CreatedAt              string   `json:"created_at" bson:"created_at"`
	UpdatedAt              string   `json:"updated_at" bson:"updated_at"`
}

type SchoolReportDebriefOverview struct {
	ID            string                         `json:"id"`
	DebriefAuthor []*UserOverview                `json:"debrief_author"`
	LessonPlan    []*SchoolReportPlanDescription `json:"lesson_plan"`
	Status        string                         `json:"status"`
	CreatedAt     string                         `json:"created_at"`
	NoteCount     int                            `json:"note_count"`
}

type SchoolReportDebriefRes struct {
	ID            string                       `json:"id"`
	DebriefAuthor *UserOverview                `json:"debrief_author"`
	LessonPlan    *SchoolReportPlanDescription `json:"lesson_plan"`
	Status        string                       `json:"status"`
	CreatedAt     string                       `json:"created_at"`
}

type SchoolReportDebriefWithNotes struct {
	ID                     string                         `json:"id"`
	DebriefAuthor          []*UserOverview                `json:"debrief_author"`
	LessonPlan             []*SchoolReportPlanDescription `json:"lesson_plan"`
	StudentCount           int                            `json:"student_count"`
	StudentList            []string                       `json:"student_list"`
	ChallengesImprovements string                         `json:"challenges_improvements"`
	Positives              string                         `json:"positives"`
	Discussion             string                         `json:"discussion"`
	Status                 string                         `json:"status"`
	CreatedAt              string                         `json:"created_at"`
	UpdatedAt              string                         `json:"updated_at"`
	Notes                  []*Note                        `json:"notes"`
}
