package core_event

type IEventService interface {
	CreateEvent(data Event) (*Event, error)
	GetEventById(eventId string) (*Event, error)
	UpdateEvent(data Event) (*Event, error)
	CreateSchedule(data Schedule) (*Schedule, error)
	GetSchedulesByEventId(eventId string) ([]Schedule, error)
	GetScheduleById(scheduleId string) (*Schedule, error)
}

type EventService struct {
	EventRepository IEventRepository
	Utils           Utils
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
