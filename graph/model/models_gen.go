// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AllEventSummaries struct {
	Event        *Event        `json:"event"`
	EventSummary *EventSummary `json:"event_summary" bson:"event_summary"`
	User         *User         `json:"user"`
	Coplanners   []*User       `json:"coplanners"`
	NoteCount    *int          `json:"noteCount"`
}

type AllEvents struct {
	Event      *Event  `json:"event"`
	User       *User   `json:"user"`
	Coplanners []*User `json:"coplanners"`
	NoteCount  *int    `json:"noteCount"`
}

type AllLogs struct {
	Log       *Log  `json:"log"`
	User      *User `json:"user"`
	NoteCount *int  `json:"noteCount"`
}

type AllSchoolReportDebriefs struct {
	SchoolReportDebrief *SchoolReportDebrief `json:"school_report_debrief" bson:"school_report_debrief"`
	User                *User                `json:"user"`
	NoteCount           *int                 `json:"noteCount"`
}

type AllSchoolReportPlans struct {
	SchoolReportPlan *SchoolReportPlan `json:"school_report_plan" bson:"school_report_plan"`
	User             *User             `json:"user"`
	NoteCount        *int              `json:"noteCount"`
}

type Contact struct {
	ID        *string `json:"id" bson:"_id"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Notes     *string `json:"notes"`
	Type      *string `json:"type"`
	Active    bool    `json:"active"`
	CreatedBy string  `json:"created_by" bson:"created_by"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt string  `json:"deleted_at" bson:"deleted_at"`
}

type ContactInfo struct {
	Contact        *Contact `json:"contact"`
	ContactCreator *User    `json:"contact_creator" bson:"contact_creator"`
}

type Event struct {
	ID                     *string   `json:"id" bson:"_id"`
	EventLead              *string   `json:"event_lead" bson:"event_lead"`
	Coplanners             []*string `json:"coplanners"`
	Title                  string    `json:"title"`
	Description            string    `json:"description"`
	StartDate              string    `json:"start_date" bson:"start_date"`
	SetUp                  string    `json:"set_up" bson:"set_up"`
	CleanUp                string    `json:"clean_up" bson:"clean_up"`
	EndDate                string    `json:"end_date" bson:"end_date"`
	GrantID                string    `json:"grant_id" bson:"grant_id"`
	Public                 bool      `json:"public"`
	Rsvp                   bool      `json:"rsvp"`
	AnnualEvent            bool      `json:"annual_event" bson:"annual_event"`
	NewEvent               bool      `json:"new_event" bson:"new_event"`
	Volunteers             bool      `json:"volunteers"`
	Agenda                 []string  `json:"agenda"`
	TargetAudience         string    `json:"target_audience" bson:"target_audience"`
	PartingGifts           []string  `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial      []string  `json:"marketing_material" bson:"marketing_material"`
	Supplies               []string  `json:"supplies"`
	SpecialOrders          []*string `json:"special_orders" bson:"special_orders"`
	Performance            string    `json:"performance"`
	Vendors                string    `json:"vendors"`
	FoodAndBeverage        []string  `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                string    `json:"caterer"`
	FoodHeadCount          int       `json:"food_head_count" bson:"food_head_count"`
	EventTeam              []*string `json:"event_team" bson:"event_team"`
	VolunteerList          []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                 float64   `json:"budget"`
	AffiliatedOrganization *string   `json:"affiliated_organization" bson:"affiliated_organization"`
	EducationalGoals       []string  `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes    []string  `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals             []string  `json:"grant_goals" bson:"grant_goals"`
	CreatedAt              string    `json:"created_at" bson:"created_at"`
	UpdatedAt              string    `json:"updated_at" bson:"updated_at"`
	Status                 string    `json:"status"`
}

type EventSummary struct {
	ID            string    `json:"id" bson:"_id"`
	EventID       string    `json:"event_id" bson:"event_id"`
	UserID        string    `json:"user_id" bson:"user_id"`
	Coplanners    []*string `json:"coplanners"`
	AttendeeCount *int      `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges"`
	Successes     string    `json:"successes"`
	Improvements  string    `json:"improvements"`
	Status        string    `json:"status"`
	CreatedAt     string    `json:"created_at" bson:"created_at"`
	UpdatedAt     string    `json:"updated_at" bson:"updated_at"`
}

type EventSummaryWithNotes struct {
	EventSummary *EventSummary `json:"event_summary" bson:"event_summary"`
	Notes        []*Note       `json:"notes"`
}

