package domain

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) GetId() string {
	return u.Id
}
