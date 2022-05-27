package model

type ProductivityLog struct {
	ID        string `json:"id"`
	Author    *User `json:"author"`
	Date      string     `json:"date"`
	Grant   *Grant `json:"grant"`
	FocusArea string        `json:"focus_area"`
	Actions   []string      `json:"actions"`
	Successes []string      `json:"successes"`
	Improvements []string    `json:"improvements"`
	NextSteps []string      `json:"next_steps"`
	Status ApprovalStatus     `json:"approval_status"`
	Notes	 []Note        `json:"notes"`
}

// func (u *User) CreateLog(args map[string]interface{}) ProductivityLog {
// 	log := ProductivityLog{}
// 	log.Status = Pending
// 	log.UserID = u.ID
// 	db.C("productivity_logs").Insert(log)
// 	return log
// }

// func (u *User) IsLogAuthor(logID bson.ObjectId) bool {
// 	filter := bson.M{"_id": logID, "user_id": u.ID}
// 	var log ProductivityLog
// 	err := db.C("productivity_logs").Find(filter).One(&log)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func (u *User) EditLog(args map[string]interface{}) {
// 	if(u.IsAdmin()) {
// 		db.C("productivity_logs").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else if(u.IsLogAuthor(args["id"].(bson.ObjectId))) {
// 		db.C("productivity_logs").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else {
// 		panic("You are not authorized to edit this log")
// 	}
// }

// func (u *User) ApproveLog(log ProductivityLog) {
// 	if(u.IsAdmin()) {
// 	db.C("productivity_logs").Update(bson.M{"_id": log.ID}, bson.M{"$set": bson.M{"approval_status": Approved}})
// 	}
// }

// func (u *User) DenyLog(log ProductivityLog) {
// 	if(u.IsAdmin()) {
// 	db.C("productivity_logs").Update(bson.M{"_id": log.ID}, bson.M{"$set": bson.M{"approval_status": Denied}})
// 	}
// }

// func (u *User) GetLogs() []ProductivityLog {
// 	if(u.IsAdmin()) {
// 		filter := nil
// 	} else {
// 		filter := bson.M{"user_id": u.ID}
// 	}
// 	var logs []ProductivityLog
// 	err := db.C("productivity_logs").Find(filter).All(&logs)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return logs
// }