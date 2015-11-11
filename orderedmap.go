package omap

type any interface{}

type Map interface {
	Set(key, value any)
	Get(key any) (any, bool)
	Delete(key any)
}

type omap struct {
	keys []any
	body map[any]any
}

var _ Map = &omap{}

func New() Map {
	return &omap{
		keys: make([]any, 0),
		body: make(map[any]any),
	}
}

func (m *omap) Set(key, value any) {
	if _, exist := m.body[key]; !exist {
		m.keys = append(m.keys, key)
	}
	m.body[key] = value
}

func (m *omap) Get(key any) (any, bool) {
	v, ok := m.body[key]
	return v, ok
}

func (m *omap) Delete(key any) {
	i := index(m.keys, key)
	if i == -1 {
		return
	}
	delete(m.body, key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
}

func index(keys []any, key any) int {
	for i, k := range keys {
		if k == key {
			return i
		}
	}
	return -1
}
