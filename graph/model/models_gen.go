// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type UserResult interface {
	IsUserResult()
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
	ID        string        `json:"id" bson:"_id"`
	Type      string        `json:"type" bson:"type"`
	Name      string        `json:"name" bson:"name"`
	Email     *string       `json:"email" bson:"email"`
	Phone     *string       `json:"phone" bson:"phone"`
	Notes     *string       `json:"notes" bson:"notes"`
	Active    bool          `json:"active" bson:"active"`
	CreatedBy *UserOverview `json:"created_by" bson:"created_by"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	UpdatedAt string        `json:"updated_at" bson:"updated_at"`
	DeletedAt string        `json:"deleted_at" bson:"deleted_at"`
}

type ContactOverview struct {
	ID     string `json:"id" bson:"_id"`
	Type   string `json:"type" bson:"type"`
	Name   string `json:"name" bson:"name"`
	Active bool   `json:"active" bson:"active"`
}

type Event struct {
	ID                      string    `json:"id" bson:"_id"`
	UserID                  string    `json:"user_id" bson:"user_id"`
	CoPlanners              []*string `json:"co_planners" bson:"co_planners"`
	Title                   string    `json:"title" bson:"title"`
	Description             string    `json:"description" bson:"description"`
	StartDate               string    `json:"start_date" bson:"start_date"`
	SetUp                   string    `json:"set_up" bson:"set_up"`
	CleanUp                 string    `json:"clean_up" bson:"clean_up"`
	EndDate                 string    `json:"end_date" bson:"end_date"`
	GrantID                 string    `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool      `json:"public_event" bson:"public_event"`
	RsvpRequired            bool      `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event" bson:"annual_event"`
	NewEvent                bool      `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []string  `json:"agenda" bson:"agenda"`
	TargetAudience          string    `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial       []*string `json:"marketing_material" bson:"marketing_material"`
	Supplies                []*string `json:"supplies" bson:"supplies"`
	SpecialOrders           []*string `json:"special_orders" bson:"special_orders"`
	Performance             string    `json:"performance" bson:"performance"`
	Vendors                 string    `json:"vendors" bson:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                 string    `json:"caterer" bson:"caterer"`
	FoodHeadCount           int       `json:"food_head_count" bson:"food_head_count"`
	EventTeam               []*string `json:"event_team" bson:"event_team"`
	VolunteerList           []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                  float64   `json:"budget" bson:"budget"`
	AffiliatedOrganizations []*string `json:"affiliated_organizations" bson:"affiliated_organizations"`
	EducationalGoals        []string  `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes     []string  `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals              []string  `json:"grant_goals" bson:"grant_goals"`
	CreatedAt               string    `json:"created_at" bson:"created_at"`
	UpdatedAt               string    `json:"updated_at" bson:"updated_at"`
	Status                  string    `json:"status" bson:"status"`
}

type EventDescription struct {
	ID        string `json:"id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	StartDate string `json:"start_date" bson:"start_date"`
}

type EventOverview struct {
	ID        string        `json:"id" bson:"_id"`
	EventLead *UserOverview `json:"event_lead" bson:"event_lead"`
	Title     string        `json:"title" bson:"title"`
	StartDate string        `json:"start_date" bson:"start_date"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	Status    string        `json:"status" bson:"status"`
	NoteCount int           `json:"note_count" bson:"note_count"`
}

type EventRes struct {
	ID        string        `json:"id" bson:"_id"`
	EventLead *UserOverview `json:"event_lead" bson:"event_lead"`
	Title     string        `json:"title" bson:"title"`
	StartDate string        `json:"start_date" bson:"start_date"`
	Status    string        `json:"status" bson:"status"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
}

