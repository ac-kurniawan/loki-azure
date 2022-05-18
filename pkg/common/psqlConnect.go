package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewPSQLConnection(
	host string, port string, username string,
	password string, dbname string, schemaName *string,
) (*gorm.DB, error) {
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname" +
		"=" + dbname + " port=" + port + " sslmode=disable"

	if schemaName == nil {
		defaultSchemaName := "public."
		schemaName = &defaultSchemaName
	}
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   *schemaName + ".",
				SingularTable: false,
			},
			Logger: logger.Default.LogMode(logger.Error),
		},
	)
	if err != nil {
		return nil, err
	}

	psql, err := db.DB()
	if err != nil {
		return nil, err
	}
	psql.SetConnMaxLifetime(300)
	psql.SetConnMaxIdleTime(3600)
	psql.SetMaxOpenConns(10)
	psql.SetMaxIdleConns(100)
	psql.Ping()
	return db, nil
}

func ClosePSQL(db *gorm.DB) error {
	psql, err := db.DB()
	if err != nil {
		return err
	}
	err = psql.Close()
	if err != nil {
		return err
	}
	return nil
}
