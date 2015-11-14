package omap

import (
	"bytes"
	"encoding/json"
)

type any interface{}

type Map interface {
	Set(key string, value any)
	Get(key string) (any, bool)
	Delete(key string)

	Each(func(any))

	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

type omap struct {
	keys []string
	body map[string]any
}

var _ Map = &omap{}
var _ json.Marshaler = &omap{}
var _ json.Unmarshaler = &omap{}

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

func (m *omap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{'{'})

	for _, key := range m.keys {
		// Marshal key
		b, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		buf.WriteByte(':')

		// Marshal value
		b, err = json.Marshal(m.body[key])
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		buf.WriteByte(',')
	}

	b := buf.Bytes()
	b[len(b)-1] = '}'
	return b, nil
}

func (m *omap) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &m.body)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(bytes.NewReader(b))

	dec.Token() // {
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}

		m.keys = append(m.keys, t.(string))

		t, err = dec.Token()
		if err != nil {
			return err
		}
		switch tok := t.(type) {
		case json.Delim:
			err := skipDelim(dec, tok)
			if err != nil {
				return err
			}
		}
	}
	dec.Token() // }

	return nil
}

func skipDelim(dec *json.Decoder, start json.Delim) error {
	var end json.Delim
	if start == '[' {
		end = ']'
	} else { // '{'
		end = '}'
	}

	cnt := 1
	for {
		t, err := dec.Token()
		if err != nil {
			return err
		}
		delim, isDelim := t.(json.Delim)
		if !isDelim {
			continue
		}
		switch delim {
		case start:
			cnt++
		case end:
			cnt--
			if cnt == 0 {
				return nil
			}
		}
	}
	return nil
}
