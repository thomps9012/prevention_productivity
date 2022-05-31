package model

type Contact struct {
	ID        string `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone	 string        `json:"phone"`
	AffiliatedOrganization string `json:"affiliated_organization"`
	ContactType string `json:"contact_type"`
	CreatedAt string     `json:"created_at"`
}

// type ContactType string
// const (
// 	Student string = "Student"
// 	Parent string = "Parent"
// 	Teacher string = "Teacher"
// 	NonProfit string = "Non-Profit"
// 	Public string = "Public"
// 	Private string = "Private"
// 	Other string = "Other"
// )

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