package model

type EventSummary struct {
	ID        string `json:"id"`
	Author *User `json:"author"`
	Event  *Event `json:"event"`
	Attendees int 		 `json:"attendees"`
	Challenges []string `json:"challenges"`
	Outcomes []string `json:"outcomes"`
	NextSteps []string `json:"next_steps"`
	ApprovalStatus string `json:"approval_status"`
	CreatedAt string     `json:"created_at"`
	Notes []Note `json:"notes"`
}

// func (u *User) NewSummary(args [string]interface{}) EventSummary {
// 	eventSummary := EventSummary{}
// 	eventSummary.UserID = u.ID
// 	eventSummary.ApprovalStatus = Pending
// 	eventSummary.CreatedAt = time.Now()
// 	db.C("event_summaries").Insert(eventSummary)
// 	return eventSummary
// }

// func (u *User) IsSummaryAuthor(summaryID bson.ObjectId) bool {
// 	filter := bson.M{"_id": summaryID, "user_id": u.ID}
// 	var summary EventSummary
// 	err := db.C("event_summaries").Find(filter).One(&summary)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func (u *User) EditSummary(args map[string]interface{}) {
// 	if(u.IsAdmin()) {
// 		db.C("event_summaries").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else if(u.IsSummaryAuthor(args["id"].(bson.ObjectId))) {
// 		db.C("event_summaries").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else {
// 		panic("You are not authorized to edit this summary")
// 	}
// }

// func (u *User) ApproveSummary(summary EventSummary) {
// 	if(u.IsAdmin()) {
// 	db.C("event_summaries").Update(bson.M{"_id": summary.ID}, bson.M{"$set": bson.M{"approval_status": Approved}})
// 	}
// }

// func (u *User) DenySummary(summary EventSummary) {
// 	if(u.IsAdmin()) {
// 	db.C("event_summaries").Update(bson.M{"_id": summary.ID}, bson.M{"$set": bson.M{"approval_status": Denied}})
// 	}
// }

// func (u *User) GetSummaries() []EventSummary {
// 	var reports []EventSummary
// 	if(u.IsAdmin()) {
// 		db.C("event_summaries").Find(nil).All(&reports)
// 	} else {
// 		db.C("event_summaries").Find(bson.M{"user_id": u.ID}).All(&reports)
// 	}
// 	return reports
// }