type EventSummary struct {
	ID            string    `json:"id" bson:"_id"`
	EventID       string    `json:"event_id" bson:"event_id"`
	UserID        string    `json:"user_id" bson:"user_id"`
	CoPlanners    []*string `json:"co_planners" bson:"co_planners"`
	AttendeeCount int       `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges" bson:"challenges"`
	Successes     string    `json:"successes" bson:"successes"`
	Improvements  string    `json:"improvements" bson:"improvements"`
	Status        string    `json:"status" bson:"status"`
	CreatedAt     string    `json:"created_at" bson:"created_at"`
	UpdatedAt     string    `json:"updated_at" bson:"updated_at"`
}

type EventSummaryOverview struct {
	ID               string            `json:"id" bson:"_id"`
	EventDescription *EventDescription `json:"event_description" bson:"event_description"`
	SummaryAuthor    *UserOverview     `json:"summary_author" bson:"summary_author"`
	AttendeeCount    int               `json:"attendee_count" bson:"attendee_count"`
	Status           string            `json:"status" bson:"status"`
	CreatedAt        string            `json:"created_at" bson:"created_at"`
	NoteCount        int               `json:"note_count" bson:"note_count"`
}

type EventSummaryRes struct {
	ID            string            `json:"id" bson:"_id"`
	Event         *EventDescription `json:"event" bson:"event"`
	SummaryAuthor *UserOverview     `json:"summary_author" bson:"summary_author"`
	Status        string            `json:"status" bson:"status"`
	CreatedAt     string            `json:"created_at" bson:"created_at"`
}

type EventSummaryWithNotes struct {
	ID               string            `json:"id" bson:"_id"`
	EventDescription *EventDescription `json:"event_description" bson:"event_description"`
	SummaryAuthor    *UserOverview     `json:"summary_author" bson:"summary_author"`
	CoPlanners       []*UserOverview   `json:"co_planners" bson:"co_planners"`
	AttendeeCount    int               `json:"attendee_count" bson:"attendee_count"`
	Challenges       string            `json:"challenges" bson:"challenges"`
	Successes        string            `json:"successes" bson:"successes"`
	Improvements     string            `json:"improvements" bson:"improvements"`
	Status           string            `json:"status" bson:"status"`
	CreatedAt        string            `json:"created_at" bson:"created_at"`
	UpdatedAt        string            `json:"updated_at" bson:"updated_at"`
	Notes            []*Note           `json:"notes" bson:"notes"`
}

