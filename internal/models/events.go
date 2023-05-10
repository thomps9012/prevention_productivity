package models

type NewEvent struct {
	CoPlanners              []*string `json:"co_planners"`
	Title                   string    `json:"title"`
	Description             string    `json:"description"`
	StartDate               string    `json:"start_date"`
	SetUp                   string    `json:"set_up"`
	CleanUp                 string    `json:"clean_up"`
	EndDate                 string    `json:"end_date"`
	GrantID                 string    `json:"grant_id"`
	PublicEvent             bool      `json:"public_event"`
	RsvpRequired            bool      `json:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event"`
	NewEvent                bool      `json:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed"`
	Agenda                  []string  `json:"agenda"`
	TargetAudience          string    `json:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials"`
	Supplies                []*string `json:"supplies"`
	SpecialOrders           []*string `json:"special_orders"`
	Performance             *string   `json:"performance"`
	Vendors                 *string   `json:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage"`
	Caterer                 *string   `json:"caterer"`
	FoodHeadCount           int       `json:"food_head_count"`
	EventTeam               []*string `json:"event_team"`
	VolunteerList           []*string `json:"volunteer_list"`
	Budget                  float64   `json:"budget"`
	AffiliatedOrganizations []*string `json:"affiliated_organizations"`
	EducationalGoals        []string  `json:"educational_goals"`
	EducationalOutcomes     []string  `json:"educational_outcomes"`
	GrantGoals              []string  `json:"grant_goals"`
}

type UpdateEvent struct {
	ID                      string    `json:"id"`
	CoPlanners              []*string `json:"co_planners"`
	Title                   string    `json:"title"`
	Description             string    `json:"description"`
	StartDate               string    `json:"start_date"`
	SetUp                   string    `json:"set_up"`
	CleanUp                 string    `json:"clean_up"`
	EndDate                 string    `json:"end_date"`
	GrantID                 string    `json:"grant_id"`
	PublicEvent             bool      `json:"public_event"`
	RsvpRequired            bool      `json:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event"`
	NewEvent                bool      `json:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed"`
	Agenda                  []string  `json:"agenda"`
	TargetAudience          string    `json:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials"`
	Supplies                []*string `json:"supplies"`
	SpecialOrders           []*string `json:"special_orders"`
	Performance             *string   `json:"performance"`
	Vendors                 *string   `json:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage"`
	Caterer                 *string   `json:"caterer"`
	FoodHeadCount           int       `json:"food_head_count"`
	EventTeam               []*string `json:"event_team"`
	VolunteerList           []*string `json:"volunteer_list"`
	Budget                  float64   `json:"budget"`
	AffiliatedOrganizations []*string `json:"affiliated_organizations"`
	EducationalGoals        []string  `json:"educational_goals"`
	EducationalOutcomes     []string  `json:"educational_outcomes"`
	GrantGoals              []string  `json:"grant_goals"`
}

type Event struct {
	ID                      string    `json:"id" bson:"_id"`
	EventLead               string    `json:"event_lead" bson:"event_lead"`
	Co_Planners             []*string `json:"co_planners" bson:"co_planners"`
	Title                   string    `json:"title" bson:"title"`
	Description             string    `json:"description" bson:"description"`
	StartDate               string    `json:"start_date" bson:"start_date"`
	SetUp                   string    `json:"set_up" bson:"set_up"`
	CleanUp                 string    `json:"clean_up" bson:"clean_up"`
	EndDate                 string    `json:"end_date" bson:"end_date"`
	GrantID                 string    `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool      `json:"public_event" bson:"public_event"`
	RSVPRequired            bool      `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event" bson:"annual_event"`
	NewEvent                bool      `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []string  `json:"agenda" bson:"agenda"`
	TargetAudience          string    `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials" bson:"marketing_materials"`
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
	ID        string `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
}

type EventOverview struct {
	ID        string          `json:"id"`
	EventLead []*UserOverview `json:"event_lead"`
	Title     string          `json:"title"`
	StartDate string          `json:"start_date"`
	CreatedAt string          `json:"created_at"`
	Status    string          `json:"status"`
	NoteCount int             `json:"note_count"`
}

type EventRes struct {
	ID        string        `json:"id"`
	EventLead *UserOverview `json:"event_lead"`
	Title     string        `json:"title"`
	StartDate string        `json:"start_date"`
	Status    string        `json:"status"`
	CreatedAt string        `json:"created_at"`
}

type EventWithNotes struct {
	ID                      string          `json:"id"`
	EventLead               []*UserOverview `json:"event_lead"`
	CoPlanners              []*UserOverview `json:"co_planners"`
	Title                   string          `json:"title"`
	Description             string          `json:"description"`
	StartDate               string          `json:"start_date"`
	SetUp                   string          `json:"set_up"`
	CleanUp                 string          `json:"clean_up"`
	EndDate                 string          `json:"end_date"`
	GrantID                 string          `json:"grant_id"`
	PublicEvent             bool            `json:"public_event"`
	RsvpRequired            bool            `json:"rsvp_required"`
	AnnualEvent             bool            `json:"annual_event"`
	NewEvent                bool            `json:"new_event"`
	VolunteersNeeded        bool            `json:"volunteers_needed"`
	Agenda                  []string        `json:"agenda"`
	TargetAudience          string          `json:"target_audience"`
	PartingGifts            []*string       `json:"parting_gifts"`
	MarketingMaterial       []*string       `json:"marketing_material"`
	Supplies                []*string       `json:"supplies"`
	SpecialOrders           []*string       `json:"special_orders"`
	Performance             string          `json:"performance"`
	Vendors                 string          `json:"vendors"`
	FoodAndBeverage         []*string       `json:"food_and_beverage"`
	Caterer                 string          `json:"caterer"`
	FoodHeadCount           int             `json:"food_head_count"`
	EventTeam               []*string       `json:"event_team"`
	VolunteerList           []*string       `json:"volunteer_list"`
	Budget                  float64         `json:"budget"`
	AffiliatedOrganizations []*string       `json:"affiliated_organizations"`
	EducationalGoals        []string        `json:"educational_goals"`
	EducationalOutcomes     []string        `json:"educational_outcomes"`
	GrantGoals              []string        `json:"grant_goals"`
	CreatedAt               string          `json:"created_at"`
	UpdatedAt               string          `json:"updated_at"`
	Status                  string          `json:"status"`
	Notes                   []*Note         `json:"notes"`
}
