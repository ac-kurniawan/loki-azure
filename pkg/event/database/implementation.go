package database_event

import (
	"errors"
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	"gorm.io/gorm"
)

type EventRepository struct {
	DbConnection *gorm.DB
}

func (e EventRepository) CreateBook(data core_event.Book) (*core_event.Book, error) {
	var model BookModel
	model.FromEntity(data)

	result := e.DbConnection.Model(BookModel{}).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) GetBookById(orderId string) (*core_event.Book, error) {
	var model BookModel

	result := e.DbConnection.Model(BookModel{}).Limit(1).Find(&model, "order_id = ?", orderId)
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) UpdateSchedule(data core_event.Schedule) (*core_event.Schedule, error) {
	model := ScheduleModel{}
	model.FromEntity(data)

	result := e.DbConnection.Model(&ScheduleModel{}).Where("schedule_id = ?", data.ScheduleId).Updates(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) GetEventById(eventId string) (*core_event.Event, error) {
	var event EventModel
	result := e.DbConnection.Model(&EventModel{}).First(&event, "event_id = ? and is_published = ?", eventId, true)

	if result.Error != nil {
		return nil, result.Error
	}

	return event.ToEntity(), nil
}

func (e EventRepository) CreateEvent(data core_event.Event) (*core_event.Event, error) {
	model := EventModel{}
	model.FromEntity(data)

	result := e.DbConnection.Model(EventModel{}).Create(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) UpdateEvent(data core_event.Event) (*core_event.Event, error) {
	model := EventModel{}
	model.FromEntity(data)

	result := e.DbConnection.Model(&EventModel{}).Where("event_id = ?", data.EventId).Updates(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) CreateSchedule(data core_event.Schedule) (*core_event.Schedule, error) {
	model := ScheduleModel{}
	model.FromEntity(data)

	result := e.DbConnection.Model(ScheduleModel{}).Create(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (e EventRepository) GetSchedulesByEventId(eventId string) ([]core_event.Schedule, error) {
	var schedules []ScheduleModel

	result := e.DbConnection.Model(&ScheduleModel{}).Find(&schedules, "event_id = ?", eventId)

	if result.Error != nil {
		return nil, result.Error
	}

	var output []core_event.Schedule
	for _, elm := range schedules {
		output = append(output, *elm.ToEntity())
	}

	return output, nil
}

func (e EventRepository) GetScheduleById(scheduleId string) (*core_event.Schedule, error) {
	var schedule ScheduleModel
	result := e.DbConnection.Model(&ScheduleModel{}).First(&schedule, "schedule_id = ?", scheduleId)

	if result.Error != nil {
		return nil, result.Error
	}

	return schedule.ToEntity(), nil
}

func NewEventRepository(module EventRepository) core_event.IEventRepository {
	return &module
}
