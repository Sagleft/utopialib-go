package utopiago

import (
	"errors"
	"reflect"
)

// GetString - get string field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetString(field string) (string, error) {
	errHandler := func(err error) (string, error) {
		return "", err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(string)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a string"))
	}

	return val, nil
}

// GetBool - get bool field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetBool(field string) (bool, error) {
	errHandler := func(err error) (bool, error) {
		return false, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(bool)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a bool"))
	}

	return val, nil
}

// GetInt - get int64 field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetInt(field string) (int64, error) {
	errHandler := func(err error) (int64, error) {
		return 0, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(int64)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a int64"))
	}

	return val, nil
}

// GetFloat - get int64 field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetFloat(field string) (float64, error) {
	errHandler := func(err error) (float64, error) {
		return 0, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(float64)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a float64"))
	}

	return val, nil
}
