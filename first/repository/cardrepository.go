package repository

import (
	"errors"
	"park/goproject/first/models"

	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(c *models.CreditCard) (*models.CreditCard, error)
	FindByCardId(card_id int) (*models.CreditCard, error)
	FindByUserId(uesr_id int) (*models.CreditCard, error)
}

func NewCardRepository(cdb *gorm.DB) CardRepository {
	return &cardRepository{db: cdb}
}

type cardRepository struct {
	db *gorm.DB
}

func (repo *cardRepository) CreateCard(c *models.CreditCard) (*models.CreditCard, error) {

	result := repo.db.Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil

}

func (repo *cardRepository) FindByCardId(card_id int) (*models.CreditCard, error) {

	data := &models.CreditCard{}

	result := repo.db.Find(data, "id = ?", card_id)

	if result.RowsAffected < 1 {
		return nil, errors.New("not found")
	}
	return data, nil
}
func (repo *cardRepository) FindByUserId(user_id int) (*models.CreditCard, error) {
	data := &models.CreditCard{}
	result := repo.db.Where("user_id = ?", user_id).First(data)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found")
	}
	return data, nil
}
