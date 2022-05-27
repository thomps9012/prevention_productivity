package models

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	Admin bool `bson:"admin"`
}

type ApprovalStatus string
const (
	Pending string = "Pending",
	Approved string = "Approved",
	Denied string = "Denied"
)

type Note struct {
	UserID bson.ObjectId `bson:"author"`
	Date time.Time `bson:"date"`
	ItemID bson.ObjectId `bson:"item_id"`
	Text string `bson:"text"`
	Section float64 `bson:"section"`
}

func NewUser(args map[string]interface{}) *User {
	user := User{}
	user.Admin = false
	db.C("users").Insert(user)
	return user
}

func (u *User) SetPassword(password string) {
	err := db.C("users").Update(bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"password": password}})
}

func (u *User) SetAdmin(admin bool) {
	err := db.C("users").Update(bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"admin": true}})
	if err != nil {
		panic(err)
	}
}

func (u *User) IsAdmin() bool {
	filter := bson.M{"_id": u.ID, "admin": true}
	var user User
	err := db.C("users").Find(filter).One(&user)
	if err != nil {
		return false
	}
	return true
}

func (u *User) GetGrants() []Grant {
	filter := bson.M{"team_members":, bson.M{{"$all", u.ID}}}
	var grants []Grant
	err := db.C("grants").Find(filter).All(&grants)
	if err != nil {
		panic(err)
	}
	return grants
}

func (u *User) CreateNote(item_id bson.ObjectId, section float64, text string) Note {
	note := Note{
		UserID: u.ID,
		Date: time.Now(),
		ItemID: item_id,
		Section: section,
		Text: text,
	}
	db.C("notes").Insert(note)
}

func GetNotes(args map[string]interface{}) []Note {
	var notes []Note
	db.C("notes").Find(args).All(&notes)
	return notes
}