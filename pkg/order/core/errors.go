package core_order

import (
	"errors"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
)

var (
	DATABASE_ERROR        = errors.New("database error")
	FAILED_TO_FETCH_EVENT = errors.New("failed to fetch to event service")
	EXCEEDED_QUOTA        = errors.New("exceeded quota")
	EVENT_EXPIRED         = errors.New("event expired")
	ORDER_EXPIRED         = errors.New("order expired")
	ORDER_CANNOT_CHECKOUT = errors.New("order cannot checkout, because already timeout/cancel/success")
)

func GetHttpError(err error) common.Response[any] {
	switch err {
	case DATABASE_ERROR:
		return common.Response[any]{
			Status:  500,
			Message: DATABASE_ERROR.Error(),
		}
	case FAILED_TO_FETCH_EVENT:
		return common.Response[any]{
			Status:  500,
			Message: FAILED_TO_FETCH_EVENT.Error(),
		}
	case EXCEEDED_QUOTA:
		return common.Response[any]{
			Status:  500,
			Message: EXCEEDED_QUOTA.Error(),
		}
	case EVENT_EXPIRED:
		return common.Response[any]{
			Status:  500,
			Message: EVENT_EXPIRED.Error(),
		}
	case ORDER_EXPIRED:
		return common.Response[any]{
			Status:  500,
			Message: ORDER_EXPIRED.Error(),
		}
	case ORDER_CANNOT_CHECKOUT:
		return common.Response[any]{
			Status:  500,
			Message: ORDER_CANNOT_CHECKOUT.Error(),
		}

	}

	return common.Response[any]{
		Status:  500,
		Message: "unexpected error",
	}
}
