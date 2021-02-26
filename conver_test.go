package convert

import "testing"

//Good candicate for generics!?
type Test struct {
	For       interface{}
	Expect    interface{}
	WithError error
}

func TestString2SliceByte(t *testing.T) {
	tests := []struct {
		have string
		want []byte
	}{
		{"test", []byte{'t', 'e', 's', 't'}},
		{"", nil},
		{"👍🙂", []byte{240, 159, 145, 141, 240, 159, 153, 130}},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v", test.have, test.want)
		got := String2SliceByte(test.have)

		if string(got) != string(test.want) {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}

func TestSliceByte2String(t *testing.T) {
	tests := []struct {
		have []byte
		want string
	}{
		{[]byte{'t', 'e', 's', 't'}, "test"},
		{[]byte{}, ""},
		{nil, ""},
		{[]byte{240, 159, 145, 141, 240, 159, 153, 130}, "👍🙂"},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v", test.have, test.want)
		got := SliceByte2String(test.have)

		if got != test.want {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}

func TestInt2String(t *testing.T) {
	tests := []struct {
		have int
		want string
	}{
		{1, "1"},
		{123, "123"},
		{0, "0"},
		{-1, "-1"},
		{-123, "-123"},
		{1231231231231231231, "1231231231231231231"},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v", test.have, test.want)
		got := Int2String(test.have)

		if got != test.want {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}

func TestFloat642String(t *testing.T) {
	tests := []struct {
		have       float64
		withFormat string
		want       string
	}{
		{123.123, "%.1f", "123.1"},
		{123.123, "%.2f", "123.12"},
		{123.123, "%.3f", "123.123"},
		{123.987, "%.3f", "123.987"},
		{123.987, "%.2f", "123.99"},
		{123.987, "%.1f", "124.0"},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v", test.have, test.want)
		got := Float642String(test.have, test.withFormat)

		if got != test.want {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}

func TestInt2Float64(t *testing.T) {
	have := int(123)
	want := float64(123)
	got := Int2Float64(have)
	if want != got {
		t.Errorf("For %#v, expected %#v, but got %#v", have, want, got)
	}
}

func TestString2Float64(t *testing.T) {
	tests := []struct {
		have string
		want float64
		err  bool
	}{
		{"123", 123.0, false},
		{"123.123", 123.123, false},
		{"123.987", 123.987, false},
		{"987.987", 987.987, false},
		{"abc", 0, true},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v & err %v", test.have, test.want, test.err)
		got, err := String2Float64(test.have)

		if !test.err && err != nil {
			t.Errorf("For %#v, got unexpected err %#v", test.have, err)
		}
		if test.err && err == nil {
			t.Errorf("For %#v, expected err but got none", test.have)
		}
		if got != test.want {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}

func TestFloat642Int(t *testing.T) {
	have := float64(123.987)
	want := int(123)
	got := Float642Int(have)
	if want != got {
		t.Errorf("For %#v, expected %#v, but got %#v", have, want, got)
	}
}

func TestString2Int(t *testing.T) {
	tests := []struct {
		have string
		want int
		err  bool
	}{
		{"123", 123, false},
		{"123.123", 0, true},
		{"1231231231231231231", 1231231231231231231, false},
		{"abc", 0, true},
	}

	for _, test := range tests {
		t.Logf("have %#v, want %#v & err %v", test.have, test.want, test.err)
		got, err := String2Int(test.have)

		if !test.err && err != nil {
			t.Errorf("For %#v, got unexpected err %#v", test.have, err)
		}
		if test.err && err == nil {
			t.Errorf("For %#v, expected err but got none", test.have)
		}
		if got != test.want {
			t.Errorf("For %#v, expected %#v, but got %#v", test.have, test.want, got)
		}
	}
}
