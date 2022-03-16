package routes


import(
	"github.com/gofiber/fiber/v2"
	"order-management/services"
)

func UserRoutes(app *fiber.App){

	app.Post("/create", services.CreateOrderStatusHistory)
	app.Post("/update", services.UpdateOrderStatusHistory)
	app.Get("/get", services.GetOrderStatusHistory)
}