package parth

import (
	"fmt"
	"testing"
)

func TestBhvrSegment(t *testing.T) {
	path := "/junk/4/key/true/other/3.3/"
	key := ""

	t.Run("bool", applyToBoolTFunc(path, key, pti(3), true))
	t.Run("float32", applyToFloat32TFunc(path, key, pti(5), 3.3))
	t.Run("float64", applyToFloat64TFunc(path, key, pti(5), 3.3))
	t.Run("int", applyToIntTFunc(path, key, pti(1), 4))
	t.Run("int16", applyToInt16TFunc(path, key, pti(1), 4))
	t.Run("int32", applyToInt32TFunc(path, key, pti(1), 4))
	t.Run("int64", applyToInt64TFunc(path, key, pti(1), 4))
	t.Run("int8", applyToInt8TFunc(path, key, pti(1), 4))
	t.Run("string", applyToStringTFunc(path, key, pti(0), "junk"))
	t.Run("uint", applyToUintTFunc(path, key, pti(1), 4))
	t.Run("uint16", applyToUint16TFunc(path, key, pti(1), 4))
	t.Run("uint32", applyToUint32TFunc(path, key, pti(1), 4))
	t.Run("uint64", applyToUint64TFunc(path, key, pti(1), 4))
	t.Run("uint8", applyToUint8TFunc(path, key, pti(1), 4))
}

func TestBhvrSequent(t *testing.T) {
	path := "/junk/4/key/true/other/3.3/"
	var i *int

	t.Run("bool", applyToBoolTFunc(path, "key", i, true))
	t.Run("float32", applyToFloat32TFunc(path, "other", i, 3.3))
	t.Run("float64", applyToFloat64TFunc(path, "other", i, 3.3))
	t.Run("int", applyToIntTFunc(path, "junk", i, 4))
	t.Run("int16", applyToInt16TFunc(path, "junk", i, 4))
	t.Run("int32", applyToInt32TFunc(path, "junk", i, 4))
	t.Run("int64", applyToInt64TFunc(path, "junk", i, 4))
	t.Run("int8", applyToInt8TFunc(path, "junk", i, 4))
	t.Run("string", applyToStringTFunc(path, "key", i, "true"))
	t.Run("uint", applyToUintTFunc(path, "junk", i, 4))
	t.Run("uint16", applyToUint16TFunc(path, "junk", i, 4))
	t.Run("uint32", applyToUint32TFunc(path, "junk", i, 4))
	t.Run("uint64", applyToUint64TFunc(path, "junk", i, 4))
	t.Run("uint8", applyToUint8TFunc(path, "junk", i, 4))
}

func TestBhvrSpan(t *testing.T) {
}

func TestBhvrSubSeg(t *testing.T) {
	path := "/junk/4/key/true/other/3.3/"

	t.Run("bool", applyToBoolTFunc(path, "junk", pti(2), true))
	t.Run("float32", applyToFloat32TFunc(path, "true", pti(1), 3.3))
	t.Run("float64", applyToFloat64TFunc(path, "true", pti(1), 3.3))
	t.Run("int", applyToIntTFunc(path, "junk", pti(0), 4))
	t.Run("int16", applyToInt16TFunc(path, "junk", pti(0), 4))
	t.Run("int32", applyToInt32TFunc(path, "junk", pti(0), 4))
	t.Run("int64", applyToInt64TFunc(path, "junk", pti(0), 4))
	t.Run("int8", applyToInt8TFunc(path, "junk", pti(0), 4))
	t.Run("string", applyToStringTFunc(path, "junk", pti(3), "other"))
	t.Run("uint", applyToUintTFunc(path, "junk", pti(0), 4))
	t.Run("uint16", applyToUint16TFunc(path, "junk", pti(0), 4))
	t.Run("uint32", applyToUint32TFunc(path, "junk", pti(0), 4))
	t.Run("uint64", applyToUint64TFunc(path, "junk", pti(0), 4))
	t.Run("uint8", applyToUint8TFunc(path, "junk", pti(0), 4))
}

func TestBhvrSubSpan(t *testing.T) {
}

func TestBhvrParth(t *testing.T) {
}

func segSeqSubSeg(path, key string, i *int, v interface{}) error {
	if path != "" && key != "" && i != nil {
		return SubSeg(path, key, *i, v)
	}

	if path != "" && key != "" {
		return Sequent(path, key, v)
	}

	if path != "" && i != nil {
		return Segment(path, *i, v)
	}

	return fmt.Errorf("see segSeqSubSeg for missing requirements")
}

func applyToBoolTFunc(path, key string, i *int, want bool) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got bool
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToFloat32TFunc(path, key string, i *int, want float32) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got float32
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToFloat64TFunc(path, key string, i *int, want float64) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got float64
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToIntTFunc(path, key string, i *int, want int) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got int
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToInt16TFunc(path, key string, i *int, want int16) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got int16
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToInt32TFunc(path, key string, i *int, want int32) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got int32
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToInt64TFunc(path, key string, i *int, want int64) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got int64
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToInt8TFunc(path, key string, i *int, want int8) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got int8
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToStringTFunc(path, key string, i *int, want string) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got string
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToUintTFunc(path, key string, i *int, want uint) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got uint
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToUint16TFunc(path, key string, i *int, want uint16) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got uint16
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToUint32TFunc(path, key string, i *int, want uint32) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got uint32
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToUint64TFunc(path, key string, i *int, want uint64) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got uint64
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
}

func applyToUint8TFunc(path, key string, i *int, want uint8) func(*testing.T) {
	return func(t *testing.T) {
		subj := subject(path, key)
		if i != nil {
			subj = subject(path, key, *i)
		}

		var got uint8
		err := segSeqSubSeg(path, key, i, &got)
		if unx(t, subj, err) {
			return
		}

		if got != want {
			t.Errorf(gwFmt, got, want)
		}
	}
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
			t.Errorf(gwxFmt, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
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
			t.Errorf(gwxFmt, tt.s, got, tt.want)
		}
	}
}

func pti(i int) *int {
	return &i
}

var (
	gwFmt  = "got %v, want %v"
	gwxFmt = "subj '%v': got %v, want %v"
)

type checkFunc func(*testing.T, interface{}, error) bool

func unx(t *testing.T, subj interface{}, err error) bool {
	b := err != nil
	if b {
		t.Errorf(gwxFmt, subj, err, nil)
	}
	return b
}

func exp(t *testing.T, subj interface{}, err error) bool {
	b := err == nil
	if b {
		t.Errorf(gwxFmt, subj, nil, "{error}")
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
