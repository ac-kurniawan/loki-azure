package eventClient_order

import core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"

type BookedRequest struct {
	OrderId    string `json:"orderId"`
	ScheduleId string `json:"scheduleId"`
	Qty        int    `json:"qty"`
}

func (b *BookedRequest) FromEntity(data core_order.Booked) {
	b.ScheduleId = data.ScheduleId
	b.OrderId = data.OrderId
	b.Qty = data.Qty
}
