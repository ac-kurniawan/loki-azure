package event

import (
	core_event "github.com/ac-kurniawan/loki-azure/pkg/event/core"
	database_event "github.com/ac-kurniawan/loki-azure/pkg/event/database"
	http_event "github.com/ac-kurniawan/loki-azure/pkg/event/http"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EventApplication struct {
	Fiber          *fiber.App
	GormConnection *gorm.DB
	Logrus         *logrus.Logger
}

func (a *EventApplication) Migrate() {
	migrate := database_event.Migrations{DbConnection: a.GormConnection}
	err := migrate.Run()
	if err != nil {
		panic(err)
	}
}

func (a *EventApplication) Run() {
	a.Logrus.Info("start event application")
	eventUtils := core_event.Utils{
		Log: a.Logrus,
	}

	eventRepo := database_event.NewEventRepository(database_event.EventRepository{DbConnection: a.GormConnection})

	eventService := core_event.NewEventService(
		core_event.EventService{
			EventRepository: eventRepo,
			Utils:           eventUtils,
		},
	)

	eventHandler := http_event.EventHandler{EventService: eventService}
	eventController := http_event.EventController{EventHandler: eventHandler, FiberApp: a.Fiber}

	eventController.Controller()
}
