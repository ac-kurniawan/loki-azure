package http_event

import (
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"time"
)

type ScheduleResponse struct {
	ScheduleId string     `json:"scheduleId"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    *time.Time `json:"endTime"`
	Location   string     `json:"location"`
	BasePrice  uint64     `json:"basePrice"`
	PromoPrice *uint64    `json:"promoPrice"`
	Quota      uint       `json:"quota"`
	Booked     uint       `json:"booked"`
}

func (s *ScheduleResponse) FromEntity(data core_event.Schedule) {
	s.ScheduleId = data.ScheduleId
	s.StartTime = data.StartTime
	s.EndTime = data.EndTime
	s.Location = data.Location
	s.BasePrice = data.BasePrice
	s.PromoPrice = data.PromoPrice
	s.Quota = data.Quota
	s.Booked = data.Booked
}

type EventEagerResponse struct {
	EventId     string             `json:"eventId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	IsPublished bool               `json:"isPublished"`
	Schedules   []ScheduleResponse `json:"schedules"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func (e *EventEagerResponse) FromEntity(data core_event.Event) {
	var schedules []ScheduleResponse
	for _, elm := range data.Schedules {
		convert := ScheduleResponse{}
		convert.FromEntity(elm)
		schedules = append(schedules, convert)
	}

	e.EventId = data.EventId
	e.Name = data.Name
	e.Description = data.Description
	e.IsPublished = data.IsPublished
	e.CreatedAt = data.CreatedAt
	e.UpdatedAt = data.UpdatedAt
	e.Schedules = schedules
}

type EventResponse struct {
	EventId     string    `json:"eventId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublished bool      `json:"isPublished"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (e *EventResponse) FromEntity(data core_event.Event) {
	e.EventId = data.EventId
	e.Name = data.Name
	e.Description = data.Description
	e.IsPublished = data.IsPublished
	e.CreatedAt = data.CreatedAt
	e.UpdatedAt = data.UpdatedAt
}
