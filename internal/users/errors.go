package users

type WrongEmailOrPassword struct {}

func (m *WrongEmailOrPassword) Error() string {
	return "Wrong email or password"
}