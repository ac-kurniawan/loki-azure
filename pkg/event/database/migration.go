package database_event

import "gorm.io/gorm"

type Migrations struct {
	DbConnection *gorm.DB
}

func (m *Migrations) Run() error {
	err := m.DbConnection.AutoMigrate(&EventModel{}, &ScheduleModel{})

	if err != nil {
		return err
	}

	return nil
}
