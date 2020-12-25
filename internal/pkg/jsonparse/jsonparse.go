package jsonparse

import "github.com/tidwall/gjson"

func Get(json, path string) interface{} {
	return gjson.Get(json, path).Value()
}
