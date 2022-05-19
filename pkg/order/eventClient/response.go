package eventClient_order

import (
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"time"
)

type ScheduleResponse struct {
	ScheduleId string     `json:"scheduleId"`
	StartTime  time.Time  `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime"`
	Location   string     `json:"location"`
	BasePrice  uint64     `json:"basePrice"`
	PromoPrice *uint64    `json:"promoPrice,omitempty"`
	Quota      uint       `json:"quota"`
	Booked     uint       `json:"booked"`
	EventId    string     `json:"eventId"`
}

func (s *ScheduleResponse) ToEntity() *core_order.Schedule {
	return &core_order.Schedule{
		ScheduleId: s.ScheduleId,
		StartTime:  s.StartTime,
		EndTime:    s.EndTime,
		Location:   s.Location,
		BasePrice:  s.BasePrice,
		PromoPrice: s.PromoPrice,
		Quota:      s.Quota,
		Booked:     s.Booked,
		EventId:    s.EventId,
	}
}

type EventResponse struct {
	EventId     string    `json:"eventId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublished bool      `json:"isPublished"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (e *EventResponse) ToEntity() *core_order.Event {
	return &core_order.Event{
		EventId:     e.EventId,
		Name:        e.Name,
		Description: e.Description,
		IsPublished: e.IsPublished,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
