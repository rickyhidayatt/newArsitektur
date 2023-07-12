package response

import (
	"context"
	"time"
)

const (
	KEY_REQUEST_TIME = "Hbm:T37.[ewrN;Ns"
)

type ResolveStructure struct {
	Data interface{} `json:"data"`
	// Metadata map[string]interface{} `json:"metadata"`
}

type RejectStructure struct {
	Code    int    `json:"-"`
	Error   error  `json:"error"`
	Message string `json:"message"`
}

func ResponseWithRequestTime(
	ctx context.Context,
	data interface{},
	// metadata map[string]interface{},
) interface{} {
	meta := make(map[string]interface{})

	// for k, v := range metadata {
	// 	meta[k] = v
	// }

	if value := ctx.Value(KEY_REQUEST_TIME); value != nil {
		meta["request_took"] = time.Since(value.(time.Time)).Seconds()
		meta["request_measure"] = "time_measure.second"
	}

	return ResolveStructure{
		Data: data,
		// Metadata: meta,
	}
}

func ResponseError(
	statusCode int,
	error error,
	message string,
) RejectStructure {
	return RejectStructure{
		Code:    statusCode,
		Error:   error,
		Message: message,
	}
}
