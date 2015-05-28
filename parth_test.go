package parth_test

import (
	"fmt"
	"testing"

	"github.com/codemodus/parth"
)

const (
	errFmtGotWant = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr  = "Did not receive expected err for segment value %v"
	printFmt      = "Segment Index = %v , Type = %T, Value = %v\n"
)

func Example() {
	urlPath := "/zero/1/2/nn3.3nn/4.4"
	parthObject := parth.New(urlPath)

	if seg, err := parthObject.String(0); err == nil {
		fmt.Printf(printFmt, 0, seg, seg)
	}

	if seg, err := parthObject.Bool(1); err == nil {
		fmt.Printf(printFmt, 1, seg, seg)
	}

	if seg, err := parthObject.Int(2); err == nil {
		fmt.Printf(printFmt, 2, seg, seg)
	}

	if seg, err := parthObject.Float32(3); err == nil {
		fmt.Printf(printFmt, 3, seg, seg)
	}

	if seg, err := parthObject.Int(4); err == nil {
		fmt.Printf(printFmt, 4, seg, seg)
	}

	// Output:
	// Segment Index = 0 , Type = string, Value = zero
	// Segment Index = 1 , Type = bool, Value = true
	// Segment Index = 2 , Type = int, Value = 2
	// Segment Index = 3 , Type = float32, Value = 3.3
	// Segment Index = 4 , Type = int, Value = 4

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
		{1, "/test5//", "", true},
		{3, "/test6", "", true},
		{0, "//test7", "", true},
		{0, "/", "", true},
	}

	for _, v := range tests {
		po := parth.New(v.p)
		seg, err := po.String(v.i)
		if err != nil && !v.e {
			t.Fatal(err)
		}
		if err == nil && v.e {
			t.Errorf(errFmtExpErr, v.p)
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
		p string
		r int
		e bool
	}{
		{"/1", 1, false},
		{"/2.2", 2, false},
		{"/.3", 0, false},
		{"/error", 0, true},
	}

	for _, v := range tests {
		po := parth.New(v.p)
		seg, err := po.Int(0)
		if err != nil && !v.e {
			t.Fatal(err)
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
		po := parth.New(v.p)
		seg, err := po.Int8(0)
		if err != nil && !v.e {
			t.Fatal(err)
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
		po := parth.New(v.p)
		seg, err := po.Int16(0)
		if err != nil && !v.e {
			t.Fatal(err)
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
		po := parth.New(v.p)
		seg, err := po.Int32(0)
		if err != nil && !v.e {
			t.Fatal(err)
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
		po := parth.New(v.p)
		seg, err := po.Int64(0)
		if err != nil && !v.e {
			t.Fatal(err)
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
		p string
		b bool
		e bool
	}{
		{"/1", true, false},
		{"/t", true, false},
		{"/T", true, false},
		{"/true", true, false},
		{"/TRUE", true, false},
		{"/True", true, false},
		{"/error", false, true},
		{"/0", false, false},
		{"/f", false, false},
		{"/F", false, false},
		{"/false", false, false},
		{"/FALSE", false, false},
		{"/False", false, false},
	}

	var path string
	for _, v := range tests {
		path += v.p
	}
	po := parth.New(path)

	for k, v := range tests {
		seg, err := po.Bool(k)
		if err != nil && !v.e {
			t.Fatal(err)
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
		p   string
		f32 float32
		f64 float64
		e   bool
	}{
		{"/0.1", 0.1, 0.1, false},
		{"/0.2a", 0.2, 0.2, false},
		{"/aaaa1.3", 1.3, 1.3, false},
		{"/4", 4.0, 4.0, false},
		{"/5aaaa", 5.0, 5.0, false},
		{"/aaa6aa", 6.0, 6.0, false},
		{"/.7.aaaa", 0.7, 0.7, false},
		{"/.8aa", 0.8, 0.8, false},
		{"/error", 0.0, 0.0, true},
		{"/.", 0.0, 0.0, true},
	}

	var path string
	for _, v := range tests {
		path += v.p
	}
	po := parth.New(path)

	for k, v := range tests {
		seg, err := po.Float32(k)
		if err != nil && !v.e {
			t.Fatal(err)
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

	for k, v := range tests {
		seg, err := po.Float64(k)
		if err != nil && !v.e {
			t.Fatal(err)
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
