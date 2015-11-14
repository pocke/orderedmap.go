package omap

import (
	"bytes"
	"encoding/json"
)

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

var _ json.Marshaler = &omap{}
var _ json.Unmarshaler = &omap{}
