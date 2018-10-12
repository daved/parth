package parth

import (
	"fmt"
	"testing"
)

func TestBhvrSegment(t *testing.T) {
	path := "/junk/key/true/other"

	t.Run("bool", func(t *testing.T) {
		i := 2
		subj := subject(path, "", i)

		var got bool
		err := Segment(path, i, &got)
		unx(t, subj, err)

		want := true
		if got != want {
			t.Errorf(gwFmt, subj, got, want)
		}
	})
}

func TestBhvrSequent(t *testing.T) {
}

func TestBhvrSpan(t *testing.T) {
}

func TestBhvrSubSeg(t *testing.T) {
}

func TestBhvrSubSpan(t *testing.T) {
}

func TestBhvrParth(t *testing.T) {
}

func TestUnitFirstFloatFromString(t *testing.T) {
	tests := []struct {
		s    string
		want string
		ck   checkFunc
	}{
		{"/0.1", "0.1", unx},
		{"/0.2a", "0.2", unx},
		{"/aaaa1.3", "1.3", unx},
		{"/4", "4", unx},
		{"/5aaaa", "5", unx},
		{"/aaa6aa", "6", unx},
		{"/.7.aaaa", ".7", unx},
		{"/.8aa", ".8", unx},
		{"/-9", "-9", unx},
		{"/10-", "10", unx},
		{"/3.14e+11", "3.14e+11", unx},
		{"/3.14e.+12", "3.14", unx},
		{"/3.14e+.13", "3.14", unx},
		{"/3.14e+.13", "3.14", unx},
		{"/error", "", exp},
		{"/.", "", exp},
	}

	for _, tt := range tests {
		got, err := firstFloatFromString(tt.s)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, got, tt.want)
		}
	}
}

