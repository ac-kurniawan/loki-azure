package http_event

import (
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	EventHandler EventHandler
	FiberApp     *fiber.App
}

func (c *EventController) Controller() {
	routes := c.FiberApp.Group("/event")

	routes.Get(
		"/health", func(ctx *fiber.Ctx) error {
			return ctx.Status(200).SendString("ok")
		},
	)

	routes.Post("/", c.EventHandler.CreateEvent)
	routes.Post("/:eventId/schedule", c.EventHandler.CreateSchedule)
	routes.Get("/:eventId", c.EventHandler.GetEventById)
	routes.Get("/:eventId/schedule", c.EventHandler.GetSchedulesByEventId)
	routes.Get("/schedule/:scheduleId", c.EventHandler.GetScheduleById)
	routes.Post("/schedule/booked", c.EventHandler.Booked)
}
