package core_event

type IEventService interface {
	CreateEvent(data Event) (*Event, error)
	GetEventById(eventId string) (*Event, error)
	UpdateEvent(data Event) (*Event, error)
	CreateSchedule(data Schedule) (*Schedule, error)
	GetSchedulesByEventId(eventId string) ([]Schedule, error)
	GetScheduleById(scheduleId string) (*Schedule, error)
	AddBooked(data Book) (*Schedule, error)
	UpdateSchedule(data Schedule) (*Schedule, error)
}

type EventService struct {
	EventRepository IEventRepository
	Utils           Utils
}

func (e EventService) UpdateSchedule(data Schedule) (*Schedule, error) {
	if result, err := e.EventRepository.UpdateSchedule(data); err != nil {
		e.Utils.Log.Errorf("error while update schedule record: %s", err.Error())
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) AddBooked(data Book) (*Schedule, error) {
	book, err := e.EventRepository.GetBookById(data.OrderId)
	if book != nil {
		return nil, BOOK_IS_EXIST
	}

	schedule, err := e.GetScheduleById(data.ScheduleId)
	if err != nil {
		return nil, DATABASE_ERROR
	}
	afterBooked := schedule.Booked + data.Qty
	if afterBooked > schedule.Quota {
		return nil, EXCEEDED_QUOTA
	}

	schedule.Booked = afterBooked
	updated, err := e.UpdateSchedule(*schedule)
	if err != nil {
		return nil, DATABASE_ERROR
	}

	_, err = e.EventRepository.CreateBook(data)
	if err != nil {
		return nil, DATABASE_ERROR
	}

	return updated, nil
}

func (e EventService) CreateEvent(data Event) (*Event, error) {
	if result, err := e.EventRepository.CreateEvent(data); err != nil {
		e.Utils.Log.Errorf("error while create new event: %s", err.Error())
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) GetEventById(eventId string) (*Event, error) {
	if result, err := e.EventRepository.GetEventById(eventId); err != nil {
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) UpdateEvent(data Event) (*Event, error) {
	if result, err := e.EventRepository.UpdateEvent(data); err != nil {
		e.Utils.Log.Errorf("error while update event record: %s", err.Error())
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) CreateSchedule(data Schedule) (*Schedule, error) {
	if result, err := e.EventRepository.CreateSchedule(data); err != nil {
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) GetSchedulesByEventId(eventId string) ([]Schedule, error) {
	if result, err := e.EventRepository.GetSchedulesByEventId(eventId); err != nil {
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func (e EventService) GetScheduleById(scheduleId string) (*Schedule, error) {
	if result, err := e.EventRepository.GetScheduleById(scheduleId); err != nil {
		return nil, DATABASE_ERROR
	} else {
		return result, nil
	}
}

func NewEventService(module EventService) IEventService {
	return &module
}
