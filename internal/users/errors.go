package users

type WrongEmailOrPassword struct {}

func (m *WrongEmailOrPassword) Error() string {
	return "Wrong email or password"
}

type UserNotActive struct {}

func (m *UserNotActive) Error() string {
	return "User not active"
}