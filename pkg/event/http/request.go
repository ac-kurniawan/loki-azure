package http_event

import (
	"fmt"
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"time"
)

type ScheduleRequest struct {
	StartTime  string  `json:"startTime"`
	EndTime    string  `json:"endTime,omitempty"`
	Location   string  `json:"location"`
	BasePrice  uint64  `json:"basePrice"`
	PromoPrice *uint64 `json:"promoPrice,omitempty"`
	Quota      uint    `json:"quota"`
	Booked     uint    `json:"booked,omitempty"`
}

func (s *ScheduleRequest) ToEntity(eventId string) *core_event.Schedule {
	startTime, err := time.Parse(time.RFC3339, s.StartTime)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	var endTime *time.Time
	if s.EndTime != "" {
		datetime, err := time.Parse(time.RFC3339, s.EndTime)
		endTime = &datetime
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	return &core_event.Schedule{
		StartTime:  startTime,
		EndTime:    endTime,
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
