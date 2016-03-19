package parth_test

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/codemodus/parth"
)

func Example() {
	testPath := "/zero/1/2/nn3.3nn/4.4"
	printFmt := "Segment Index = %v , Type = %T, Value = %v\n"

	if s0, err := parth.SegmentToString(testPath, 0); err == nil {
		fmt.Printf(printFmt, 0, s0, s0)
	}

	if b1, err := parth.SegmentToBool(testPath, 1); err == nil {
		fmt.Printf(printFmt, 1, b1, b1)
	}

	if i2, err := parth.SegmentToInt(testPath, 2); err == nil {
		fmt.Printf(printFmt, 2, i2, i2)
	}

	if f3, err := parth.SegmentToFloat32(testPath, 3); err == nil {
		fmt.Printf(printFmt, 3, f3, f3)
	}

	if i4, err := parth.SegmentToInt(testPath, 4); err == nil {
		fmt.Printf(printFmt, 4, i4, i4)
	}

	if iN1, err := parth.SegmentToInt(testPath, -1); err == nil {
		fmt.Printf(printFmt, -1, iN1, iN1)
	}

	if s0N3, err := parth.SpanToString(testPath, 0, -2); err == nil {
		fmt.Printf("First Segment = 0, Last Segment = -2, Value = %v", s0N3)
	}

	// Output:
	// Segment Index = 0 , Type = string, Value = zero
	// Segment Index = 1 , Type = bool, Value = true
	// Segment Index = 2 , Type = int, Value = 2
	// Segment Index = 3 , Type = float32, Value = 3.3
	// Segment Index = 4 , Type = int, Value = 4
	// Segment Index = -1 , Type = int, Value = 4
	// First Segment = 0, Last Segment = -2, Value = /zero/1/2
}

var (
	errFmtGotWant  = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr   = "Did not receive expected err for segment value %v"
	errFmtUnexpErr = "Received unexpected err for segment type %T: %v"
)

func TestFunctString(t *testing.T) {
	var tests = []struct {
		ind   int
		path  string
		s     string
		isErr bool
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
		s, err := parth.SegmentToString(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctInts(t *testing.T) {
	var tests = []struct {
		ind   int
		path  string
		i     int
		isErr bool
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
		i, err := parth.SegmentToInt(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, i, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt8(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, seg, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := int8(v.i)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt16(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, i, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := int16(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt32(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, i, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := int32(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt64(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, i, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := int64(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctBool(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		b     bool
		isErr bool
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
		b, err := parth.SegmentToBool(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, b, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := v.b
		got := b
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctFloats(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		f32   float32
		f64   float64
		isErr bool
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
		f32, err := parth.SegmentToFloat32(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, f32, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := v.f32
		got := f32
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		f64, err := parth.SegmentToFloat64(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, f64, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := v.f64
		got := f64
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSpan(t *testing.T) {
	var tests = []struct {
		firstInd int
		lastInd  int
		path     string
		s        string
		isErr    bool
	}{
		{0, 0, "/test1", "/test1", false},
		{0, 1, "/test1", "/test1", false},
		{0, 1, "/test1/test-2", "/test1", false},
		{1, 2, "/test1/test-2/test_3/", "/test-2", false},
		{0, 0, "test4/t4", "test4/t4", false},
		{0, 1, "t444/t4", "t444", false},
		{0, 1, "//test5", "/", false},
		{0, 1, "/test6//", "/test6", false},
		{0, 2, "/t6//", "/t6/", false},
		{0, 3, "/66//", "/66//", false},
		{1, 2, "/test7", "", true},
		{0, -1, "/test8", "", false},
		{1, 1, "/t/9", "", false},
		{0, 0, "/", "/", false},
		{1, 1, "/", "", true},
		{-1, -1, "/", "", false},
		{0, -1, "/", "", false},
		{-1, 0, "/", "/", false},
		{-1, 0, "/test1", "/test1", false},
		{0, -1, "/test1/test-2", "/test1", false},
		{-3, -1, "/test1/test-2/test_3", "/test1/test-2", false},
		{-1, -1, "/test11/test-12", "", false},
		{-1, -3, "/test11/test-12", "", true},
		{-2, -1, "test4/t4/", "/t4", false},
		{-1, -3, "/test5/test-6/test_7", "", true},
		{-3, 0, "/test7", "", true},
	}

	for _, v := range tests {
		s, err := parth.SpanToString(v.path, v.firstInd, v.lastInd)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

var (
	bmri int
	bmrs string
)

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
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = standardSegment(p, 1)
	}

	bmri = r
}

func BenchmarkParthInt(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SegmentToInt(p, 1)
	}

	bmri = r
}

func BenchmarkParthIntNeg(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SegmentToInt(p, -1)
	}

	bmri = r
}

func BenchmarkStandardSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		cs := strings.Split(p, "/")
		if p[0] == '/' {
			cs[1] = "/" + cs[1]
		}
		r = path.Join(cs[0:3]...)
	}

	bmrs = r
}

func BenchmarkParthSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SpanToString(p, 0, 1)
	}

	bmrs = r
}
