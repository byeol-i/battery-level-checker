package user

type User struct {
	UserInterface
}

type UserImpl struct {
	Id string
}

type UserInterface interface {
	GetId() string
	SetId(string) 
}

func NewUser(uid string) *User {
	return &User{}
}

func (u *UserImpl) GetId() string {
	return u.Id
}

func (u *UserImpl) SetId(uid string) {
	u.Id = uid
}