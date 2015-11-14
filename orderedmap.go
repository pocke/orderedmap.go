package omap

type any interface{}

type Map interface {
	Set(key string, value any)
	Get(key string) (any, bool)
	Delete(key string)

	Each(func(any))
}

type omap struct {
	keys []string
	body map[string]any
}

var _ Map = &omap{}

func New() Map {
	return &omap{
		keys: make([]string, 0),
		body: make(map[string]any),
	}
}

func (m *omap) Set(key string, value any) {
	if _, exist := m.body[key]; !exist {
		m.keys = append(m.keys, key)
	}
	m.body[key] = value
}

func (m *omap) Get(key string) (any, bool) {
	v, ok := m.body[key]
	return v, ok
}

func (m *omap) Delete(key string) {
	i := index(m.keys, key)
	if i == -1 {
		return
	}
	delete(m.body, key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
}

func index(keys []string, key string) int {
	for i, k := range keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (m *omap) Each(f func(any)) {
	for _, v := range m.keys {
		f(m.body[v])
	}
}

func MarshalJSON(b []byte) error {
	return nil
}
