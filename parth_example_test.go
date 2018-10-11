package parth_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/parth"
)

var (
	r, rErr = http.NewRequest("GET", "/zero/1/2/nn3.3nn/key/5.5", nil)
)

func init() {
	if rErr != nil {
		panic(rErr)
	}
}

func Example() {
	fmt.Println(r.URL.Path)

	var s string
	if err := parth.Segment(r.URL.Path, 0, &s); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(s)

	var f float32
	if err := parth.SubSeg(r.URL.Path, "key", &f); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(f)

	// Output:
	// /zero/1/2/nn3.3nn/key/5.5
	// zero
	// 5.5
}

func Example_parthType() {
	fmt.Println(r.URL.Path)

	var s string
	var f float32

	p := parth.New(r.URL.Path)
	p.Segment(0, &s)
	p.SubSeg("key", &f)
	if err := p.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(s)
	fmt.Println(f)

	// Output:
	// /zero/1/2/nn3.3nn/key/5.5
	// zero
	// 5.5
}
