package core_event

type IEventRepository interface {
	GetEventById(eventId string) (*Event, error)
	CreateEvent(data Event) (*Event, error)
	UpdateEvent(data Event) (*Event, error)
	CreateSchedule(data Schedule) (*Schedule, error)
	GetSchedulesByEventId(eventId string) ([]Schedule, error)
	GetScheduleById(scheduleId string) (*Schedule, error)
	UpdateSchedule(data Schedule) (*Schedule, error)
	CreateBook(data Book) (*Book, error)
	GetBookById(orderId string) (*Book, error)
}
