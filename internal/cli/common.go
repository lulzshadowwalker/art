package cli

import (
	"errors"
	"fmt"
	"reflect"
	"text/template"

	"github.com/lulzshadowwalker/art/internal/util"
)

var defaultFuncMap = template.FuncMap{
	"required":    required,
	"upperFirst":  util.UpperFirst,
	"lowerFirst":  util.LowerFirst,
	"firstLetter": util.FirstLetter,
}

var ErrFieldRequired = errors.New("field required")

func required(field string, value interface{}) (interface{}, error) {
	if value == nil {
		return nil, fmt.Errorf("%s %w", field, ErrFieldRequired)
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Bool:
		return value, nil
	default:
		if v.IsZero() {
			return nil, fmt.Errorf("%s %w", field, ErrFieldRequired)
		}

		return value, nil
	}
}
