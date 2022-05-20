package http_order

import core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"

type OrderRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	ScheduleId  string `json:"scheduleId"`
	Qty         int    `json:"qty"`
}

func (i *OrderRequest) ToEntity() core_order.Order {
	return core_order.Order{
		PhoneNumber: i.PhoneNumber,
		Email:       i.Email,
		ScheduleId:  i.ScheduleId,
		Qty:         i.Qty,
	}
}

type CheckoutRequest struct {
	OrderId string `json:"orderId"`
}