type EventWithNotes struct {
	ID                      string          `json:"id" bson:"_id"`
	EventLead               *UserOverview   `json:"event_lead" bson:"event_lead"`
	CoPlanners              []*UserOverview `json:"co_planners" bson:"co_planners"`
	Title                   string          `json:"title" bson:"title"`
	Description             string          `json:"description" bson:"description"`
	StartDate               string          `json:"start_date" bson:"start_date"`
	SetUp                   string          `json:"set_up" bson:"set_up"`
	CleanUp                 string          `json:"clean_up" bson:"clean_up"`
	EndDate                 string          `json:"end_date" bson:"end_date"`
	GrantID                 string          `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool            `json:"public_event" bson:"public_event"`
	RsvpRequired            bool            `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool            `json:"annual_event" bson:"annual_event"`
	NewEvent                bool            `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool            `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []string        `json:"agenda" bson:"agenda"`
	TargetAudience          string          `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string       `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial       []*string       `json:"marketing_material" bson:"marketing_material"`
	Supplies                []*string       `json:"supplies" bson:"supplies"`
	SpecialOrders           []*string       `json:"special_orders" bson:"special_orders"`
	Performance             string          `json:"performance" bson:"performance"`
	Vendors                 string          `json:"vendors" bson:"vendors"`
	FoodAndBeverage         []*string       `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                 string          `json:"caterer" bson:"caterer"`
	FoodHeadCount           int             `json:"food_head_count" bson:"food_head_count"`
	EventTeam               []*string       `json:"event_team" bson:"event_team"`
	VolunteerList           []*string       `json:"volunteer_list" bson:"volunteer_list"`
	Budget                  float64         `json:"budget" bson:"budget"`
	AffiliatedOrganizations []*string       `json:"affiliated_organizations" bson:"affiliated_organizations"`
	EducationalGoals        []string        `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes     []string        `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals              []string        `json:"grant_goals" bson:"grant_goals"`
	CreatedAt               string          `json:"created_at" bson:"created_at"`
	UpdatedAt               string          `json:"updated_at" bson:"updated_at"`
	Status                  string          `json:"status" bson:"status"`
	Notes                   []*Note         `json:"notes" bson:"notes"`
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
	Active      bool     `json:"active" bson:"active"`
	CreatedBy   string   `json:"created_by" bson:"created_by"`
	CreatedAt   string   `json:"created_at" bson:"created_at"`
	UpdatedAt   string   `json:"updated_at" bson:"updated_at"`
}

type GrantDetail struct {
	ID          string          `json:"id" bson:"_id"`
	Name        string          `json:"name" bson:"name"`
	Description string          `json:"description" bson:"description"`
	Goals       []string        `json:"goals" bson:"goals"`
	Objectives  []string        `json:"objectives" bson:"objectives"`
	StartDate   string          `json:"start_date" bson:"start_date"`
	AwardDate   string          `json:"award_date" bson:"award_date"`
	EndDate     string          `json:"end_date" bson:"end_date"`
	AwardNumber string          `json:"award_number" bson:"award_number"`
	Budget      float64         `json:"budget" bson:"budget"`
	Active      bool            `json:"active" bson:"active"`
	CreatedBy   []*UserOverview `json:"created_by" bson:"created_by"`
	CreatedAt   string          `json:"created_at" bson:"created_at"`
	UpdatedAt   string          `json:"updated_at" bson:"updated_at"`
}

type GrantOverview struct {
	ID          string  `json:"id" bson:"_id"`
	Name        string  `json:"name" bson:"name"`
	StartDate   string  `json:"start_date" bson:"start_date"`
	AwardDate   string  `json:"award_date" bson:"award_date"`
	EndDate     string  `json:"end_date" bson:"end_date"`
	AwardNumber string  `json:"award_number" bson:"award_number"`
	Budget      float64 `json:"budget" bson:"budget"`
	Active      bool    `json:"active" bson:"active"`
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
	ID        string        `json:"id" bson:"_id"`
	LogAuthor *UserOverview `json:"log_author" bson:"log_author"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	UpdatedAt string        `json:"updated_at" bson:"updated_at"`
	Status    string        `json:"status" bson:"status"`
	NoteCount int           `json:"note_count" bson:"note_count"`
}

