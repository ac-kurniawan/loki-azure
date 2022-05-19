package order

import (
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	database_order "github.com/ac-kurniawan/loki-azure/pkg/order/database"
	eventClient_order "github.com/ac-kurniawan/loki-azure/pkg/order/eventClient"
	http_order "github.com/ac-kurniawan/loki-azure/pkg/order/http"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderApplication struct {
	Fiber          *fiber.App
	GormConnection *gorm.DB
	Logrus         *logrus.Logger
}

func (a *OrderApplication) Migrate() {
	migrate := database_order.Migrations{DbConnection: a.GormConnection}
	err := migrate.Run()
	if err != nil {
		panic(err)
	}
}

func (a *OrderApplication) Run() {
	a.Logrus.Info("start order application")
	httpClient := common.GetHttpClient()

	orderUtils := core_order.Utils{
		Log: a.Logrus,
	}

	orderRepo := database_order.NewOrderRepository(
		database_order.OrderRepository{
			DbConnection: a.GormConnection,
		},
	)

	eventClient := eventClient_order.NewEventClient(
		eventClient_order.EventClient{
			BaseUrl:    "http://127.0.0.1:3222",
			HttpClient: httpClient,
		},
	)

	orderService := core_order.NewOrderService(
		core_order.OrderService{
			Repository:      orderRepo,
			EventRepository: eventClient,
			Utils:           orderUtils,
		},
	)

	orderHandler := http_order.OrderHandler{OrderService: orderService}
	orderController := http_order.OrderController{OrderHandler: orderHandler, FiberApp: a.Fiber}

	orderController.Controller()
}
