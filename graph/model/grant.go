package model

type Grant struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	AwardNumber string        `bson:"award_number"`
	Title      string        `bson:"title"`
	Description string        `bson:"description"`
	StartDate  time.Time     `bson:"start_date"`
	EndDate    time.Time     `bson:"end_date"`
	Status     string        `bson:"status"`
	TeamMembers []User `bson:"team_members"`
}


// func NewGrant(awardNumber string, title string, description string, startDate time.Time, endDate time.Time) *Grant {
// 	grant := &Grant{
// 		AwardNumber: awardNumber,
// 		Title: title,
// 		Description: description,
// 		StartDate: startDate,
// 		EndDate: endDate,
// 		Status: "Active"
// 	}
// 	db.C("grants").Insert(grant)
// 	return grant
// }

// func (u *User) EditGrant(args map[string]interface{}) {
// 	if(u.IsAdmin()){
// 		db.C("grants").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 		} else {
// 			panic("You are not authorized to edit grants.")
// 		}
// }

// func (g *Grant) GetProductivityLogs() []ProductivityLog {
// 	filter := bson.M{"grant_id": g.ID}
// 	var logs []ProductivityLog
// 	err := db.C("productivity_logs").Find(filter).All(&logs)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return logs
// }