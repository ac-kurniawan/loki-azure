package http_order

import (
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderService core_order.IOrderService
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var request OrderRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	event, err := h.OrderService.CreateOrder(request.ToEntity())
	if err != nil {
		return c.Status(500).JSON(core_order.GetHttpError(err))
	}
	var response OrderResponse
	response.FromEntity(*event)

	return c.Status(201).JSON(
		common.Response[OrderResponse]{
			Status: 201,
			Data:   response,
		},
	)
}

func (h *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	event, err := h.OrderService.GetOrderById(orderId)
	if err != nil {
		return c.Status(500).JSON(core_order.GetHttpError(err))
	}

	var response OrderResponse
	response.FromEntity(*event)
	return c.Status(200).JSON(
		common.Response[OrderResponse]{
			Status: 200,
			Data:   response,
		},
	)
}
