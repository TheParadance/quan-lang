package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"theparadance.com/quan-lang/env"
	errorexception "theparadance.com/quan-lang/error-exception"
	lang "theparadance.com/quan-lang/quan-lang"
)

func Init() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Log the error (you can use a logger here)
			println("Error:", err.Error())

			switch e := err.(type) {
			case *errorexception.RuntimeError:
				// If it's a runtime error, return a specific JSON response
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Runtime Error",
					"message": e.Message,
				})
			case *fiber.Error:
				// If it's a fiber error, return the status code and message
				return c.Status(e.Code).JSON(fiber.Map{
					"error":   e.Message,
					"message": e.Message,
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
			}
		},
	})

	app.Use(recover.New())

	app.Post("/execute", func(c *fiber.Ctx) error {
		var request struct {
			Program string                 `json:"program"`
			Vars    map[string]interface{} `json:"vars"`
		}

		println("Request received", request.Program)

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		env, err := lang.Execuate(request.Program, &env.Env{
			Vars: request.Vars,
		})

		if err != nil {
			println("panic here")
			panic(&errorexception.RuntimeError{
				Message: err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Program executed successfully",
			"payload": map[string]interface{}{
				"program": request.Program,
				"inputs":  request.Vars,
				"outputs": env.Vars,
			},
		})
	})

	app.Listen(":3000")
}
