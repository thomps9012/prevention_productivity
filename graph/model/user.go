package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Admin bool `json:"admin"`
}

type ApprovalStatus string
const (
	PENDING string = "PENDING"
	APPROVED string = "APPROVED"
	REJECTED string = "REJECTED"
)


// func NewUser(args map[string]interface{}) *User {
// 	user := User{}
// 	user.Admin = false
// 	db.C("users").Insert(user)
// 	return user
// }

// func (u *User) SetPassword(password string) {
// 	err := db.C("users").Update(bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"password": password}})
// }

// func (u *User) SetAdmin(admin bool) {
// 	err := db.C("users").Update(bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"admin": true}})
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (u *User) IsAdmin() bool {
// 	filter := bson.M{"_id": u.ID, "admin": true}
// 	var user User
// 	err := db.C("users").Find(filter).One(&user)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func (u *User) GetGrants() []Grant {
// 	filter := bson.M{"team_members":, bson.M{{"$all", u.ID}}}
// 	var grants []Grant
// 	err := db.C("grants").Find(filter).All(&grants)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return grants
// }

// func (u *User) CreateNote(item_id bson.ObjectId, section float64, text string) Note {
// 	note := Note{
// 		UserID: u.ID,
// 		Date: time.Now(),
// 		ItemID: item_id,
// 		Section: section,
// 		Text: text,
// 	}
// 	db.C("notes").Insert(note)
// }

// func GetNotes(args map[string]interface{}) []Note {
// 	var notes []Note
// 	db.C("notes").Find(args).All(&notes)
// 	return notes
// }