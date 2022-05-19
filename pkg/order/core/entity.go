package core_order

import "time"

type Order struct {
	OrderId     string
	PhoneNumber string
	Email       string
	Status      string
	ScheduleId  string
	Qty         int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
