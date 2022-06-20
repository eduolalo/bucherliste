package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

func SignUp(c *fiber.Ctx) error {

	res := common.Response{}
	res.Created("")
	return c.Status(fiber.StatusCreated).JSON(&res)
}
