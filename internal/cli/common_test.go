package cli

import (
	"errors"
	"reflect"
	"testing"
)

func TestRequired(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    interface{}
		expected interface{}
		err      error
	}{
		{"Nil value", "name", nil, nil, ErrFieldRequired},
		{"Empty string", "name", "", nil, ErrFieldRequired},
		{"Non-empty string", "name", "John", "John", nil},
		{"Empty slice", "list", []int{}, []int{}, nil},
		{"Non-empty slice", "list", []int{1, 2, 3}, []int{1, 2, 3}, nil},
		{"Empty map", "dict", map[string]int{}, map[string]int{}, nil},
		{"Non-empty map", "dict", map[string]int{"key": 1}, map[string]int{"key": 1}, nil},
		{"Nil pointer", "ptr", (*int)(nil), nil, ErrFieldRequired},
		{"Non-nil pointer", "ptr", new(int), new(int), nil},
		{"Empty struct", "struct", struct{}{}, nil, ErrFieldRequired},
		{"Non-empty struct", "struct", struct{ Name string }{Name: "John"}, struct{ Name string }{Name: "John"}, nil},
		{"Number zero", "zero", 0, 0, nil},
		{"Integer", "int", 42, 42, nil},
		{"Float64", "float64", 3.14, 3.14, nil},
		{"Float32", "float32", float32(3.14), float32(3.14), nil},
		{"True boolean", "bool", true, true, nil},
		{"False boolean", "bool", false, false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := required(tt.field, tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if (err == nil && tt.err != nil) || (err != nil && tt.err == nil) || (err != nil && tt.err != nil && !errors.Is(err, tt.err)) {
				t.Errorf("expected error %v, got error %v", tt.err, err)
			}
		})
	}
}
