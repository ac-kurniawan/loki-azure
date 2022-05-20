package core_order

import "time"

type IOrderService interface {
	CreateOrder(data Order) (*Order, error)
	GetOrderById(orderId string) (*Order, error)
	UpdateOrder(data Order) (*Order, error)
	UpdateOrderStatus(orderId string, status string) (*Order, error)
	Checkout(orderId string) (*Order, error)
}

type OrderService struct {
	Repository      IOrderRepository
	EventRepository IEventRepository
	Utils           Utils
}

func (o OrderService) Checkout(orderId string) (*Order, error) {
	order, err := o.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	expired := order.CreatedAt.Add(300 * time.Second)
	timeNow := time.Now()
	if timeNow.After(expired) {
		o.Utils.Log.Errorf(
			"error while checkout caused, %s, timeNow: %s order time: %s", ORDER_EXPIRED,
			timeNow.String(), order.CreatedAt.String(),
		)
		o.UpdateOrderStatus(orderId, ORDER_STATUS_TIMEOUT)
		return nil, ORDER_EXPIRED
	}

	if order.Status != ORDER_STATUS_WAITING_FOR_PAYMENT {
		o.Utils.Log.Errorf(
			"error while checkout caused, %s, %v", ORDER_CANNOT_CHECKOUT,
			err,
		)
		return nil, ORDER_CANNOT_CHECKOUT
	}

	updated, err := o.UpdateOrderStatus(orderId, ORDER_STATUS_SUCCESS)
	if err != nil {
		o.Utils.Log.Errorf(
			"error while checkout caused, %s, %v", DATABASE_ERROR,
			err,
		)
		return nil, DATABASE_ERROR
	}

	_, err = o.EventRepository.OrderBooked(
		Booked{
			OrderId:    orderId,
			ScheduleId: order.ScheduleId,
			Qty:        order.Qty,
		},
	)

	if err != nil {
		o.Utils.Log.Errorf(
			"error while checkout caused, %s, %v", ORDER_CANNOT_CHECKOUT,
			err,
		)
		return nil, ORDER_CANNOT_CHECKOUT
	}

	return updated, nil
}

func (o OrderService) CreateOrder(data Order) (*Order, error) {
	schedule, err := o.EventRepository.GetScheduleById(data.ScheduleId)
	if err != nil {
		o.Utils.Log.Errorf("error while create order, %v", err)
		return nil, FAILED_TO_FETCH_EVENT
	}

	timeNow := time.Now()
	if timeNow.After(schedule.StartTime) {
		o.Utils.Log.Errorf(
			"error while creation caused, %s, createdAt: %s eventTime: %s", EVENT_EXPIRED,
			timeNow.String(), schedule.StartTime.String(),
		)
		return nil, EVENT_EXPIRED
	}

	if schedule.Booked >= schedule.Quota {
		o.Utils.Log.Errorf("error while creation caused, %s", EXCEEDED_QUOTA)
		return nil, EXCEEDED_QUOTA
	}

	data.Status = ORDER_STATUS_WAITING_FOR_PAYMENT

	result, err := o.Repository.CreateOrder(data)
	if err != nil {
		o.Utils.Log.Errorf("error while create order, %v", err)
		return nil, DATABASE_ERROR
	}

	return result, nil
}

func (o OrderService) GetOrderById(orderId string) (*Order, error) {
	result, err := o.Repository.GetOrderById(orderId)
	if err != nil {
		o.Utils.Log.Errorf("error while get order, %v", err)
		return nil, DATABASE_ERROR
	}

	return result, nil
}

func (o OrderService) UpdateOrder(data Order) (*Order, error) {
	result, err := o.Repository.UpdateOrder(data)
	if err != nil {
		o.Utils.Log.Errorf("error while update order, %v", err)
		return nil, DATABASE_ERROR
	}

	return result, nil
}

func (o OrderService) UpdateOrderStatus(orderId string, status string) (*Order, error) {
	order, err := o.Repository.GetOrderById(orderId)
	if err != nil {
		o.Utils.Log.Errorf("error while get order, %v", err)
		return nil, DATABASE_ERROR
	}

	order.Status = status
	result, err := o.Repository.UpdateOrder(*order)
	if err != nil {
		o.Utils.Log.Errorf("error while get order, %v", err)
		return nil, DATABASE_ERROR
	}

	return result, nil
}

func NewOrderService(module OrderService) IOrderService {
	return &module
}
