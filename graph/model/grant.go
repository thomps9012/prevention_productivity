package model

type Grant struct {
	ID        string `json:"id"`
	AwardNumber string        `json:"award_number"`
	Title      string        `json:"title"`
	Description string        `json:"description"`
	StartDate  string     `json:"start_date"`
	EndDate    string     `json:"end_date"`
	Status     string        `json:"status"`
	TeamMembers []User `json:"team_members"`
}


// func NewGrant(awardNumber string, title string, description string, startDate string, endDate string) *Grant {
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