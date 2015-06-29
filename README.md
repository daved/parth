# parth

    go get "github.com/codemodus/parth"

Package parth provides functions for accessing path segments.

When returning an int of any size, the first whole number within the specified 
segment will be returned.  When returning a float of any size, the first 
decimal number within the specified segment will be returned.

## Usage

```go
func SegmentToString(path string, i int) (string, error)
func SegmentToBool(path string, i int) (bool, error)
func SegmentToFloat32(path string, i int) (float32, error)
func SegmentToFloat64(path string, i int) (float64, error)
func SegmentToInt(path string, i int) (int, error)
func SegmentToInt16(path string, i int) (int16, error)
func SegmentToInt32(path string, i int) (int32, error)
func SegmentToInt64(path string, i int) (int64, error)
func SegmentToInt8(path string, i int) (int8, error)
func SpanToString(path string, firstSeg, lastSeg int) (string, error)
```

### Setup

```go
import (
    "fmt"

    "github.com/codemodus/parth"
)

func main() {
    path := "/zero/1/2"

    out0, err := parth.SegmentToString(path, 0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out0) // Prints: "zero"

    out1, err := parth.SegmentToInt(path, 1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out1) // Prints: 1

    out2, err := parth.SegmentToInt(path, -1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out2) // Prints: 2

    out3, err := parth.SpanToString(path, 0, -2)
    if err != nil {
		fmt.Println(err)
    }
    fmt.Println(out3) // Prints: "/zero/1"

}
```

## More Info

### Caution (restated): First Whole, First Decimal

When returning an int of any size, the first whole number within the specified 
segment will be returned.  When returning a float of any size, the first 
decimal number within the specified segment will be returned.

Please review the test cases for working examples.

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/parth)

## Benchmarks

These results compare standard library functions to parth functions.

    benchmark                  iter      time/iter   bytes alloc        allocs
    ---------                  ----      ---------   -----------        ------
    BenchmarkStandardInt    5000000   394.00 ns/op       64 B/op   3 allocs/op
    BenchmarkParthInt      20000000    70.00 ns/op        0 B/op   0 allocs/op
    BenchmarkParthIntNeg   30000000    50.00 ns/op        0 B/op   0 allocs/op
    BenchmarkStandardSpan   3000000   557.00 ns/op       88 B/op   4 allocs/op
    BenchmarkParthSpan     30000000    36.30 ns/op        0 B/op   0 allocs/op
