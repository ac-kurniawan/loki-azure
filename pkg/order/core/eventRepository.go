package core_order

import "time"

type Schedule struct {
	ScheduleId string
	StartTime  time.Time
	EndTime    *time.Time
	Location   string
	BasePrice  uint64
	PromoPrice *uint64
	Quota      uint
	Booked     uint
	EventId    string
}

type Event struct {
	EventId     string
	Name        string
	Description string
	IsPublished bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Booked struct {
	OrderId    string
	ScheduleId string
	Qty        int
}

type IEventRepository interface {
	GetScheduleById(scheduleId string) (*Schedule, error)
	GetEventById(eventId string) (*Event, error)
	OrderBooked(data Booked) (*Schedule, error)
}