type EventWithNotes struct {
	Event *Event  `json:"event"`
	Notes []*Note `json:"notes"`
}

type Grant struct {
	ID          *string   `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Goals       []*string `json:"goals"`
	Objectives  []*string `json:"objectives"`
	StartDate   string    `json:"start_date" bson:"start_date"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
	EndDate     string    `json:"end_date" bson:"end_date"`
	AwardNumber string    `json:"award_number" bson:"award_number"`
	Budget      *float64  `json:"budget"`
	Active      bool      `json:"active"`
	CreatedBy   string    `json:"created_by" bson:"created_by"`
	CreatedAt   string    `json:"created_at" bson:"created_at"`
	UpdatedAt   string    `json:"updated_at" bson:"updated_at" `
}

type Log struct {
	ID            *string `json:"id" bson:"_id"`
	UserID        *string `json:"user_id" bson:"user_id"`
	DailyActivity string  `json:"daily_activity" bson:"daily_activity"`
	Positives     string  `json:"positives"`
	Improvements  string  `json:"improvements"`
	NextSteps     string  `json:"next_steps" bson:"next_steps"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at" bson:"created_at"`
	UpdatedAt     string  `json:"updated_at"  bson:"updated_at"`
}

type LogWithNotes struct {
	Log   *Log    `json:"log"`
	Notes []*Note `json:"notes"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewContact struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Notes *string `json:"notes"`
}

type NewEvent struct {
	Coplanners             []*string `json:"coplanners"`
	Title                  *string   `json:"title"`
	Description            *string   `json:"description"`
	StartDate              *string   `json:"start_date" bson:"start_date"`
	SetUp                  *string   `json:"set_up" bson:"set_up"`
	CleanUp                *string   `json:"clean_up" bson:"clean_up"`
	EndDate                *string   `json:"end_date" bson:"end_date"`
	GrantID                *string   `json:"grant_id" bson:"grant_id"`
	Public                 *bool     `json:"public"`
	Rsvp                   *bool     `json:"rsvp"`
	AnnualEvent            *bool     `json:"annual_event" bson:"annual_event"`
	NewEvent               *bool     `json:"new_event" bson:"new_event"`
	Volunteers             *bool     `json:"volunteers"`
	Agenda                 []*string `json:"agenda"`
	TargetAudience         *string   `json:"target_audience" bson:"target_audience"`
	PartingGifts           []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial      []*string `json:"marketing_material" bson:"marketing_material"`
	Supplies               []*string `json:"supplies"`
	SpecialOrders          []*string `json:"special_orders" bson:"special_orders"`
	Performance            *string   `json:"performance"`
	Vendors                *string   `json:"vendors"`
	FoodAndBeverage        []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                *string   `json:"caterer"`
	FoodHeadCount          *int      `json:"food_head_count" bson:"food_head_count"`
	EventTeam              []*string `json:"event_team" bson:"event_team"`
	VolunteerList          []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                 *float64  `json:"budget"`
	AffiliatedOrganization *string   `json:"affiliated_organization" bson:"affiliated_organization"`
	EducationalGoals       []*string `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes    []*string `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals             []*string `json:"grant_goals" bson:"grant_goals"`
}

type NewEventSummary struct {
	EventID       *string   `json:"event_id" bson:"event_id"`
	Coplanners    []*string `json:"coplanners"`
	AttendeeCount *int      `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges"`
	Successes     string    `json:"successes"`
	Improvements  string    `json:"improvements"`
}

type NewGrant struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Goals       []*string `json:"goals"`
	Objectives  []*string `json:"objectives"`
	StartDate   *string   `json:"start_date" bson:"start_date"`
	EndDate     *string   `json:"end_date" bson:"end_date"`
	Budget      *float64  `json:"budget"`
	AwardNumber *string   `json:"award_number" bson:"award_number"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
}

type NewLog struct {
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives"`
	Improvements  string `json:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
}

type NewNote struct {
	ItemID  string `json:"item_id" bson:"item_id`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewSchoolReportDebrief struct {
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           *int     `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
}

type NewSchoolReportPlan struct {
	Cofacilitators []*string `json:"cofacilitators"`
	Curriculum     string    `json:"curriculum"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	School         string    `json:"school"`
}

type NewUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Note struct {
	ID        *string `json:"id" bson:"_id"`
	ItemID    *string `json:"item_id" bson:"item_id"`
	UserID    *string `json:"user_id" bson:"user_id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type SchoolReportDebrief struct {
	ID                     string   `json:"id" bson:"_id"`
	UserID                 string   `json:"user_id" bson:"user_id"`
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           *int     `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
	Status                 string   `json:"status"`
	CreatedAt              string   `json:"created_at" bson:"created_at"`
	UpdatedAt              string   `json:"updated_at" bson:"updated_at"`
}

type SchoolReportDebriefWithNotes struct {
	SchoolReportDebrief *SchoolReportDebrief `json:"school_report_debrief" bson:"school_report_debrief"`
	Notes               []*Note              `json:"notes"`
}

type SchoolReportPlan struct {
	ID             *string   `json:"id" bson:"_id"`
	UserID         *string   `json:"user_id" bson:"user_id"`
	Cofacilitators []*string `json:"cofacilitators"`
	Curriculum     string    `json:"curriculum"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	School         string    `json:"school"`
	Status         string    `json:"status"`
	CreatedAt      string    `json:"created_at" bson:"created_at"`
	UpdatedAt      string    `json:"updated_at" bson:"updated_at"`
}

type SchoolReportPlanWithNotes struct {
	SchoolReportPlan *SchoolReportPlan `json:"school_report_plan"`
	Notes            []*Note           `json:"notes"`
}

type UpdateContact struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Notes *string `json:"notes"`
}

type UpdateEvent struct {
	Coplanners             []*string `json:"coplanners"`
	Title                  *string   `json:"title"`
	Description            *string   `json:"description"`
	StartDate              *string   `json:"start_date" bson:"start_date"`
	SetUp                  *string   `json:"set_up" bson:"set_up"`
	CleanUp                *string   `json:"clean_up" bson:"clean_up"`
	EndDate                *string   `json:"end_date" bson:"end_date"`
	GrantID                *string   `json:"grant_id" bson:"grant_id"`
	Public                 *bool     `json:"public"`
	Rsvp                   *bool     `json:"rsvp"`
	AnnualEvent            *bool     `json:"annual_event" bson:"annual_event"`
	NewEvent               *bool     `json:"new_event" bson:"new_event"`
	Volunteers             *bool     `json:"volunteers"`
	Agenda                 []*string `json:"agenda"`
	TargetAudience         *string   `json:"target_audience" bson:"target_audience"`
	PartingGifts           []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial      []*string `json:"marketing_material" bson:"marketing_materials"`
	Supplies               []*string `json:"supplies"`
	SpecialOrders          []*string `json:"special_orders" bson:"special_orders"`
	Performance            *string   `json:"performance"`
	Vendors                *string   `json:"vendors"`
	FoodAndBeverage        []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                *string   `json:"caterer"`
	FoodHeadCount          *int      `json:"food_head_count" bson:"food_head_count"`
	EventTeam              []*string `json:"event_team" bson:"event_team"`
	VolunteerList          []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                 *float64  `json:"budget"`
	AffiliatedOrganization *string   `json:"affiliated_organization" bson:"affiliated_organization"`
	EducationalGoals       []*string `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes    []*string `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals             []*string `json:"grant_goals" bson:"grant_goals"`
	Status                 *string   `json:"status"`
}

type UpdateEventSummary struct {
	AttendeeCount *int      `json:"attendee_count" bson:"attendee_count"`
	Coplanners    []*string `json:"coplanners"`
	Challenges    string    `json:"challenges"`
	Successes     string    `json:"successes"`
	Improvements  string    `json:"improvements"`
	Status        *string   `json:"status"`
}

type UpdateGrant struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Goals       []*string `json:"goals"`
	Objectives  []*string `json:"objectives"`
	StartDate   *string   `json:"start_date" bson:"start_date"`
	EndDate     *string   `json:"end_date" bson:"end_date"`
	Budget      *float64  `json:"budget"`
	AwardNumber *string   `json:"award_number" bson:"award_number"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
	Active      *bool     `json:"active"`
}

type UpdateLog struct {
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives"`
	Improvements  string `json:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
	Status        string `json:"status"`
}

type UpdateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateSchoolReportDebrief struct {
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           *int     `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
	Status                 *string  `json:"status"`
}

type UpdateSchoolReportPlan struct {
	Cofacilitators []*string `json:"cofacilitators"`
	Curriculum     string    `json:"curriculum"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	School         string    `json:"school"`
	Status         *string   `json:"status"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
}

type User struct {
	ID        *string `json:"id" bson:"_id"`
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Admin     bool    `json:"admin"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt *string `json:"deleted_at" bson:"deleted_at"`
	Active    bool    `json:"active"`
}
