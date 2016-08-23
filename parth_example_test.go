package parth_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/parth"
)

func Example() {
	r, err := http.NewRequest("GET", "/zero/1/2/nn3.3nn/key/5.5", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	printFmt := "Type = %T, Value = %v\n"

	if s, err := parth.SegmentToString(r.URL.Path, 0); err == nil {
		fmt.Printf(printFmt, s, s)
	}

	if b, err := parth.SegmentToBool(r.URL.Path, 1); err == nil {
		fmt.Printf(printFmt, b, b)
	}

	if i, err := parth.SegmentToInt(r.URL.Path, -4); err == nil {
		fmt.Printf(printFmt, i, i)
	}

	if f, err := parth.SegmentToFloat32(r.URL.Path, 3); err == nil {
		fmt.Printf(printFmt, f, f)
	}

	if s, err := parth.SpanToString(r.URL.Path, 0, -3); err == nil {
		fmt.Printf(printFmt, s, s)
	}

	if i, err := parth.SubSegToInt(r.URL.Path, "key"); err == nil {
		fmt.Printf(printFmt, i, i)
	}

	if s, err := parth.SubSpanToString(r.URL.Path, "zero", 2); err == nil {
		fmt.Printf(printFmt, s, s)
	}

	// Output:
	// Type = string, Value = zero
	// Type = bool, Value = true
	// Type = int, Value = 2
	// Type = float32, Value = 3.3
	// Type = string, Value = /zero/1/2
	// Type = int, Value = 5
	// Type = string, Value = /1/2
}

func Example_parthType() {
	r, err := http.NewRequest("GET", "/zero/1/2/nn3.3nn/key/5.5", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	printFmt := "Type = %T, Value = %v\n"

	p := parth.New(r.URL.Path)

	s := p.SegmentToString(0)
	f := p.SegmentToFloat32(3)
	ss := p.SubSpanToString("zero", 2)

	if p.Err() == nil {
		fmt.Printf(printFmt, s, s)
		fmt.Printf(printFmt, f, f)
		fmt.Printf(printFmt, ss, ss)
	}

	// Output:
	// Type = string, Value = zero
	// Type = float32, Value = 3.3
	// Type = string, Value = /1/2
}
