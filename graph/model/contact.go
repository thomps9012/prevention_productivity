package model

type Contact struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Email     string        `bson:"email"`
	Phone	 string        `bson:"phone"`
	AffiliatedOrganization string `bson:"affiliated_organization"`
	ContactType string `bson:"contact_type"`
	CreatedAt time.Time     `bson:"created_at"`
}

type ContactType string
const (
	Student,
	Parent,
	Teacher,
	NonProfit,
	Public,
	Private,
	Other
)

// func NewContact(firstName string, lastName string, email string, phone string, affiliatedOrganization string, contactType string) *Contact {
// 	contact := &Contact{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Email:     email,
// 		Phone: phone,
// 		AffiliatedOrganization: affiliatedOrganization,
// 		ContactType: contactType,
// 		CreatedAt: time.Now(),
// 	}
// 	db.C("contacts").Insert(contact)
// }

// func(u *User) EditContact(args map[string]interface{}) {
// 	if(u.IsAdmin()){
// 		db.C("contacts").Update(bson.M{"_id": args["id"]}, bson.M{"$set": args})
// 	} else {
// 		panic("You are not authorized to edit contacts.")
// 	}
// }

// func (u *User) DisplayContacts() []Contact {
// 	var contacts []Contact
// 	if(u.IsAdmin()) {
// 		err := db.C("contacts").Find(nil).All(&contacts)
// 		if err != nil {
// 			panic(err)
// 		}
// 		} else {
// 		for item := range ["Student", "Parent", "Teacher"] {
// 			err := db.C("contacts").Find(bson.M{"contact_type": item}).All(&contacts)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// 	return contacts
// }