package parth_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/codemodus/parth"
)

const (
	errFmtGotWant  = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr   = "Did not receive expected err for segment value %v"
	errFmtUnexpErr = "Received unexpected err for segment type %T: %v"
	printFmt       = "Segment Index = %v , Type = %T, Value = %v\n"
)

func Example() {
	path := "/zero/1/2/nn3.3nn/4.4"

	if seg, err := parth.SegmentToString(path, 0); err == nil {
		fmt.Printf(printFmt, 0, seg, seg)
	}

	if seg, err := parth.SegmentToBool(path, 1); err == nil {
		fmt.Printf(printFmt, 1, seg, seg)
	}

	if seg, err := parth.SegmentToInt(path, 2); err == nil {
		fmt.Printf(printFmt, 2, seg, seg)
	}

	if seg, err := parth.SegmentToFloat32(path, 3); err == nil {
		fmt.Printf(printFmt, 3, seg, seg)
	}

	if seg, err := parth.SegmentToInt(path, 4); err == nil {
		fmt.Printf(printFmt, 4, seg, seg)
	}

	if seg, err := parth.SegmentToInt(path, -1); err == nil {
		fmt.Printf(printFmt, -1, seg, seg)
	}

	if path, err := parth.SpanToString(path, 0, -3); err == nil {
		fmt.Printf("First Segment = 0, Last Segment = -3, Value = %v", path)
	}

	// Output:
	// Segment Index = 0 , Type = string, Value = zero
	// Segment Index = 1 , Type = bool, Value = true
	// Segment Index = 2 , Type = int, Value = 2
	// Segment Index = 3 , Type = float32, Value = 3.3
	// Segment Index = 4 , Type = int, Value = 4
	// Segment Index = -1 , Type = int, Value = 4
	// First Segment = 0, Last Segment = -3, Value = /zero/1/2
}

