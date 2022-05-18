package core_event

type IEventRepository interface {
	GetEventById(eventId string) (*Event, error)
	CreateEvent(data Event) (*Event, error)
	UpdateEvent(data Event) (*Event, error)
	CreateSchedule(data Schedule) (*Schedule, error)
	GetSchedulesByEventId(eventId string) ([]Schedule, error)
	GetScheduleById(scheduleId string) (*Schedule, error)
}