func TestUnitFirstIntFromString(t *testing.T) {
	var tests = []struct {
		s    string
		want string
		ck   checkFunc
	}{
		{"0.1", "0", unx},
		{"0.2a", "0", unx},
		{"aaaa1.3", "1", unx},
		{"4", "4", unx},
		{"5aaaa", "5", unx},
		{"aaa6aa", "6", unx},
		{".7.aaaa", "0", unx},
		{".8aa", "0", unx},
		{"-9", "-9", unx},
		{"10-", "10", unx},
		{"3.14e+11", "3", unx},
		{"3.14e.+12", "3", unx},
		{"3.14e+.13", "3", unx},
		{"18446744073709551615", "18446744073709551615", unx},
		{".", "", exp},
		{"error", "", exp},
	}

	for _, tt := range tests {
		got, err := firstIntFromString(tt.s)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitFirstUintFromString(t *testing.T) {
	var tests = []struct {
		s    string
		want string
		ck   checkFunc
	}{
		{"0.1", "0", unx},
		{"0.2a", "0", unx},
		{"aaaa1.3", "1", unx},
		{"4", "4", unx},
		{"5aaaa", "5", unx},
		{"aaa6aa", "6", unx},
		{".7.aaaa", "0", unx},
		{".8aa", "0", unx},
		{"-9", "9", unx},
		{"10-", "10", unx},
		{"3.14e+11", "3", unx},
		{"3.14e.+12", "3", unx},
		{"3.14e+.13", "3", unx},
		{"18446744073709551615", "18446744073709551615", unx},
		{".", "", exp},
		{"error", "", exp},
	}

	for _, tt := range tests {
		got, err := firstUintFromString(tt.s)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitSegEndIndexFromEnd(t *testing.T) {
	tests := []struct {
		i    int
		s    string
		want int
		ck   checkFunc
	}{
		{0, "/test1", 6, unx},
		{-1, "/t1", 0, unx},
		{-1, "/test1/test-2", 6, unx},
		{-2, "/test1/t-2/t_3", 6, unx},
		{-3, "test3/t3/", 0, unx},
		{-1, "test4/t4/", 8, unx},
		{-2, "/t5/f/fiv/55/5/fi/ve", 14, unx},
		{-1, "/", 0, unx},
		{-4, "/test/out", 0, exp},
		{4, "/test/out", 0, exp},
	}

	for _, tt := range tests {
		got, err := segEndIndexFromEnd(tt.s, tt.i)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitSegEndIndexFromStart(t *testing.T) {
	tests := []struct {
		i    int
		s    string
		want int
		ck   checkFunc
	}{
		{1, "/test1", 6, unx},
		{2, "/test1/test-2", 13, unx},
		{2, "/test1/test-2/test_3", 13, unx},
		{1, "test3/t3/", 5, unx},
		{2, "test4/t4/", 8, unx},
		{5, "/t5/f/fiv/55/5/fi/ve", 14, unx},
		{1, "/", 1, unx},
		{-4, "/test/out", 0, exp},
		{4, "/test/out", 0, exp},
	}

	for _, tt := range tests {
		got, err := segEndIndexFromStart(tt.s, tt.i)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitSegIndexByKey(t *testing.T) {
	tests := []struct {
		k    string
		s    string
		want int
		ck   checkFunc
	}{
		{"test", "/1/test/3", 2, unx},
		{"2", "/2/t/3", 0, unx},
		{"3", "/1/test/3", 7, unx},
		{"4", "/44/44/33", 0, exp},
		{"best", "12/best/3", 2, unx},
		{"6", "6/tt/66", 0, unx},
		{"7", "1/test/7", 6, unx},
		{"first", "first/2/three", 0, unx},
		{"bad", "/ba/d/", 0, exp},
		{"11", "/4/56/11/", 5, unx},
		{"", "/4/56/11/", 0, exp},
		{"t", "", 0, exp},
	}

	for _, tt := range tests {
		got, err := segIndexByKey(tt.s, tt.k)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitSegStartIndexFromEnd(t *testing.T) {
	tests := []struct {
		i    int
		s    string
		want int
		ck   checkFunc
	}{
		{-1, "/test1", 0, unx},
		{-1, "/test1/test-2", 6, unx},
		{-1, "/test1/test-2/test_3", 13, unx},
		{-3, "test3/t3/", 0, unx},
		{-1, "test4/t4/", 8, unx},
		{-2, "/t5/f/fiv/55/5/fi/ve", 14, unx},
		{-1, "/", 0, unx},
		{-2, "/", 0, exp},
		{-4, "/test/out", 0, exp},
		{1, "/test/out", 0, exp},
	}

	for _, tt := range tests {
		got, err := segStartIndexFromEnd(tt.s, tt.i)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

func TestUnitSegStartIndexFromStart(t *testing.T) {
	tests := []struct {
		i    int
		s    string
		want int
		ck   checkFunc
	}{
		{0, "/test1", 0, unx},
		{1, "/test1/test-2", 6, unx},
		{2, "/t1-2/fd", 0, exp},
		{2, "/test1/test-2/test_3", 13, unx},
		{0, "test3/t3/", 0, unx},
		{1, "test4/t4/", 5, unx},
		{6, "/t5/f/fiv/55/5/fi/ve", 17, unx},
		{0, "/", 0, unx},
		{1, "/", 0, exp},
		{2, "/", 0, exp},
		{4, "/test/out", 0, exp},
		{-1, "/test/out", 0, exp},
		{2, "/0/1//", 4, unx},
		{3, "/0/1//", 5, unx},
	}

	for _, tt := range tests {
		got, err := segStartIndexFromStart(tt.s, tt.i)
		if tt.ck(t, tt.s, err) {
			continue
		}

		if got != tt.want {
			t.Errorf(gwFmt, tt.s, got, tt.want)
		}
	}
}

var gwFmt = "subj '%v': got %v, want %v"

type checkFunc func(*testing.T, interface{}, error) bool

func unx(t *testing.T, subj interface{}, err error) bool {
	b := err != nil
	if b {
		t.Errorf("subj '%v': got %v, want nil", subj, err)
	}
	return b
}

func exp(t *testing.T, subj interface{}, err error) bool {
	b := err == nil
	if b {
		t.Errorf("subj '%v': got nil, want error", subj)
	}
	return b
}

func subject(path, key string, indexes ...int) string {
	s := "path " + path
	if key != "" {
		s += ", key " + key
	}

	if len(indexes) > 0 {
		s += ", indexes"
	}

	for _, i := range indexes {
		s += fmt.Sprintf(" %d", i)
	}

	return s
}
