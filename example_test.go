package parth_test

import (
	"fmt"
	"net/http"

	"github.com/daved/parth"
)

var req, _ = http.NewRequest("GET", "/zero/1/2/key/nn4.4nn/5.5", nil)

func Example() {
	var segFour string
	if err := parth.Segment(&segFour, req.URL.Path, 4); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v (%[1]T)\n", segFour)

	// Output:
	// nn4.4nn (string)
}

func ExampleSegment() {
	var segFour string
	if err := parth.Segment(&segFour, req.URL.Path, 4); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v (%[1]T)\n", segFour)

	// Output:
	// nn4.4nn (string)
}

func ExampleSequent() {
	var afterKey float32
	if err := parth.Sequent(&afterKey, req.URL.Path, "key"); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v (%[1]T)\n", afterKey)

	// Output:
	// 4.4 (float32)
}

func ExampleSpan() {
	span, err := parth.Span(req.URL.Path, 2, 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(span)

	// Output:
	// /2/key
}

func ExampleSubSeg() {
	var twoAfterKey float64
	if err := parth.SubSeg(&twoAfterKey, req.URL.Path, "key", 1); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v (%[1]T)\n", twoAfterKey)

	// Output:
	// 5.5 (float64)
}

func ExampleSubSpan() {
	subSpanZero, err := parth.SubSpan(req.URL.Path, "zero", 2, 4)
	if err != nil {
		fmt.Println(err)
	}

	subSpanOne, err := parth.SubSpan(req.URL.Path, "1", 1, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(subSpanZero)
	fmt.Println(subSpanOne)

	// Output:
	// /key/nn4.4nn
	// /key/nn4.4nn
}

func ExampleParth() {
	var segZero string
	var twoAfterKey float32

	p := parth.New(req.URL.Path)
	p.Segment(&segZero, 0)
	p.SubSeg(&twoAfterKey, "key", 1)
	if err := p.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v (%[1]T)\n", segZero)
	fmt.Printf("%[1]v (%[1]T)\n", twoAfterKey)

	// Output:
	// zero (string)
	// 5.5 (float32)
}

type MyType []byte

// UnmarshalText implements encoding.TextUnmarshaler. Let's pretend something
// interesting is actually happening here.
func (m *MyType) UnmarshalText(text []byte) error {
	*m = text
	return nil
}

func Example_encodingTextUnmarshaler() {
	var m MyType
	if err := parth.Segment(&m, req.URL.Path, 4); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%[1]v == %[1]q (%[1]T)\n", m)

	// Output:
	// [110 110 52 46 52 110 110] == "nn4.4nn" (parth_test.MyType)
}