type LogRes struct {
	ID        string        `json:"id" bson:"_id"`
	LogAuthor *UserOverview `json:"log_author" bson:"log_author"`
	Status    string        `json:"status" bson:"status"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
}

type LogWithNotes struct {
	ID            string        `json:"id" bson:"_id"`
	LogAuthor     *UserOverview `json:"log_author" bson:"log_author"`
	DailyActivity string        `json:"daily_activity" bson:"daily_activity"`
	Positives     string        `json:"positives" bson:"positives"`
	Improvements  string        `json:"improvements" bson:"improvements"`
	NextSteps     string        `json:"next_steps" bson:"next_steps"`
	Status        string        `json:"status" bson:"status"`
	CreatedAt     string        `json:"created_at" bson:"created_at"`
	UpdatedAt     string        `json:"updated_at" bson:"updated_at"`
	Notes         []*Note       `json:"notes" bson:"notes"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type LoginRes struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Active    bool   `json:"active" bson:"active"`
	Token     string `json:"token" bson:"token"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}

type NewContact struct {
	Name  string  `json:"name" bson:"name"`
	Type  string  `json:"type" bson:"type"`
	Email *string `json:"email" bson:"email"`
	Phone *string `json:"phone" bson:"phone"`
	Notes *string `json:"notes" bson:"notes"`
}

type NewEvent struct {
	CoPlanners              []*string `json:"co_planners" bson:"co_planners"`
	Title                   string    `json:"title" bson:"title"`
	Description             string    `json:"description" bson:"description"`
	StartDate               string    `json:"start_date" bson:"start_date"`
	SetUp                   string    `json:"set_up" bson:"set_up"`
	CleanUp                 string    `json:"clean_up" bson:"clean_up"`
	EndDate                 string    `json:"end_date" bson:"end_date"`
	GrantID                 string    `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool      `json:"public_event" bson:"public_event"`
	RsvpRequired            bool      `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event" bson:"annual_event"`
	NewEvent                bool      `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []string  `json:"agenda" bson:"agenda"`
	TargetAudience          string    `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials" bson:"marketing_materials"`
	Supplies                []*string `json:"supplies" bson:"supplies"`
	SpecialOrders           []*string `json:"special_orders" bson:"special_orders"`
	Performance             *string   `json:"performance" bson:"performance"`
	Vendors                 *string   `json:"vendors" bson:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                 *string   `json:"caterer" bson:"caterer"`
	FoodHeadCount           int       `json:"food_head_count" bson:"food_head_count"`
	EventTeam               []*string `json:"event_team" bson:"event_team"`
	VolunteerList           []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                  float64   `json:"budget" bson:"budget"`
	AffiliatedOrganizations []*string `json:"affiliated_organizations" bson:"affiliated_organizations"`
	EducationalGoals        []string  `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes     []string  `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals              []string  `json:"grant_goals" bson:"grant_goals"`
}

type NewEventSummary struct {
	EventID       string    `json:"event_id" bson:"event_id"`
	CoPlanners    []*string `json:"co_planners" bson:"co_planners"`
	AttendeeCount int       `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges" bson:"challenges"`
	Successes     string    `json:"successes" bson:"successes"`
	Improvements  string    `json:"improvements" bson:"improvements"`
}

type NewGrant struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Goals       []string `json:"goals" bson:"goals"`
	Objectives  []string `json:"objectives" bson:"objectives"`
	StartDate   string   `json:"start_date" bson:"start_date"`
	EndDate     string   `json:"end_date" bson:"end_date"`
	Budget      float64  `json:"budget" bson:"budget"`
	AwardNumber string   `json:"award_number" bson:"award_number"`
	AwardDate   string   `json:"award_date" bson:"award_date"`
}

