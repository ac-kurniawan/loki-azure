package http_event

import (
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"time"
)

type ScheduleRequest struct {
	StartTime  time.Time  `json:"startTime"`
	EndTime    *time.Time `json:"endTime"`
	Location   string     `json:"location"`
	BasePrice  uint64     `json:"basePrice"`
	PromoPrice *uint64    `json:"promoPrice"`
	Quota      uint       `json:"quota"`
	Booked     uint       `json:"booked"`
}

func (s *ScheduleRequest) ToEntity(eventId string) *core_event.Schedule {
	return &core_event.Schedule{
		StartTime:  s.StartTime,
		EndTime:    s.EndTime,
		Location:   s.Location,
		BasePrice:  s.BasePrice,
		PromoPrice: s.PromoPrice,
		Quota:      s.Quota,
		Booked:     s.Booked,
		EventId:    eventId,
	}
}

type EventRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublished bool   `json:"isPublished"`
}

func (e *EventRequest) ToEntity() *core_event.Event {
	return &core_event.Event{
		Name:        e.Name,
		Description: e.Description,
		IsPublished: e.IsPublished,
	}
}
