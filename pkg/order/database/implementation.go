package database_order

import (
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DbConnection *gorm.DB
}

func (o OrderRepository) CreateOrder(data core_order.Order) (*core_order.Order, error) {
	var model OrderModel
	model.FromEntity(data)
	result := o.DbConnection.Model(OrderModel{}).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (o OrderRepository) GetOrderById(orderId string) (*core_order.Order, error) {
	var model OrderModel

	result := o.DbConnection.Model(OrderModel{}).First(&model, "order_id = ?", orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func (o OrderRepository) UpdateOrder(data core_order.Order) (*core_order.Order, error) {
	var model OrderModel
	model.FromEntity(data)
	result := o.DbConnection.Model(OrderModel{}).Where("order_id = ?", data.OrderId).Updates(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.ToEntity(), nil
}

func NewOrderRepository(module OrderRepository) core_order.IOrderRepository {
	return &module
}
