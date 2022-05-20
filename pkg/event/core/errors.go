package core_event

import (
	"errors"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
)

var (
	DATABASE_ERROR = errors.New("database error")
	EXCEEDED_QUOTA = errors.New("exceeded quota")
	BOOK_IS_EXIST  = errors.New("book already processed")
)

func GetHttpError(err error) common.Response[any] {
	switch err {
	case DATABASE_ERROR:
		return common.Response[any]{
			Status:  500,
			Message: DATABASE_ERROR.Error(),
		}
	case EXCEEDED_QUOTA:
		return common.Response[any]{
			Status:  500,
			Message: EXCEEDED_QUOTA.Error(),
		}
	case BOOK_IS_EXIST:
		return common.Response[any]{
			Status:  500,
			Message: BOOK_IS_EXIST.Error(),
		}
	}

	return common.Response[any]{
		Status:  500,
		Message: "unexpected error",
	}
}
