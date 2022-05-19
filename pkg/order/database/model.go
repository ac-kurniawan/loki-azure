package database_order

import (
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"github.com/google/uuid"
	"time"
)

type OrderModel struct {
	OrderId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;index"`
	PhoneNumber string
	Email       string
	Status      string
	ScheduleId  string
	Qty         int
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
}

func (o *OrderModel) FromEntity(data core_order.Order) {
	o.OrderId, _ = uuid.Parse(data.OrderId)
	o.PhoneNumber = data.PhoneNumber
	o.Email = data.Email
	o.Status = data.Status
	o.ScheduleId = data.ScheduleId
	o.Qty = data.Qty
	o.CreatedAt = data.CreatedAt
	o.UpdateAt = data.UpdateAt
}

func (o *OrderModel) ToEntity() *core_order.Order {
	return &core_order.Order{
		OrderId:     o.OrderId.String(),
		PhoneNumber: o.PhoneNumber,
		Email:       o.Email,
		Status:      o.Status,
		ScheduleId:  o.ScheduleId,
		Qty:         o.Qty,
		CreatedAt:   o.CreatedAt,
		UpdateAt:    o.UpdateAt,
	}
}
