package config

import "fmt"

type hashmap struct {
	storage map[string]map[string]interface{}
}

func (h *hashmap) String() string {
	return fmt.Sprintf("%v", h.storage)
}

func (h *hashmap) Add(name string, key string, value interface{}) {
	_, ok := h.storage[name]
	if !ok {
		h.storage[name] = make(map[string]interface{})
	}
	h.storage[name][key] = value
}

func (h *hashmap) Get(name string) (map[string]interface{}, bool) {
	hmap, ok := h.storage[name]
	return hmap, ok
}

func (h *hashmap) GetValue(name string, key string) (interface{}, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value, true
		}
	}
	return nil, false
}

func (h *hashmap) GetValueBool(name string, key string) (bool, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(bool), true
		}
	}
	return false, false
}

func (h *hashmap) GetValueFloat64(name string, key string) (float64, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(float64), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueInt(name string, key string) (int, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(int), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueInt32(name string, key string) (int32, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(int32), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueInt64(name string, key string) (int64, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(int64), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueString(name string, key string) (string, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(string), true
		}
	}
	return "", false
}

func (h *hashmap) GetValueUint(name string, key string) (uint, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(uint), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueUint32(name string, key string) (uint32, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(uint32), true
		}
	}
	return 0, false
}

func (h *hashmap) GetValueUint64(name string, key string) (uint64, bool) {
	if hmap, ok := h.storage[name]; ok {
		if value, ok := hmap[key]; ok {
			return value.(uint64), true
		}
	}
	return 0, false
}
