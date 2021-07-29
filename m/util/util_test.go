package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	// Based on https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6#gistcomment-3515624
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"snake_case", "snake_case"},
		{"Pascal_Snake", "pascal_snake"},
		{"SCREAMING_SNAKE", "screaming_snake"},
		{"A", "a"},
		{"AA", "aa"},
		{"AAA", "aaa"},
		{"AAAA", "aaaa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"BatteryLifeValue", "battery_life_value"},
		{"Id0Value", "id0_value"},
		{"ID0Value", "id0_value"},
		{"DBError", "db_error"},
		{"UserID", "user_id"},
		{"UserId", "user_id"},
	}

	for _, test := range tests {
		result := SnakeCase(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestStructFields(t *testing.T) {
	input1 := struct {
		A int
		B uint64
		C byte
		D float64
		E string
		F complex128
		G bool
	}{}
	expected1 := map[string]reflect.Type{
		"A": reflect.TypeOf(input1.A),
		"B": reflect.TypeOf(input1.B),
		"C": reflect.TypeOf(input1.C),
		"D": reflect.TypeOf(input1.D),
		"E": reflect.TypeOf(input1.E),
		"F": reflect.TypeOf(input1.F),
		"G": reflect.TypeOf(input1.G),
	}
	output1, err := StructFields(input1)

	if assert.NoError(t, err) {
		assert.True(t, reflect.DeepEqual(output1, expected1))
	}
}
