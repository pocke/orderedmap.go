package omap

import "testing"

func TestMap(t *testing.T) {
	m := New()

	m.Set("foo", "bar")

	if got, exist := m.Get("foo"); got != "bar" {
		t.Errorf("Expected bar but got %s", got)
	} else if !exist {
		t.Error("not exist")
	}

	m.Delete("honya")
	m.Delete("foo")

	if _, exist := m.Get("foo"); exist {
		t.Errorf("should not be exist, but exist")
	}
}
