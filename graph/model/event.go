package model

type Event struct {
	ID        string `json:"id"`
	Grant   *Grant `json:"grant"`
	Name 	string        `json:"name"`
	Description string        `json:"description"`
	Location string        `json:"location"`
	LocationDetails string        `json:"location_details"`
	NewEvent bool `json:"new_event"`
	PublicEvent bool        `json:"public_event"`
	RSVPsNeeded int `json:"rsvps_needed"`
	AnnualEvent bool `json:"annual_event"`
	StartTime  string     `json:"start_time"`
	SetUpTime string `json:"set_up_time"`
	CleanUpTime string `json:"clean_up_time"`
	Agenda []string `json:"agenda"`
	TargetPopulation string `json:"target_population"`
	AgeGroup string `json:"age_group"`
	PartingGifts []string `json:"parting_gifts"`
	Raffle []string `json:"raffle_items"`
	Marketing string `json:"marketing_material"`
	SpecialOrder string `json:"special_order"`
	Vendors []string `json:"vendors"`
	FoodBeverage string `json:"food_beverage"`
	Caterer string `json:"caterer"`
	FoodHeadCount int `json:"food_head_count"`
	PreventionTeam []*User `json:"prevention_team"`
	StaffTimeCommitment float64 `json:"staff_time_commitment"`
	StaffDuties []string `json:"staff_duties"`
	Clients []string `json:"nora_clients"`
	VolunteersNeeded int `json:"volunteers_needed"`
	Volunteers []string `json:"volunteer_list"`
	Budget float64 `json:"budget"`
	Affiliates []string `json:"affiliates"`
	EventLead *User `json:"event_lead"`
	EducationalGoal string `json:"educational_goal"`
	CurriculumPlan string `json:"curriculum_plan"`
	Outreach string `json:"outreach"`
	GrantGoals []string `json:"grant_goals"`
	GuestList []string `json:"guest_list"`
	EventStatus ApprovalStatus `json:"event_status"`
	Notes []Note `json:"notes"`
}

// type EventStatus string
// const (
// 	Pending string = "Pending"
// 	Approved string = "Approved"
// 	Denied string = "Denied"
// 	Cancelled string = "Cancelled"
// 	Complete string = "Complete"
// )

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