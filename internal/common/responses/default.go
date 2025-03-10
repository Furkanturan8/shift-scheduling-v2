package responses

import (
	"shift-scheduling-V2/internal/domain/enum"
)

var (
	DefaultSuccessResponse = General{
		Status:    200,
		Code:      enum.OK,
		ErrorCode: enum.COMMON_CODE,
		Message:   "success",
		Data:      nil,
	}

	DefaultErrorResponse = General{
		Status:    500,
		Code:      enum.OK,
		ErrorCode: "",
		Message:   "Internal server error",
		Data:      nil,
	}
)
