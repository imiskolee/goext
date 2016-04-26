//Package goextgoextmap is a php key-value map impl.
//fetures:
//safely dynamic basic type support
package goext

import (
	"errors"
	"fmt"
	"strconv"
)

//Map is define a php like map
type Map map[string]interface{}

func NewMap(m map[string]interface{}) Map {
	return Map(m)
}

//GetInt 将值解析为一个int64
func (m Map) GetInt(key string) (int64, error) {

	item, ok := m[key]

	if !ok {
		return 0, errors.New("key not exists.")
	}

	return strconv.ParseInt(fmt.Sprint(item), 10, 64)
}

//GetUint 将值解析为一个uint64
func (m Map) GetUint(key string) (uint64, error) {

	item, ok := m[key]

	if !ok {
		return 0, errors.New("key not exists.")
	}

	return strconv.ParseUint(fmt.Sprint(item), 10, 64)

}

//GetFloat 将值解析成一个float64
func (m Map) GetFloat(key string) (float64, error) {

	item, ok := m[key]

	if !ok {
		return 0, errors.New("key not exists.")
	}

	return strconv.ParseFloat(fmt.Sprint(item), 64)

}

//GetString 将值解析成一个String
func (m Map) GetString(key string) string {

	item, ok := m[key]
	if !ok {
		return ""
	}
	return fmt.Sprint(item)
}

//GetBool 将值解析成一个bool
func (m Map) GetBool(key string) (bool, error) {

	item, ok := m[key]

	if !ok {
		return false, errors.New("key not exists.")
	}

	return strconv.ParseBool(fmt.Sprint(item))

}
