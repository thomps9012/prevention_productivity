package model

type Event struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	GrantID   *Grant `bson:"grant"`
	Name 	string        `bson:"name"`
	Description string        `bson:"description"`
	Location string        `bson:"location"`
	LocationDetails string        `bson:"location_details"`
	NewEvent bool `bson:"new_event"`
	PublicEvent bool        `bson:"public_event"`
	RSVPsNeeded int `bson:"rsvps_needed"`
	AnnualEvent bool `bson:"annual_event"`
	StartDate  time.Time     `bson:"start_date"`
	SetUpTime time.Time `bson:"set_up_time"`
	CleanUpTime time.Time `bson:"clean_up_time"`
	Agenda []AgendaItem `bson:"agenda"`
	TargetPopulation []string `bson:"target_population"`
	AgeGroup []string `bson:"age_group"`
	PartingGifts []string `bson:"parting_gifts"`
	Raffle []string `bson:"raffle"`
	Marketing []string `bson:"marketing"`
	SpecialOrder []SpecialOrderItem `bson:"special_order"`
	Vendor []VendorItem `bson:"vendor"`
	FoodBeverage []string `bson:"food_beverage"`
	Caterer []string `bson:"caterer"`
	FoodHeadCount int `bson:"food_head_count"`
	PreventionTeam []*User `bson:"prevention_team"`
	StaffTimeCommitment double `bson:"staff_time_commitment"`
	StaffDuties []string `bson:"staff_duties"`
	Clients []string `bson:"nora_clients"`
	VolunteersNeeded int `bson:"volunteers_needed"`
	Volunteers []string `bson:"nora_volunteers"`
	Budget float64 `bson:"budget"`
	Affiliates []string `bson:"affiliates"`
	Notes []Note `bson:"notes"`
	EventLead *User `bson:"event_lead"`
	EducationalGoal []string `bson:"educational_goal"`
	Curriculum []string `bson:"curriculum"`
	Outreach []string `bson:"outreach"`
	GrantGoals []string `bson:"grant_goals"`
	GuestList []GuestListItem `bson:"guest_list"`
	EventStatus enum `bson:"event_status"`
	Notes []Note `bson:"notes"`
}

type EventStatus string
const (
	Pending,
	Approved,
	Denied,
	Cancelled,
	Complete
)

// func (u *User) NewEvent(args [string]interface{}) Event {
// 	event := Event{}
// 	event.EventLead = u.ID
// 	event.EventStatus = Pending
// 	db.C("events").Insert(event)
// 	return event
// }

// func (u *User) IsEventLead(eventId bson.ObjectId) bool {
// 	filter := bson.M{"_id": logID, "user_id": u.ID}
// 	var event Event
// 	err := db.C("events").Find(filter).One(&log)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func (u *User) EditEvent(args map[string]interface{}) {
// 	if(u.IsAdmin()) {
// 		db.C("events").UpdateId(e.ID, bson.M{"$set": args})
// 	} else if(u.IsEventLead(args["id"].(bson.ObjectId))) {
// 		db.C("events").UpdateId(e.ID, bson.M{"$set": args})
// 	} else {
// 		panic("You are not authorized to edit events.")
// 	}
// }

// func (e *Event) GetEventLead() User {
// 	var user User
// 	db.C("users").FindId(e.EventLead).One(&user)
// 	return user
// }


// func (u *User) ApproveEvent(event_id bson.ObjectId) {
// 	if(u.IsAdmin()) {
// 		db.C("events").UpdateId(event_id, bson.M{"$set": bson.M{"event_status": Approved}})
// 	}
// }

// func (u *User) DenyEvent(event_id bson.ObjectId) {
// 	if(u.IsAdmin()) {
// 		db.C("events").UpdateId(event_id, bson.M{"$set": bson.M{"event_status": Denied}})
// 	}
// }

// func (u *User) CancelEvent(event_id bson.ObjectId) {
// 	if(u.IsAdmin()) {
// 		db.C("events").UpdateId(event_id, bson.M{"$set": bson.M{"event_status": Cancelled}})
// 	}
// }

// func (u *User) CompleteEvent(event_id bson.ObjectId) {
// 	if(u.IsAdmin()) {
// 		db.C("events").UpdateId(event_id, bson.M{"$set": bson.M{"event_status": Complete}})
// 	}
// }

// func (u *User) GetEvents() []Event {
// 	if(u.IsAdmin()) {
// 		var events []Event
// 		db.C("events").Find(nil).All(&events)
// 		return events
// 	} else {
// 		var events []Event
// 		db.C("events").Find(bson.M{"event_lead": u.ID}).All(&events)
// 		return events
// 	}
// }