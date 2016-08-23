package parth_test

import (
	"fmt"

	"github.com/codemodus/parth"
)

func Example() {
	testPath := "/zero/1/2/nn3.3nn/key/5.5"
	printFmt := "Segment Index = %v, Type = %T, Value = %v\n"

	if s, err := parth.SegmentToString(testPath, 0); err == nil {
		fmt.Printf(printFmt, 0, s, s)
	}

	if b, err := parth.SegmentToBool(testPath, 1); err == nil {
		fmt.Printf(printFmt, 1, b, b)
	}

	if i, err := parth.SegmentToInt(testPath, -4); err == nil {
		fmt.Printf(printFmt, -4, i, i)
	}

	if f, err := parth.SegmentToFloat32(testPath, 3); err == nil {
		fmt.Printf(printFmt, 3, f, f)
	}

	if s, err := parth.SpanToString(testPath, 0, -3); err == nil {
		fmt.Printf("First Segment = %d, Last Segment = %d, Value = %q\n", 0, -3, s)
	}

	if i, err := parth.SubSegToInt(testPath, "key"); err == nil {
		fmt.Printf("Segment Key = %q, Type = %T, Value = %v\n", "key", i, i)
	}

	if s, err := parth.SubSpanToString(testPath, "zero", 2); err == nil {
		fmt.Printf("Segment Key = %q, Last Segment = %d, Value = %q\n", "zero", 2, s)
	}

	// Output:
	// Segment Index = 0, Type = string, Value = zero
	// Segment Index = 1, Type = bool, Value = true
	// Segment Index = -4, Type = int, Value = 2
	// Segment Index = 3, Type = float32, Value = 3.3
	// First Segment = 0, Last Segment = -3, Value = "/zero/1/2"
	// Segment Key = "key", Type = int, Value = 5
	// Segment Key = "zero", Last Segment = 2, Value = "/1/2"
}
