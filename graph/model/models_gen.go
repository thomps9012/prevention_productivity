package model

type AllLogs struct {
	Log       *Log  `json:"log"`
	User      *User `json:"user"`
	NoteCount *int  `json:"noteCount"`
}

type Log struct {
	ID           *string `json:"id" bson:"_id"`
	UserID       *string `json:"user_id" bson:"user_id"`
	FocusArea    string  `json:"focus_area" bson:"focus_area"`
	Actions      string  `json:"actions"`
	Successes    string  `json:"successes"`
	Improvements string  `json:"improvements"`
	NextSteps    string  `json:"next_steps" bson:"next_steps"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at" bson:"created_at"`
	UpdatedAt    string  `json:"updated_at" bson:"updated_at"`
}

type LogWithNotes struct {
	Log   *Log    `json:"log"`
	Notes []*Note `json:"notes"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewLog struct {
	FocusArea    string `json:"focus_area" bson:"focus_area"`
	Actions      string `json:"actions"`
	Successes    string `json:"successes"`
	Improvements string `json:"improvements"`
	NextSteps    string `json:"next_steps" bson:"next_steps"`
}

type NewNote struct {
	ItemID  string `json:"item_id" bson:"item_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Note struct {
	ID        *string `json:"id" bson:"_id"`
	ItemID    *string `json:"item_id" bson:"item_id"`
	UserID    *string `json:"user_id" bson:"user_id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UpdateLog struct {
	FocusArea    string `json:"focus_area" bson:"focus_area"`
	Actions      string `json:"actions"`
	Successes    string `json:"successes"`
	Improvements string `json:"improvements"`
	NextSteps    string `json:"next_steps" bson:"next_steps"`
	Status       string `json:"status"`
}

type UpdateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	ID        *string `json:"id" bson:"_id"`
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	IsAdmin   bool    `json:"is_admin"`
}
