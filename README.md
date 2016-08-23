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
    r, err := http.NewRequest("GET", "/zero/1/2/nn3.3nn/key/5.5", nil)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    printFmt := "Type = %T, Value = %v\n"

    if s, err := parth.SegmentToString(r.URL.Path, 0); err == nil {
        fmt.Printf(printFmt, s, s) // Outputs: Type = string, Value = zero
    }

    if b, err := parth.SegmentToBool(r.URL.Path, 1); err == nil {
        fmt.Printf(printFmt, b, b) // Outputs: Type = bool, Value = true
    }

    if i, err := parth.SegmentToInt(r.URL.Path, -4); err == nil {
        fmt.Printf(printFmt, i, i) // Outputs: Type = int, Value = 2
    }

    if f, err := parth.SegmentToFloat32(r.URL.Path, 3); err == nil {
        fmt.Printf(printFmt, f, f) // Outputs: Type = float32, Value = 3.3
    }

    if s, err := parth.SpanToString(r.URL.Path, 0, -3); err == nil {
        fmt.Printf(printFmt, s, s) // Outputs: Type = string, Value = /zero/1/2
    }

    if i, err := parth.SubSegToInt(r.URL.Path, "key"); err == nil {
        fmt.Printf(printFmt, i, i) // Outputs: Type = int, Value = 5
    }

    if s, err := parth.SubSpanToString(r.URL.Path, "zero", 2); err == nil {
        fmt.Printf(printFmt, s, s) // Outputs: Type = string, Value = /1/2
    }
}
```

## More Info

### Path parameters via global, alternate HandlerFunc, or Context? Why?

The most obvious use case for parth is when working with http.Request data 
within an http.Handler. parth is fast enough that it can be used 20+ times when 
compared to similar router-parameter/Context usage. Why pass data that is 
already being passed? The request type holds URL data and parth loves handling 
it! Additionally, parth takes care of parsing segments into the types actually 
needed. It's not only fast, it does more, and requires less code.  

### Caution (restated): First Whole, First Decimal

When returning an int of any size, the first whole number within the specified 
segment will be returned.  When returning a float of any size, the first 
decimal number within the specified segment will be returned.

Please review the test cases for working examples.

### SpanToString

SpanToString receives two int values representing path segments, and returns 
the content between those segments, including the first segment, as a string 
and a nil error. If any error is encountered, a zero value string and error are 
returned. The segments can be of negative values, but the first segment must 
come before the last segment. Providing a 0 int for the second int is a special 
case which indicates the end of the path.

### SubSpanToString

SubSpanToString receives a key which is used to search for the first matching 
path segment and an int value representing a second segment by it's distance 
from the matched segment, then returns the content between those segments as a 
string and a nil error. If any error is encountered, a zero value string and 
error are returned. The int representing a segment can be of negative values. 
Providing a 0 int is a special case which indicates the end of the path.

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


    Go 1.7

    benchmark                                 iter       time/iter   bytes alloc        allocs
    ---------                                 ----       ---------   -----------        ------
    BenchmarkVsCtxParthString2x-8         20000000     83.80 ns/op        0 B/op   0 allocs/op
    BenchmarkVsCtxParthString3x-8         10000000    125.00 ns/op        0 B/op   0 allocs/op
    BenchmarkVsCtxContextGetSetGet-8       1000000   1629.00 ns/op      336 B/op   5 allocs/op
    BenchmarkVsCtxContextGetSetGetGet-8    1000000   2044.00 ns/op      352 B/op   6 allocs/op
