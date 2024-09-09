package facing

import (
	"github.com/gofiber/fiber/v2"
	add_item_use_case "github.com/isidroevc/blist_api/domain/usecase"
	"github.com/isidroevc/blist_api/facing/models"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

func BuildApp() *fiber.App {
	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("items/", createItem)

	return app
}

func createItem(c *fiber.Ctx) error {

	input := new(models.CreateItemInput)

	if err := c.BodyParser(input); err != nil {
		return err
	}
	createItemInput, err := input.ToCreateItemInput()
	if err != nil {
		return err
	}
	result, err := add_item_use_case.CreateItem(createItemInput)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
