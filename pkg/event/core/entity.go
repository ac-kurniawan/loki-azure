package core_event

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
	Schedules   []Schedule
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Book struct {
	OrderId    string
	ScheduleId string
	Qty        uint
}
