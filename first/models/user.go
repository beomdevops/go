package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; AUTO_INCREMENT"`
	Name        string
	CreditCards []CreditCard
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u *User) ToDto() *UserDto {
	return &UserDto{User_Id: u.ID, User_Name: u.Name}
}
