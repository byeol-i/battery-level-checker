package user

type User struct {
	UserInterface
	UserImpl 
}

type UserImpl struct {
	Id string `validate:"required" json:"id" example:"123"`
	Name string `validate:"required" json:"name" example:"gil dong"`
}

type UserInterface interface {
	GetId() string
	SetId(string) 
	SetName(string)
	GetName(string)
}

func NewUser() *User {
	newUser := &User{}
	return newUser
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) SetId(uid string) {
	u.Id = uid
}


func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(uName string) {
	u.Name = uName
}

