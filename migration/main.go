package main

import (
	"fmt"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	config2 "github.com/ac-kurniawan/loki-azure/pkg/config"
	"github.com/ac-kurniawan/loki-azure/pkg/event"
)

func main() {
	fmt.Println("start migrations")
	config := config2.Config{}
	config.GetConfig()

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
	defer common.ClosePSQL(eventPsqlConnection)

	eventApp := event.EventApplication{
		GormConnection: eventPsqlConnection,
	}

	eventPsqlConnection.Exec(
		fmt.Sprintf(
			"CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION %s", eventSchemaName, config.PSQL_USERNAME,
		),
	)
	eventApp.Migrate()

	fmt.Println("finish migrations")
}
