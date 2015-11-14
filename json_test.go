package omap

import (
	"reflect"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	m := New()
	m.Set("foo", "1")
	m.Set("bar", "2")

	dump, err := m.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(dump, []byte(`{"foo":"1","bar":"2"}`)) {
		t.Errorf("Got %s", dump)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	s := `
	{
		"foo": "1",
		"bar": "2",
		"baz": [1,2,3],
		"foobar": {
			"hoge": "fuga"
		}
	}`
	m := New()
	err := m.UnmarshalJSON([]byte(s))
	if err != nil {
		t.Error(err)
	}

	keys := []string{"foo", "bar", "baz", "foobar"}
	if !reflect.DeepEqual(keys, m.(*omap).keys) {
		t.Errorf("Expected %v, but got %v", keys, m.(*omap).keys)
	}

	gotBody := m.(*omap).body
	if gotBody["foo"] != any("1") {
		t.Errorf("foo should be 1")
	}
	if gotBody["bar"] != any("2") {
		t.Errorf("bar should be 2")
	}
}
