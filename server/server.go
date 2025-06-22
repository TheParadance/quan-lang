package server

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/recover"
// 	lang "theparadance.com/quan-lang/quan-lang"
// 	builtinfunc "theparadance.com/quan-lang/src/builtin-func"
// 	debuglevel "theparadance.com/quan-lang/src/debug/debug-level"
// 	"theparadance.com/quan-lang/src/env"
// 	errorexception "theparadance.com/quan-lang/src/error-exception"
// 	systemconsole "theparadance.com/quan-lang/src/system-console"
// )

// func Init() {

// 	app := fiber.New(fiber.Config{
// 		ErrorHandler: func(c *fiber.Ctx, err error) error {
// 			// Log the error (you can use a logger here)
// 			println("Error:", err.Error())

// 			switch e := err.(type) {
// 			case errorexception.QuanLangEngineError:
// 				// If it's a runtime error, return a specific JSON response
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error":   "QuanLang Engine error",
// 					"message": e.GetMessage(),
// 				})
// 			case *fiber.Error:
// 				// If it's a fiber error, return the status code and message
// 				return c.Status(e.Code).JSON(fiber.Map{
// 					"error":   e.Message,
// 					"message": e.Message,
// 				})
// 			default:
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
// 			}
// 		},
// 	})
// 	app.Use(recover.New())
// 	app.Post("/execute", func(c *fiber.Ctx) error {
// 		var request struct {
// 			Program string                 `json:"program"`
// 			Vars    map[string]interface{} `json:"vars"`
// 		}

// 		println("Request received", request.Program)

// 		if err := c.BodyParser(&request); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 		}

// 		debugLv := []debuglevel.DebugLevel{debuglevel.LEXER_TOKENS, debuglevel.AST_TREE}
// 		console := systemconsole.NewVirtualSystemConsole()
// 		langOptions := lang.NewExecuationOption(console, lang.RELEASE_MODE, &debugLv)
// 		e := &env.Env{
// 			Vars:    request.Vars,
// 			Builtin: builtinfunc.BuildInFuncs(console),
// 		}
// 		result, err := lang.Execuate(request.Program, e, langOptions)

// 		if err != nil {
// 			println("panic here")
// 			panic(err)
// 		}

// 		return c.JSON(fiber.Map{
// 			"message": "Program executed successfully",
// 			"payload": map[string]interface{}{
// 				"program": request.Program,
// 				"inputs":  request.Vars,
// 				"outputs": result.Env.Vars,
// 				"console": result.ConsoleMessages,
// 				"tokens":  result.Tokens,
// 				"ast":     result.Expression,
// 			},
// 		})
// 	})

// 	app.Listen(":3000")
// }
