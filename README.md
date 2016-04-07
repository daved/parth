# parth

    go get github.com/codemodus/parth

Package parth provides functions for accessing path segments.

When returning an int of any size, the first whole number within the specified 
segment will be returned.  When returning a float of any size, the first 
decimal number within the specified segment will be returned.

## Usage

```go
func SegmentToBool(path string, i int) (bool, error)
func SegmentToFloat32(path string, i int) (float32, error)
func SegmentToFloat64(path string, i int) (float64, error)
func SegmentToInt(path string, i int) (int, error)
func SegmentToInt16(path string, i int) (int16, error)
func SegmentToInt32(path string, i int) (int32, error)
func SegmentToInt64(path string, i int) (int64, error)
func SegmentToInt8(path string, i int) (int8, error)
func SegmentToString(path string, i int) (string, error)
func SpanToString(path string, firstSeg, lastSeg int) (string, error)
func SubSegToBool(path, key string) (bool, error)
func SubSegToFloat32(path, key string) (float32, error)
func SubSegToFloat64(path, key string) (float64, error)
func SubSegToInt(path, key string) (int, error)
func SubSegToInt16(path, key string) (int16, error)
func SubSegToInt32(path, key string) (int32, error)
func SubSegToInt64(path, key string) (int64, error)
func SubSegToInt8(path, key string) (int8, error)
func SubSegToString(path, key string) (string, error)
func SubSpanToString(path, key string, lastSeg int) (string, error)
type Parth
    func New(path string) Parth
    func NewFromSpan(path string, firstSeg, lastSeg int) Parth
    func NewFromSubSpan(path, key string, lastSeg int) Parth
    func (p *Parth) Err() error
    func (p *Parth) SegmentToBool(i int) bool
    func (p *Parth) SegmentToFloat32(i int) float32
    func (p *Parth) SegmentToFloat64(i int) float64
    func (p *Parth) SegmentToInt(i int) int
    func (p *Parth) SegmentToInt16(i int) int16
    func (p *Parth) SegmentToInt32(i int) int32
    func (p *Parth) SegmentToInt64(i int) int64
    func (p *Parth) SegmentToInt8(i int) int8
    func (p *Parth) SegmentToString(i int) string
    func (p *Parth) SpanToString(firstSeg, lastSeg int) string
    func (p *Parth) SubSegToBool(key string) bool
    func (p *Parth) SubSegToFloat32(key string) float32
    func (p *Parth) SubSegToFloat64(key string) float64
    func (p *Parth) SubSegToInt(key string) int
    func (p *Parth) SubSegToInt16(key string) int16
    func (p *Parth) SubSegToInt32(key string) int32
    func (p *Parth) SubSegToInt64(key string) int64
    func (p *Parth) SubSegToInt8(key string) int8
    func (p *Parth) SubSegToString(key string) string
    func (p *Parth) SubSpanToString(key string, lastSeg int) string
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

	out1, err := parth.SegmentToBool(path, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out1) // Prints: true

	out2, err := parth.SegmentToInt(path, -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out2) // Prints: 2

	out3, err := parth.SpanToString(path, 0, -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out3) // Prints: "/zero/1"

	out4, err := parth.SubSegToInt(path, "zero")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out4) // Prints: 1
}
```

## More Info

### Caution (restated): First Whole, First Decimal

When returning an int of any size, the first whole number within the specified 
segment will be returned.  When returning a float of any size, the first 
decimal number within the specified segment will be returned.

Please review the test cases for working examples.

### SpanToString (Indexed similarly to slices/arrays)

SpanToString receives two int values representing path segments, and returns 
the content between those segments, including the first segment, as a string 
and a nil error. The segments can be of negative values, but the first segment 
must come before the last segment. Providing a 0 int for the second int is a 
special case which indicates the end of the path.

### SubSpanToString (Indexed similarly to slices/arrays - Starting from key)

SubSpanToString receives a key which is used to search for the first matching 
path segment and an int value representing a second segment by it's distance 
from the matched segment, and returns the content between those segments as a 
string and a nil error. The int representing a segment can be of negative 
values. Providing a 0 int is a special case which indicates the end of the 
path.

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/parth)

## Benchmarks

    Go 1.6

    benchmark                     iter      time/iter   bytes alloc        allocs
    ---------                     ----      ---------   -----------        ------
    BenchmarkStandardInt-8     5000000   371.00 ns/op       64 B/op   2 allocs/op
    BenchmarkParthInt-8       20000000    78.50 ns/op        0 B/op   0 allocs/op
    BenchmarkParthIntNeg-8    30000000    52.70 ns/op        0 B/op   0 allocs/op
    BenchmarkParthSubSeg-8    20000000    93.60 ns/op        0 B/op   0 allocs/op
    BenchmarkStandardSpan-8    3000000   527.00 ns/op       88 B/op   4 allocs/op
    BenchmarkParthSpan-8      50000000    29.70 ns/op        0 B/op   0 allocs/op
    BenchmarkParthSubSpan-8   20000000    82.60 ns/op        0 B/op   0 allocs/op
