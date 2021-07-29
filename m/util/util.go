package util

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(camelCaseStr string) string {
	// Based on https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6#gistcomment-3515624
	str := matchFirstCap.ReplaceAllString(camelCaseStr, "${1}_${2}")
	str = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(str)
}

func StructFields(s interface{}) (map[string]reflect.Type, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("not struct")
	}

	m := make(map[string]reflect.Type)
	num := t.NumField()
	for i := 0; i < num; i++ {
		f := t.Field(i)
		m[f.Name] = f.Type
	}
	return m, nil
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}
