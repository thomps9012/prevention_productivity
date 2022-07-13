package model

type AllEventSummaries struct {
	Event        *Event        `json:"event"`
	EventSummary *EventSummary `json:"event_summary" bson:"event_summary"`
	User         *User         `json:"user"`
	NoteCount    *int          `json:"noteCount"`
}

type AllEvents struct {
	Event     *Event `json:"event"`
	User      *User  `json:"user"`
	NoteCount *int   `json:"noteCount"`
}

type AllLogs struct {
	Log       *Log  `json:"log"`
	User      *User `json:"user"`
	NoteCount *int  `json:"noteCount"`
}

type AllSchoolReports struct {
	SchoolReport *SchoolReport `json:"school_report" bson:"school_report"`
	User         *User         `json:"user"`
	NoteCount    *int          `json:"noteCount"`
}

type Contact struct {
	ID        *string `json:"id" bson:"_id"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Notes     *string `json:"notes"`
	Type      *string `json:"type"`
	IsActive  bool    `json:"is_active" bson:"is_active"`
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
	Volunteers             bool      `json:"volunteers" bson:"volunteers"`
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
	ID            string `json:"id" bson:"_id"`
	EventID       string `json:"event_id" bson:"event_id"`
	UserID        string `json:"user_id" bson:"user_id"`
	AttendeeCount int    `json:"attendee_count" bson:"attendee_count"`
	Challenges    string `json:"challenges"`
	Successes     string `json:"successes"`
	Improvements  string `json:"improvements"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
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
	StartDate   string    `json:"start_date" bson:"start_date"`
	Goals       []*string `json:"goals" bson:"goals"`
	Objectives  []*string `json:"objectives" bson:"objectives"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
	EndDate     string    `json:"end_date" bson:"end_date"`
	AwardNumber string    `json:"award_number" bson:"award_number"`
	Budget      *float64  `json:"budget"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by" bson:"created_by"`
	CreatedAt   string    `json:"created_at" bson:"created_at"`
	UpdatedAt   string    `json:"updated_at" bson:"updated_at"`
}

type Log struct {
	ID           *string `json:"id" bson:"_id"`
	UserID       *string `json:"user_id" bson:"user_id"`
	FocusArea    string  `json:"focus_area" bson:"focus_area"`
	Actions      string  `json:"actions"`
	Successes    string  `json:"successes"`
	Improvements string  `json:"improvements"`
	NextSteps    string  `json:"next_steps" bson:"next_steps"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at" bson:"created_at"`
	UpdatedAt    string  `json:"updated_at" bson:"updated_at"`
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
	Volunteers             *bool     `json:"volunteers" bson:"volunteers"`
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
	EventID       *string `json:"event_id" bson:"event_id"`
	AttendeeCount *int    `json:"attendee_count" bson:"attendee_count"`
	Challenges    *string `json:"challenges"`
	Successes     *string `json:"successes"`
	Improvements  *string `json:"improvements"`
}

type NewGrant struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Goals       []*string `json:"goals" bson:"goals"`
	Objectives  []*string `json:"objectives" bson:"objectives"`
	StartDate   *string   `json:"start_date" bson:"start_date"`
	EndDate     *string   `json:"end_date" bson:"end_date"`
	Budget      *float64  `json:"budget"`
	AwardNumber *string   `json:"award_number" bson:"award_number"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
}

type NewLog struct {
	FocusArea    string `json:"focus_area" bson:"focus_area"`
	Actions      string `json:"actions"`
	Successes    string `json:"successes"`
	Improvements string `json:"improvements"`
	NextSteps    string `json:"next_steps" bson:"next_steps"`
}

type NewNote struct {
	ItemID  string `json:"item_id" bson:"item_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewSchoolReport struct {
	Curriculum   *string  `json:"curriculum"`
	LessonPlan   *string  `json:"lesson_plan" bson:"lesson_plan"`
	School       *string  `json:"school"`
	Topics       *string  `json:"topics"`
	StudentCount *int     `json:"student_count" bson:"student_count"`
	StudentList  []string `json:"student_list" bson:"student_list"`
	Challenges   *string  `json:"challenges"`
	Successes    *string  `json:"successes"`
	Improvements *string  `json:"improvements"`
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
	Title     string  `json:"title" bson:"title"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type SchoolReport struct {
	ID           *string  `json:"id" bson:"_id"`
	UserID       *string  `json:"user_id" bson:"user_id"`
	Curriculum   string   `json:"curriculum"`
	LessonPlan   string   `json:"lesson_plan" bson:"lesson_plan"`
	School       string   `json:"school"`
	Topics       string   `json:"topics"`
	StudentCount int      `json:"student_count" bson:"student_count"`
	StudentList  []string `json:"student_list" bson:"student_list"`
	Challenges   string   `json:"challenges"`
	Successes    string   `json:"successes"`
	Improvements string   `json:"improvements"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at" bson:"created_at"`
	UpdatedAt    string   `json:"updated_at" bson:"updated_at"`
}

type SchoolReportWithNotes struct {
	SchoolReport *SchoolReport `json:"school_report" bson:"school_report"`
	Notes        []*Note       `json:"notes"`
}

type UpdateContact struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Type  *string `json:"type"`
	Phone *string `json:"phone"`
	Notes *string `json:"notes"`
}

type UpdateEvent struct {
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
	Status                 *string   `json:"status"`
}

type UpdateEventSummary struct {
	AttendeeCount *int    `json:"attendee_count" bson:"attendee_count"`
	Challenges    *string `json:"challenges"`
	Successes     *string `json:"successes"`
	Improvements  *string `json:"improvements"`
	Status        *string `json:"status"`
}

type UpdateGrant struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Goals       []*string `json:"goals" bson:"goals"`
	Objectives  []*string `json:"objectives" bson:"objectives"`
	StartDate   *string   `json:"start_date" bson:"start_date"`
	EndDate     *string   `json:"end_date" bson:"end_date"`
	Budget      *float64  `json:"budget"`
	AwardNumber *string   `json:"award_number" bson:"award_number"`
	AwardDate   *string   `json:"award_date" bson:"award_date"`
	IsActive    *bool     `json:"is_active" bson:"is_active"`
}

type UpdateLog struct {
	FocusArea    string `json:"focus_area" bson:"focus_area"`
	Actions      string `json:"actions"`
	Successes    string `json:"successes"`
	Improvements string `json:"improvements"`
	NextSteps    string `json:"next_steps" bson:"next_steps"`
	Status       string `json:"status"`
}

type UpdateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateSchoolReport struct {
	Curriculum   *string  `json:"curriculum"`
	LessonPlan   *string  `json:"lesson_plan" bson:"lesson_plan"`
	School       *string  `json:"school"`
	Topics       *string  `json:"topics"`
	StudentCount *int     `json:"student_count" bson:"student_count"`
	StudentList  []string `json:"student_list" bson:"student_list"`
	Challenges   *string  `json:"challenges"`
	Successes    *string  `json:"successes"`
	Improvements *string  `json:"improvements"`
	Status       *string  `json:"status"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Admin     bool
	Active    bool
}

type User struct {
	ID        *string `json:"id" bson:"_id"`
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Admin     bool
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt *string `json:"deleted_at" bson:"deleted_at"`
	Active    bool
}
