package util

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	str2 := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	str2 = matchAllCap.ReplaceAllString(str2, "${1}_${2}")
	//str2 = strings.ReplaceAll(str2, "-", "_")
	return strings.ToLower(str2)
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

func Encrypt(s string) string {
	return s
}
