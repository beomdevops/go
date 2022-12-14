package service

import (
	"park/goproject/first/models"
	"park/goproject/first/repository"
)

type CardService struct {
	cardRepository repository.CardRepository
	userRepository repository.UserRepository
}

func NewCardService(
	di_cardRepository repository.CardRepository,
	di_userRepo repository.UserRepository,
) *CardService {

	return &CardService{
		cardRepository: di_cardRepository,
		userRepository: di_userRepo,
	}
}

func (c *CardService) FindByCardId(card_id int) (*models.CreditCard, error) {
	data, err := c.cardRepository.FindByCardId(card_id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CardService) FindByUserId(user_id int) (*models.CreditCard, error) {
	data, err := c.cardRepository.FindByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CardService) CreateCard(user_id int) (*models.CreditCard, error) {

	user, err := c.userRepository.FindById(user_id)
	if err != nil {
		return nil, err
	}

	card := models.NewCreditCard(int(user.ID))

	card, err = c.cardRepository.CreateCard(card)

	if err != nil {
		return nil, err
	}

	return card, nil
}
