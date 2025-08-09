package http

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Index(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Store(*fiber.Ctx) error
	Show(*fiber.Ctx) error
	Edit(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Destroy(*fiber.Ctx) error
}
