package encoding

import (
	"context"
	"encoding/json"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func Encode() kitHttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if nil == response {
			w.WriteHeader(http.StatusNoContent)
			_ = json.NewEncoder(w).Encode(nil)
			return nil
		}

		_ = json.NewEncoder(w).Encode(response)

		return nil
	}
}

// TransformObject used to transform source object to result object based on json tag
func TransformObject(source interface{}, result interface{}) error {
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sourceBytes, &result)
	if err != nil {
		return err
	}

	return nil
}
