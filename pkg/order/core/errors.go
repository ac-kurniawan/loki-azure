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
	}

	return common.Response[any]{
		Status:  500,
		Message: "unexpected error",
	}
}
