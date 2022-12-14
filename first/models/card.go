package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	ID     uint `gorm:"primaryKey; AUTO_INCREMENT"`
	Number string
	UserID int
}

func NewCreditCard(user_id int) *CreditCard {
	card := &CreditCard{UserID: user_id}
	card.Number = card.getNumber()
	return card
}

func (c *CreditCard) getNumber() string {
	return uuid.New().String()
}
