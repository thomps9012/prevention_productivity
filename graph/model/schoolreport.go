package model

type SchoolReport struct {
	ID        string `json:"id"`
	Author    *User `json:"author"`
	Grant    *Grant `json:"grant"`
	Curriculum string 	  `json:"curriculum"`
	LessonPlan string 	  `json:"lesson_plan"`
	School string 	  `json:"school"`
	StudentRoster []string `json:"student_roster"`
	TopicsCovered []string `json:"topics_covered"`
	Challenges []string `json:"challenges"`
	Successes []string `json:"successes"`
	ApprovalStatus string `json:"approval_status"`
	Notes []Note `json:"notes"`
	created_at string     `json:"created_at"`
}

// func (u *User) NewReport(args map[string]interface{}) {
// 	report := &SchoolReport{}
// 	report.PreventionTeamMember = u.ID
// 	report.ApprovalStatus = "Pending"
// 	db.C("school_reports").Insert(report)
// }

// func (u *User) IsReportAuthor(reportID bson.ObjectId) bool {
// 	filter := bson.M{"_id": reportID, "user_id": u.ID}
// 	var report SchoolReport
// 	err := db.C("school_reports").Find(filter).One(&report)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func (u *User) EditSchoolReport(args map[string]interface{}) {
// 	if(u.IsAdmin()){
// 		db.C("school_reports").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else if(u.IsReportAuthor(args["id"].(bson.ObjectId))){
// 		db.C("school_reports").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else {
// 		panic("You are not authorized to edit school reports.")
// 	}
// }

// func (u *User) ApproveSchoolReport(report SchoolReport) {
// 	if(u.IsAdmin()){
// 		db.C("school_reports").Update(bson.M{"_id": report.ID}, bson.M{"$set": bson.M{"approval_status": "Approved"}})
// 	}
// }

// func (u *User) DenySchoolReport(report SchoolReport) {
// 	if(u.IsAdmin()){
// 		db.C("school_reports").Update(bson.M{"_id": report.ID}, bson.M{"$set": bson.M{"approval_status": "Denied"}})
// 	}
// }

// func (u *User) GetSchoolReports() []SchoolReport {
// 	if(u.IsAdmin()){
// 		var reports []SchoolReport
// 		err := db.C("school_reports").Find(nil).All(&reports)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return reports
// 	} else {
// 		var reports []SchoolReport
// 		err := db.C("school_reports").Find(bson.M{"user_id": u.ID}).All(&reports)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return reports
// 	}
// }