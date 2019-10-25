package config

import "fmt"

type array struct {
	storage map[string][]interface{}
}

func (a *array) String() string {
	return fmt.Sprintf("%v", a.storage)
}

func (a *array) Add(name string, value ...interface{}) {
	_, ok := a.storage[name]
	if !ok {
		a.storage[name] = make([]interface{}, 0)
	}
	a.storage[name] = append(a.storage[name], value...)
}

func (a *array) Get(name string) ([]interface{}, bool) {
	arr, ok := a.storage[name]
	return arr, ok
}

func (a *array) GetIndex(name string, index int) (interface{}, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index], true
	}
	return nil, false
}

func (a *array) GetIndexBool(name string, index int) (bool, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(bool), true
	}
	return false, false
}

func (a *array) GetIndexFloat64(name string, index int) (float64, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(float64), true
	}
	return 0, false
}

func (a *array) GetIndexInt(name string, index int) (int, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(int), true
	}
	return 0, false
}

func (a *array) GetIndexInt32(name string, index int) (int32, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(int32), true
	}
	return 0, false
}

func (a *array) GetIndexInt64(name string, index int) (int64, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(int64), true
	}
	return 0, false
}

func (a *array) GetIndexString(name string, index int) (string, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(string), true
	}
	return "", false
}

func (a *array) GetIndexUint(name string, index int) (uint, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(uint), true
	}
	return 0, false
}

func (a *array) GetIndexUint32(name string, index int) (uint32, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(uint32), true
	}
	return 0, false
}

func (a *array) GetIndexUint64(name string, index int) (uint64, bool) {
	if arr, ok := a.storage[name]; ok && index < len(arr) {
		return arr[index].(uint64), true
	}
	return 0, false
}
