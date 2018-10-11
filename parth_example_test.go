package parth_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/parth"
)

var (
	r, rErr = http.NewRequest("GET", "/zero/1/2/key/nn4.4nn/5.5", nil)
)

func init() {
	if rErr != nil {
		panic(rErr)
	}
}

func Example_segment() {
	var s string
	if err := parth.Segment(r.URL.Path, 4, &s); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Printf("%v (%T)\n", s, s)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// nn4.4nn (string)
}

func Example_sequent() {
	var f float32
	if err := parth.Sequent(r.URL.Path, "key", &f); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Printf("%v (%T)\n", f, f)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// 4.4 (float32)
}

func Example_span() {
	s, err := parth.Span(r.URL.Path, 2, 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Println(s)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// /2/key
}

func Example_subSeg() {
	var f float64
	if err := parth.SubSeg(r.URL.Path, "key", 1, &f); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Printf("%v (%T)\n", f, f)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// 5.5 (float64)
}

func Example_subSpan() {
	s0, err := parth.SubSpan(r.URL.Path, "zero", 2, 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	s1, err := parth.SubSpan(r.URL.Path, "1", 1, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Println(s0)
	fmt.Println(s1)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// /key/nn4.4nn
	// /key/nn4.4nn
}

func ExampleParth() {
	var s string
	var f float32

	p := parth.New(r.URL.Path)
	p.Segment(0, &s)
	p.SubSeg("key", 1, &f)
	if err := p.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(r.URL.Path)
	fmt.Println(s)
	fmt.Println(f)

	// Output:
	// /zero/1/2/key/nn4.4nn/5.5
	// zero
	// 5.5
}
