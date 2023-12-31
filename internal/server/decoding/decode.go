package decoding

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi/v5"
	gokitHttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
)

func Decode(model interface{}) gokitHttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		if model == nil {
			return nil, nil
		}

		if reflect.TypeOf(model).Kind() != reflect.Ptr {
			return nil, errors.New("decode model must be pointer")
		}

		var err error

		if r.ContentLength > 0 {
			contentType := r.Header["Content-Type"]
			for _, ct := range contentType {
				if ct == "application/json" {
					err = parseJSON(ctx, model, r)
					if nil != err {
						return nil, err
					}
				}
			}
		}

		err = getURLParam(ctx, model, r)
		if nil != err {
			return nil, err
		}

		err = getURLQueryString(ctx, model, r)
		if nil != err {
			return nil, err
		}

		err = getHeader(ctx, model, r)
		if nil != err {
			return nil, err
		}

		return model, err
	}
}

func parseJSON(ctx context.Context, model interface{}, r *http.Request) (err error) {
	err = json.NewDecoder(r.Body).Decode(model)
	return err
}

func getHeader(ctx context.Context, model interface{}, r *http.Request) error {
	var err error

	typeOf := reflect.TypeOf(model)
	elem := typeOf.Elem()
	for i := 0; i < elem.NumField(); i++ {
		tag := elem.Field(i).Tag.Get("header")
		if tag == "" {
			continue
		}

		value := r.Header.Get(tag)
		err = assignValue(model, value, i)
		if nil != err {
			return err
		}
	}

	return nil
}

func getURLParam(ctx context.Context, model interface{}, r *http.Request) error {
	var err error

	typeOf := reflect.TypeOf(model)
	elem := typeOf.Elem()

	for i := 0; i < elem.NumField(); i++ {
		tag := elem.Field(i).Tag.Get("url_param")
		if tag == "" {
			continue
		}

		value := chi.URLParam(r, tag)
		err = assignValue(model, value, i)
		if nil != err {
			return err
		}
	}

	return nil
}

func getURLQueryString(ctx context.Context, model interface{}, r *http.Request) error {
	var err error

	typeOf := reflect.TypeOf(model)
	elem := typeOf.Elem()

	for i := 0; i < elem.NumField(); i++ {
		tag := elem.Field(i).Tag.Get("qs")
		if tag == "" {
			continue
		}

		value := r.URL.Query().Get(tag)
		if len(value) <= 0 {
			continue
		}
		err = assignValue(model, value, i)
		if nil != err {
			return err
		}
	}

	return nil
}

func assignValue(model interface{}, value string, fieldIndex int) error {
	elem := reflect.ValueOf(model).Elem()

	switch elem.Field(fieldIndex).Type().String() {
	case "int", "int32", "int64":
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).SetInt(v)
	case "*int", "*int32", "*int64":
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).Set(reflect.ValueOf(&v))
	case "string":
		elem.Field(fieldIndex).SetString(value)
	case "*string":
		elem.Field(fieldIndex).Set(reflect.ValueOf(&value))
	case "bool":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).SetBool(v)
	case "*bool":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).Set(reflect.ValueOf(&v))
	case "uuid.UUID":
		v, err := uuid.Parse(value)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).Set(reflect.ValueOf(v))
	case "*uuid.UUID":
		v, err := uuid.Parse(value)
		if err != nil {
			return err
		}
		elem.Field(fieldIndex).Set(reflect.ValueOf(&v))
	case "map[string]interface {}":
		var v map[string]interface{}

		err := json.Unmarshal([]byte(value), &v)
		if err != nil {
			return err
		}

		elem.Field(fieldIndex).Set(reflect.ValueOf(v))
	}
	return nil
}
