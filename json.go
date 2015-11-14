package omap

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var endOfSliceError = fmt.Errorf("End of slice")

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
	dec := json.NewDecoder(bytes.NewReader(b))

	dec.Token() // {
	m.unmarshalJSON(dec)
	dec.Token() // }

	return nil
}

func (m *omap) unmarshalJSON(dec *json.Decoder) error {
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}

		key, isKey := t.(string)
		if !isKey {
			return fmt.Errorf("%t %s is not string(expected key)", t, t)
		}

		val, err := getVal(dec)
		if err != nil {
			return err
		}
		m.Set(key, val)
	}
	return nil
}

func decToSlice(dec *json.Decoder) ([]any, error) {
	res := make([]any, 0)

	for {
		v, err := getVal(dec)
		if err == endOfSliceError {
			return res, nil
		}
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
}

func getVal(dec *json.Decoder) (any, error) {
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}

	switch tok := t.(type) {
	case json.Delim:
		switch tok {
		case '[':
			return decToSlice(dec)
		case '{':
			om := New()
			err := om.(*omap).unmarshalJSON(dec)
			if err != nil {
				return nil, err
			}
			_, err = dec.Token() // }
			return om, err
		case ']':
			return nil, endOfSliceError
		}
	default:
		return tok, nil
	}

	panic("unreachable code")
}

var _ json.Marshaler = &omap{}
var _ json.Unmarshaler = &omap{}
