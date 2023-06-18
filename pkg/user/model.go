package user

import (
	pb_unit_user "github.com/byeol-i/battery-level-checker/pb/unit/user"
)

type User struct {
	UserInterface
	UserImpl 
}

type Token struct {
	Uid string 
	Token string
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
	ToProto() (*pb_unit_user.User, error)
}

func NewUser() *User {
	newUser := &User{}
	return newUser
}

func NewUserFromProto(pbUser *pb_unit_user.User) (*User, error) {
	newUser := &User{
		UserImpl: UserImpl{
			Id: pbUser.Id.Id,
			Name: pbUser.Name,
		},	
	}

	err := UserValidator(&newUser.UserImpl)
	if err != nil {
		return nil, err
	}
	
	return newUser, nil
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

func (u *User) ToProto() (*pb_unit_user.User, error) {
	pbUnit := &pb_unit_user.User{}

	id := &pb_unit_user.ID{
		Id: u.Id,
	}

	pbUnit.Id = id
	pbUnit.Name = u.Name

	return pbUnit, nil
}