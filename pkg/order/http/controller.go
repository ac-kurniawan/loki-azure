package http_order

import "github.com/gofiber/fiber/v2"

type OrderController struct {
	OrderHandler OrderHandler
	FiberApp     *fiber.App
}

func (o *OrderController) Controller() {
	routes := o.FiberApp.Group("/order")

	routes.Get(
		"/health", func(ctx *fiber.Ctx) error {
			return ctx.Status(200).SendString("ok")
		},
	)

	routes.Post("/", o.OrderHandler.CreateOrder)
	routes.Get("/:orderId", o.OrderHandler.GetOrderById)
}
