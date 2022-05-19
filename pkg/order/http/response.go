package http_order

import (
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"time"
)

type OrderResponse struct {
	OrderId     string    `json:"orderId"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	ScheduleId  string    `json:"scheduleId"`
	Qty         int       `json:"qty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func (i *OrderResponse) FromEntity(data core_order.Order) {
	i.OrderId = data.OrderId
	i.PhoneNumber = data.PhoneNumber
	i.Email = data.Email
	i.Status = data.Status
	i.ScheduleId = data.ScheduleId
	i.Qty = data.Qty
	i.CreatedAt = data.CreatedAt
	i.UpdateAt = data.UpdateAt
}
