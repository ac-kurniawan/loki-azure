package database_order

import "gorm.io/gorm"

type Migrations struct {
	DbConnection *gorm.DB
}

func (m *Migrations) Run() error {
	err := m.DbConnection.AutoMigrate(&OrderModel{})

	if err != nil {
		return err
	}

	return nil
}
