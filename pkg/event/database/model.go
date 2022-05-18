package database_event

import (
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"github.com/google/uuid"
	"time"
)

type ScheduleModel struct {
	ScheduleId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;index"`
	StartTime  time.Time
	EndTime    *time.Time
	Location   string
	BasePrice  uint64 `gorm:"not null"`
	PromoPrice *uint64
	Quota      uint
	Booked     uint
	EventId    string
}

type EventModel struct {
	EventId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;index"`
	Name        string    `gorm:"not null"`
	Description string
	IsPublished bool
	Schedules   []ScheduleModel `gorm:"foreignKey:EventId;references:EventId"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}

func (s *ScheduleModel) ToEntity() *core_event.Schedule {
	return &core_event.Schedule{
		ScheduleId: s.ScheduleId.String(),
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

func (s *ScheduleModel) FromEntity(data core_event.Schedule) {
	s.ScheduleId, _ = uuid.Parse(data.ScheduleId)
	s.StartTime = data.StartTime
	s.EndTime = data.EndTime
	s.Location = data.Location
	s.BasePrice = data.BasePrice
	s.PromoPrice = data.PromoPrice
	s.Quota = data.Quota
	s.Booked = data.Booked
	s.EventId = data.EventId
}

func (e *EventModel) ToEntity() *core_event.Event {
	var schedules []core_event.Schedule
	for _, elm := range e.Schedules {
		convert := elm.ToEntity()
		schedules = append(schedules, *convert)
	}
	return &core_event.Event{
		EventId:     e.EventId.String(),
		Name:        e.Name,
		Description: e.Description,
		IsPublished: e.IsPublished,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Schedules:   schedules,
	}
}

func (e *EventModel) FromEntity(data core_event.Event) {
	var schedules []ScheduleModel
	for _, elm := range data.Schedules {
		convert := ScheduleModel{}
		convert.FromEntity(elm)
		schedules = append(schedules, convert)
	}

	e.EventId, _ = uuid.Parse(data.EventId)
	e.Name = data.Name
	e.Description = data.Description
	e.IsPublished = data.IsPublished
	e.CreatedAt = data.CreatedAt
	e.UpdatedAt = data.UpdatedAt
	e.Schedules = schedules
}