func TestString(t *testing.T) {
	var tests = []struct {
		i int
		p string
		r string
		e bool
	}{
		{0, "/test1", "test1", false},
		{1, "/test1/test-2", "test-2", false},
		{2, "/test1/test-2/test_3/", "test_3", false},
		{0, "test4/t4", "test4", false},
		{1, "//test5", "test5", false},
		{1, "/test6//", "", true},
		{3, "/test7", "", true},
		{0, "//test8", "", true},
		{0, "/", "", true},
		{-1, "/test1", "test1", false},
		{-1, "/test1/test-2", "test-2", false},
		{-2, "/test1/test-2", "test1", false},
		{-3, "/test1/test-2/test_3", "test1", false},
		{-1, "test4/t4/", "t4", false},
		{-1, "//test5", "test5", false},
		{-1, "/test6//", "", true},
		{-3, "/test7", "", true},
		{-2, "//test8", "", true},
		{-1, "/", "", true},
	}

	for _, v := range tests {
		seg, err := parth.SegmentToString(v.p, v.i)
		if err != nil && !v.e {
			t.Errorf(errFmtUnexpErr, seg, err)
			continue
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
			continue
		}

		want := v.r
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestInts(t *testing.T) {
	var tests = []struct {
		i int
		p string
		r int
		e bool
	}{
		{0, "/0.1", 0, false},
		{0, "/0.2a", 0, false},
		{0, "/aaaa1.3", 1, false},
		{0, "/4", 4, false},
		{0, "/5aaaa", 5, false},
		{0, "/aaa6aa", 6, false},
		{0, "/.7.aaaa", 0, false},
		{0, "/.8aa", 0, false},
		{0, "/-9", -9, false},
		{-1, "/-9", -9, false},
		{0, "/10-", 10, false},
		{0, "/3.14e+11", 3, false},
		{0, "/3.14e.+12", 3, false},
		{0, "/3.14e+.13", 3, false},
		{-1, "/3.14e+.13", 3, false},
		{1, "/8", 0, true},
		{0, "/.", 0, true},
		{0, "/error", 0, true},
		{0, "/12414143242534534346456456457457456346756868686524234", 0, true},
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := int(v.r)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt8(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := int8(v.r)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt16(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := int16(v.r)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt32(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := int32(v.r)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt64(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := int64(v.r)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		i int
		p string
		b bool
		e bool
	}{
		{0, "/1", true, false},
		{0, "/t", true, false},
		{0, "/T", true, false},
		{0, "/true", true, false},
		{0, "/TRUE", true, false},
		{0, "/True", true, false},
		{0, "/0", false, false},
		{0, "/f", false, false},
		{0, "/F", false, false},
		{-1, "/F", false, false},
		{0, "/false", false, false},
		{0, "/FALSE", false, false},
		{0, "/False", false, false},
		{1, "/True", false, true},
		{0, "/error", false, true},
	}

	for _, v := range tests {
		seg, err := parth.SegmentToBool(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := v.b
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFloats(t *testing.T) {
	tests := []struct {
		i   int
		p   string
		f32 float32
		f64 float64
		e   bool
	}{
		{0, "/0.1", 0.1, 0.1, false},
		{0, "/0.2a", 0.2, 0.2, false},
		{0, "/aaaa1.3", 1.3, 1.3, false},
		{0, "/4", 4.0, 4.0, false},
		{0, "/5aaaa", 5.0, 5.0, false},
		{0, "/aaa6aa", 6.0, 6.0, false},
		{0, "/.7.aaaa", 0.7, 0.7, false},
		{0, "/.8aa", 0.8, 0.8, false},
		{0, "/-9", -9.0, -9.0, false},
		{0, "/10-", 10.0, 10.0, false},
		{0, "/3.14e+11", 3.14e+11, 3.14e+11, false},
		{0, "/3.14e.+12", 3.14, 3.14, false},
		{0, "/3.14e+.13", 3.14, 3.14, false},
		{-1, "/3.14e+.13", 3.14, 3.14, false},
		{1, "/14", 0.0, 0.0, true},
		{0, "/error", 0.0, 0.0, true},
		{0, "/.", 0.0, 0.0, true},
		{0, "/3.14e+407", 0.0, 0.0, true},
	}

	for _, v := range tests {
		seg, err := parth.SegmentToFloat32(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := v.f32
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToFloat64(v.p, v.i)
		if err != nil && !v.e {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
		}

		want := v.f64
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestSpan(t *testing.T) {
	var tests = []struct {
		i int
		n int
		p string
		r string
		e bool
	}{
		{0, 0, "/test1", "/test1", false},
		{0, 1, "/test1/test-2", "/test1/test-2", false},
		{0, 1, "/test1/test-2/test_3/", "/test1/test-2", false},
		{0, 0, "test4/t4", "test4", false},
		{0, 1, "//test5", "//test5", false},
		{0, 1, "/test6//", "/test6/", false},
		{1, 2, "/test7", "", true},
		{0, -1, "/test8", "/test8", false},
		{0, 0, "/", "/", false},
		{-1, -1, "/", "/", false},
		{0, -1, "/", "/", false},
		{-1, 0, "/", "/", false},
		{-1, 0, "/test1", "/test1", false},
		{-1, 0, "/test1/test-2", "", true},
		{-2, 0, "/test1/test-2", "/test1", false},
		{-3, -1, "/test1/test-2/test_3", "/test1/test-2/test_3", false},
		{-1, -1, "/test11/test-12", "/test-12", false},
		{-1, -1, "test4/t4/", "/t4", false},
		{-1, -3, "/test5/test-6/test_7", "", true},
		{-3, 0, "/test7", "", true},
	}

	for _, v := range tests {
		spn, err := parth.SpanToString(v.p, v.i, v.n)
		if err != nil && !v.e {
			t.Errorf(errFmtUnexpErr, spn, err)
			continue
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
			continue
		}

		want := v.r
		got := spn
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func standardSegment(path string, i int) (int, error) {
	ss := strings.Split(strings.TrimLeft(path, "/"), "/")
	if len(ss) == 0 || i > len(ss) {
		err := fmt.Errorf("segment out of bounds")
		return 0, err
	}
	v, err := strconv.ParseInt(ss[i], 10, 0)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func BenchmarkStandardInt(b *testing.B) {
	path := "/zero/1"
	for n := 0; n < b.N; n++ {
		_, _ = standardSegment(path, 1)
	}
}

func BenchmarkParthInt(b *testing.B) {
	path := "/zero/1"
	for n := 0; n < b.N; n++ {
		_, _ = parth.SegmentToInt(path, 1)
	}
}

func BenchmarkParthIntNeg(b *testing.B) {
	path := "/zero/1"
	for n := 0; n < b.N; n++ {
		_, _ = parth.SegmentToInt(path, -1)
	}
}

func BenchmarkParthSpan(b *testing.B) {
	path := "/zero/1/2"
	for n := 0; n < b.N; n++ {
		_, _ = parth.SpanToString(path, 0, 1)
	}
}
