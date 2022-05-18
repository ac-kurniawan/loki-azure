package core_event

import (
	"errors"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
)

var (
	DATABASE_ERROR = errors.New("database error")
)

func GetHttpError(err error) common.Response[any] {
	switch err {
	case DATABASE_ERROR:
		return common.Response[any]{
			Status:  500,
			Message: DATABASE_ERROR.Error(),
		}
	}

	return common.Response[any]{
		Status:  500,
		Message: "unexpected error",
	}
}