type NewLog struct {
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives" bson:"positives"`
	Improvements  string `json:"improvements" bson:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
}

type NewNote struct {
	ItemID  string `json:"item_id" bson:"item_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

type NewSchoolReportDebrief struct {
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           int      `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives" bson:"positives"`
	Discussion             string   `json:"discussion" bson:"discussion"`
}

type NewSchoolReportPlan struct {
	Date           string    `json:"date" bson:"date"`
	CoFacilitators []*string `json:"co_facilitators" bson:"co_facilitators"`
	Curriculum     string    `json:"curriculum" bson:"curriculum"`
	School         string    `json:"school" bson:"school"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
}

type NewUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
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

type NoteDetail struct {
	ID        string        `json:"id" bson:"_id"`
	ItemID    string        `json:"item_id" bson:"item_id"`
	Author    *UserOverview `json:"author" bson:"author"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content" bson:"content"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	UpdatedAt string        `json:"updated_at" bson:"updated_at"`
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
	ID            string                       `json:"id" bson:"_id"`
	DebriefAuthor *UserOverview                `json:"debrief_author" bson:"debrief_author"`
	LessonPlan    *SchoolReportPlanDescription `json:"lesson_plan" bson:"lesson_plan"`
	Status        string                       `json:"status" bson:"status"`
	CreatedAt     string                       `json:"created_at" bson:"created_at"`
	NoteCount     int                          `json:"note_count" bson:"note_count"`
}

type SchoolReportDebriefRes struct {
	ID            string                       `json:"id" bson:"_id"`
	DebriefAuthor *UserOverview                `json:"debrief_author" bson:"debrief_author"`
	LessonPlan    *SchoolReportPlanDescription `json:"lesson_plan" bson:"lesson_plan"`
	Status        string                       `json:"status" bson:"status"`
	CreatedAt     string                       `json:"created_at" bson:"created_at"`
}

type SchoolReportDebriefWithNotes struct {
	ID                     string                       `json:"id" bson:"_id"`
	DebriefAuthor          *UserOverview                `json:"debrief_author" bson:"debrief_author"`
	LessonPlan             *SchoolReportPlanDescription `json:"lesson_plan" bson:"lesson_plan"`
	StudentCount           int                          `json:"student_count" bson:"student_count"`
	StudentList            []string                     `json:"student_list" bson:"student_list"`
	ChallengesImprovements string                       `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string                       `json:"positives" bson:"positives"`
	Discussion             string                       `json:"discussion" bson:"discussion"`
	Status                 string                       `json:"status" bson:"status"`
	CreatedAt              string                       `json:"created_at" bson:"created_at"`
	UpdatedAt              string                       `json:"updated_at" bson:"updated_at"`
	Notes                  []*Note                      `json:"notes" bson:"notes"`
}

type SchoolReportPlan struct {
	ID             string    `json:"id" bson:"_id"`
	UserID         string    `json:"user_id" bson:"user_id"`
	Date           string    `json:"date" bson:"date"`
	CoFacilitators []*string `json:"co_facilitators" bson:"co_facilitators"`
	Curriculum     string    `json:"curriculum" bson:"curriculum"`
	School         string    `json:"school" bson:"school"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	Status         string    `json:"status" bson:"status"`
	CreatedAt      string    `json:"created_at" bson:"created_at"`
	UpdatedAt      string    `json:"updated_at" bson:"updated_at"`
}

type SchoolReportPlanDescription struct {
	ID     string `json:"id" bson:"_id"`
	School string `json:"school" bson:"school"`
	Date   string `json:"date" bson:"date"`
}

type SchoolReportPlanOverview struct {
	ID           string        `json:"id" bson:"_id"`
	Date         string        `json:"date" bson:"date"`
	ReportAuthor *UserOverview `json:"report_author" bson:"report_author"`
	School       string        `json:"school" bson:"school"`
	Status       string        `json:"status" bson:"status"`
	CreatedAt    string        `json:"created_at" bson:"created_at"`
	NoteCount    int           `json:"note_count" bson:"note_count"`
}

type SchoolReportPlanRes struct {
	ID         string        `json:"id" bson:"_id"`
	PlanAuthor *UserOverview `json:"plan_author" bson:"plan_author"`
	Date       string        `json:"date" bson:"date"`
	School     string        `json:"school" bson:"school"`
	Status     string        `json:"status" bson:"status"`
	CreatedAt  string        `json:"created_at" bson:"created_at"`
}

type SchoolReportPlanWithNotes struct {
	ID             string          `json:"id" bson:"_id"`
	Date           string          `json:"date" bson:"date"`
	ReportAuthor   *UserOverview   `json:"report_author" bson:"report_author"`
	CoFacilitators []*UserOverview `json:"co_facilitators" bson:"co_facilitators"`
	Curriculum     string          `json:"curriculum" bson:"curriculum"`
	School         string          `json:"school" bson:"school"`
	LessonTopics   string          `json:"lesson_topics" bson:"lesson_topics"`
	Status         string          `json:"status" bson:"status"`
	CreatedAt      string          `json:"created_at" bson:"created_at"`
	UpdatedAt      string          `json:"updated_at" bson:"updated_at"`
	Notes          []*Note         `json:"notes" bson:"notes"`
}

type UpdateContact struct {
	ID    string  `json:"id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Type  string  `json:"type" bson:"type"`
	Email *string `json:"email" bson:"email"`
	Phone *string `json:"phone" bson:"phone"`
	Notes *string `json:"notes" bson:"notes"`
}

type UpdateEvent struct {
	ID                      string    `json:"id" bson:"_id"`
	CoPlanners              []*string `json:"co_planners" bson:"co_planners"`
	Title                   string    `json:"title" bson:"title"`
	Description             string    `json:"description" bson:"description"`
	StartDate               string    `json:"start_date" bson:"start_date"`
	SetUp                   string    `json:"set_up" bson:"set_up"`
	CleanUp                 string    `json:"clean_up" bson:"clean_up"`
	EndDate                 string    `json:"end_date" bson:"end_date"`
	GrantID                 string    `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool      `json:"public_event" bson:"public_event"`
	RsvpRequired            bool      `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event" bson:"annual_event"`
	NewEvent                bool      `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []string  `json:"agenda" bson:"agenda"`
	TargetAudience          string    `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials" bson:"marketing_materials"`
	Supplies                []*string `json:"supplies" bson:"supplies"`
	SpecialOrders           []*string `json:"special_orders" bson:"special_orders"`
	Performance             *string   `json:"performance" bson:"performance"`
	Vendors                 *string   `json:"vendors" bson:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                 *string   `json:"caterer" bson:"caterer"`
	FoodHeadCount           int       `json:"food_head_count" bson:"food_head_count"`
	EventTeam               []*string `json:"event_team" bson:"event_team"`
	VolunteerList           []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                  float64   `json:"budget" bson:"budget"`
	AffiliatedOrganizations []*string `json:"affiliated_organizations" bson:"affiliated_organizations"`
	EducationalGoals        []string  `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes     []string  `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals              []string  `json:"grant_goals" bson:"grant_goals"`
}

type UpdateEventSummary struct {
	ID            string    `json:"id" bson:"_id"`
	EventID       string    `json:"event_id" bson:"event_id"`
	CoPlanners    []*string `json:"co_planners" bson:"co_planners"`
	AttendeeCount int       `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges" bson:"challenges"`
	Successes     string    `json:"successes" bson:"successes"`
	Improvements  string    `json:"improvements" bson:"improvements"`
	Status        string    `json:"status" bson:"status"`
}

type UpdateGrant struct {
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Goals       []string `json:"goals" bson:"goals"`
	Objectives  []string `json:"objectives" bson:"objectives"`
	StartDate   string   `json:"start_date" bson:"start_date"`
	EndDate     string   `json:"end_date" bson:"end_date"`
	Budget      float64  `json:"budget" bson:"budget"`
	AwardNumber string   `json:"award_number" bson:"award_number"`
	AwardDate   string   `json:"award_date" bson:"award_date"`
	Active      bool     `json:"active" bson:"active"`
}

type UpdateLog struct {
	ID            string `json:"id" bson:"_id"`
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives" bson:"positives"`
	Improvements  string `json:"improvements" bson:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
	Status        string `json:"status" bson:"status"`
}

type UpdateNote struct {
	ID      string `json:"id" bson:"_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

type UpdateSchoolReportDebrief struct {
	ID                     string   `json:"id" bson:"_id"`
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           int      `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives" bson:"positives"`
	Discussion             string   `json:"discussion" bson:"discussion"`
	Status                 string   `json:"status" bson:"status"`
}

type UpdateSchoolReportPlan struct {
	ID             string    `json:"id" bson:"_id"`
	Date           string    `json:"date" bson:"date"`
	CoFacilitators []*string `json:"co_facilitators" bson:"co_facilitators"`
	Curriculum     string    `json:"curriculum" bson:"curriculum"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	School         string    `json:"school" bson:"school"`
	Status         string    `json:"status" bson:"status"`
}

type UpdateUser struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Admin     bool   `json:"admin" bson:"admin"`
	Active    bool   `json:"active" bson:"active"`
}

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Admin     bool   `json:"admin" bson:"admin"`
	Active    bool   `json:"active" bson:"active"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}

func (User) IsUserResult() {}

type UserOverview struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

func (UserOverview) IsUserResult() {}

type UserUpdateRes struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Admin     bool   `json:"admin" bson:"admin"`
	Active    bool   `json:"active" bson:"active"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
