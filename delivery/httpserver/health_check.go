package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s Server) healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "everything is good! " + time.Now().String(),
	})
}
