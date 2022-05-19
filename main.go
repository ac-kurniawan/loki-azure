package main

import (
	"fmt"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	config2 "github.com/ac-kurniawan/loki-azure/pkg/config"
	"github.com/ac-kurniawan/loki-azure/pkg/event"
	"github.com/ac-kurniawan/loki-azure/pkg/order"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	config := config2.Config{}
	config.GetConfig()

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	log.Info("start server")

	eventSchemaName := "event"
	eventPsqlConnection, err := common.NewPSQLConnection(
		config.PSQL_HOST,
		config.PSQL_PORT,
		config.PSQL_USERNAME,
		config.PSQL_PASSWORD,
		config.PSQL_DB_NAME,
		&eventSchemaName,
	)
	if err != nil {
		panic("Error connect to db")
	}
	orderSchemaName := "orders"
	orderPsqlConnection, err := common.NewPSQLConnection(
		config.PSQL_HOST,
		config.PSQL_PORT,
		config.PSQL_USERNAME,
		config.PSQL_PASSWORD,
		config.PSQL_DB_NAME,
		&orderSchemaName,
	)
	if err != nil {
		panic("Error connect to db")
	}

	fiberApp := fiber.New()
	fiberApp.Use(recover.New())
	fiberApp.Use(recover.New())
	fiberApp.Use(requestid.New(requestid.ConfigDefault))

	eventApp := event.EventApplication{
		Fiber:          fiberApp,
		GormConnection: eventPsqlConnection,
		Logrus:         log,
	}

	orderApp := order.OrderApplication{
		Fiber:          fiberApp,
		GormConnection: orderPsqlConnection,
		Logrus:         log,
	}

	eventApp.Run()
	orderApp.Run()

	// run server
	go func() {
		if err := fiberApp.Listen(":" + config.ServerPort); err != nil {
			fmt.Printf("[HTTP Server] - %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := fiberApp.Shutdown(); err != nil {
		fmt.Printf("[HTTP Server] - %s", err.Error())
	}
}
