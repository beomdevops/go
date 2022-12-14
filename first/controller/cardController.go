package controller

import (
	"park/goproject/first/service"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

type CardController struct {
	cardService *service.CardService
}

func NewCardController(di_cardService *service.CardService) *CardController {

	return &CardController{cardService: di_cardService}

}

func (cardController *CardController) FindByCardId(ctx *fiber.Ctx) error {
	parma := ctx.Params("cardId")
	id, err := strconv.Atoi(parma)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	find_card, err := cardController.cardService.FindByCardId(id)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_card})

}

func (cardController *CardController) FindByUserId(ctx *fiber.Ctx) error {
	parma := ctx.Params("userId")
	id, _ := strconv.Atoi(parma)
	find_card, err := cardController.cardService.FindByUserId(id)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_card})
}

type CardCreateRequest struct {
	UserID int `json:"id"`
}

func (cardController *CardController) CreateCard(ctx *fiber.Ctx) error {

	p := new(CardCreateRequest)
	err := ctx.BodyParser(p)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}
	create_card, err := cardController.cardService.CreateCard(p.UserID)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": create_card})
}
