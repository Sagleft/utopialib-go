package utopia

import (
	"math"
	"reflect"
)

type uMap map[string]interface{}

func (u uMap) add(fieldTag string, value any) uMap {
	if !reflect.ValueOf(value).IsZero() {
		u[fieldTag] = value
	}
	return u
}

func (u uMap) set(fieldTag string, value interface{}) uMap {
	u[fieldTag] = value
	return u
}

func roundFloat(v float64, digits int) float64 {
	digitsPow := math.Pow10(digits)
	return math.Round(v*digitsPow) / digitsPow
}
