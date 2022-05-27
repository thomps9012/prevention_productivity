package models

type ProductivityLog struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    bson.ObjectId `bson:"user_id"`
	Date      time.Time     `bson:"date"`
	GrantID   bson.ObjectId `bson:"grant_id"`
	FocusArea string        `bson:"focus_area"`
	Actions   []string      `bson:"actions"`
	Successes []string      `bson:"successes"`
	Improvements []string    `bson:"improvements"`
	NextSteps []string      `bson:"next_steps"`
	Status ApprovalStatus     `bson:"approval_status"`
	Notes	 []Note        `bson:"notes"`
}

func (u *User) CreateLog(args map[string]interface{}) ProductivityLog {
	log := ProductivityLog{}
	log.Status = Pending
	log.UserID = u.ID
	db.C("productivity_logs").Insert(log)
	return log
}

func (u *User) IsLogAuthor(logID bson.ObjectId) bool {
	filter := bson.M{"_id": logID, "user_id": u.ID}
	var log ProductivityLog
	err := db.C("productivity_logs").Find(filter).One(&log)
	if err != nil {
		return false
	}
	return true
}

func (u *User) EditLog(args map[string]interface{}) {
	if(u.IsAdmin()) {
		db.C("productivity_logs").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
	} else if(u.IsLogAuthor(args["id"].(bson.ObjectId))) {
		db.C("productivity_logs").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
	} else {
		panic("You are not authorized to edit this log")
	}
}

func (u *User) ApproveLog(log ProductivityLog) {
	if(u.IsAdmin()) {
	db.C("productivity_logs").Update(bson.M{"_id": log.ID}, bson.M{"$set": bson.M{"approval_status": Approved}})
	}
}

func (u *User) DenyLog(log ProductivityLog) {
	if(u.IsAdmin()) {
	db.C("productivity_logs").Update(bson.M{"_id": log.ID}, bson.M{"$set": bson.M{"approval_status": Denied}})
	}
}

func (u *User) GetLogs() []ProductivityLog {
	if(u.IsAdmin()) {
		filter := nil
	} else {
		filter := bson.M{"user_id": u.ID}
	}
	var logs []ProductivityLog
	err := db.C("productivity_logs").Find(filter).All(&logs)
	if err != nil {
		panic(err)
	}
	return logs
}