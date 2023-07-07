package liberror

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := http.StatusInternalServerError
	message := "Something Went Wrong"

	if sc, ok := err.(*Error); ok {
		code = sc.StatusCode
		message = sc.Message
	}

	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   err.Error(),
		"code":    code,
		"message": message,
	})
}
