package entity

type User struct {
	Name     string
	Email    string
	Password string
	Status   string
	Code     string
}

func NewUser(name, email, password, status, code string) User {
	return User{Name: name, Email: email, Password: password, Status: status, Code: code}
}
