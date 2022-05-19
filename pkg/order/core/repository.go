package core_order

type IOrderRepository interface {
	CreateOrder(data Order) (*Order, error)
	GetOrderById(orderId string) (*Order, error)
	UpdateOrder(data Order) (*Order, error)
}
