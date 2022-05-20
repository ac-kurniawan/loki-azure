package http_event

import (
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	EventService core_event.IEventService
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var request EventRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	event, err := h.EventService.CreateEvent(*request.ToEntity())
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}
	var response EventResponse
	response.FromEntity(*event)

	return c.Status(201).JSON(
		common.Response[EventResponse]{
			Status: 201,
			Data:   response,
		},
	)
}

func (h *EventHandler) GetEventById(c *fiber.Ctx) error {
	eventId := c.Params("eventId")
	event, err := h.EventService.GetEventById(eventId)
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}

	var response EventResponse
	response.FromEntity(*event)
	return c.Status(200).JSON(
		common.Response[EventResponse]{
			Status: 200,
			Data:   response,
		},
	)
}

func (h *EventHandler) CreateSchedule(c *fiber.Ctx) error {
	eventId := c.Params("eventId")
	var request ScheduleRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(
			common.Response[error]{
				Status:  400,
				Message: err.Error(),
			},
		)
	}

	schedule, err := h.EventService.CreateSchedule(*request.ToEntity(eventId))
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}

	var response ScheduleResponse
	response.FromEntity(*schedule)
	return c.Status(201).JSON(
		common.Response[ScheduleResponse]{
			Status: 201,
			Data:   response,
		},
	)
}

func (h *EventHandler) GetSchedulesByEventId(c *fiber.Ctx) error {
	eventId := c.Params("eventId")

	schedules, err := h.EventService.GetSchedulesByEventId(eventId)
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}

	var responses []ScheduleResponse
	for _, elm := range schedules {
		var response ScheduleResponse
		response.FromEntity(elm)
		responses = append(responses, response)
	}

	return c.Status(200).JSON(
		common.Response[[]ScheduleResponse]{
			Status: 200,
			Data:   responses,
		},
	)
}

func (h *EventHandler) GetScheduleById(c *fiber.Ctx) error {
	scheduleId := c.Params("scheduleId")
	schedule, err := h.EventService.GetScheduleById(scheduleId)
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}

	var response ScheduleResponse
	response.FromEntity(*schedule)
	return c.Status(200).JSON(
		common.Response[ScheduleResponse]{
			Status: 200,
			Data:   response,
		},
	)
}

func (h *EventHandler) Booked(c *fiber.Ctx) error {
	var request BookedRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	schedule, err := h.EventService.AddBooked(request.ToEntity())
	if err != nil {
		return c.Status(500).JSON(core_event.GetHttpError(err))
	}
	var response ScheduleResponse
	response.FromEntity(*schedule)

	return c.Status(201).JSON(
		common.Response[ScheduleResponse]{
			Status: 201,
			Data:   response,
		},
	)
}
