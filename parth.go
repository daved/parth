// Package parth provides path parsing for segment slicing and unmarshaling. In
// other words, parth provides simple and flexible access to (URL) path
// parameters.
//
// Valid values are:
//   - builtin: *string, *bool, *int, *int64, *int32, *int16, *int8, *uint,
//     *uint64, *uint32, *uint16, *uint8, *float64, *float32
//   - stdlib: [*time.Duration], [encoding.TextUnmarshaler], [flag.Value]
//
// When handling any size of int, uint, or float, the first valid value within
// the specified segment will be used. Three important terms used in this
// package are "segment", "sequent", and "span". A segment is any single path
// section. A sequent is a segment that follows a "key" path section. Segments
// are able to be unmarshaled into variables. A span is any multiple path
// sections, and is handled as a string.
package parth

import (
	"encoding"
	"errors"
	"flag"
	"time"
)

// Err{Name} values facilitate error identification.
var (
	ErrUnknownType = errors.New("unknown type provided")

	ErrFirstSegNotFound = errors.New("first segment not found by index")
	ErrLastSegNotFound  = errors.New("last segment not found by index")
	ErrSegOrderReversed = errors.New("first segment must precede last segment")
	ErrKeySegNotFound   = errors.New("segment not found by key")

	ErrDataUnparsable = errors.New("data cannot be parsed")
)

// Segment locates the path segment indicated by index i. If the index is
// negative, the negative count begins with the last segment.
func Segment(v any, path string, i int) error {
	var err error

	switch v := v.(type) {
	case *bool:
		*v, err = segmentToBool(path, i)

	case *float32:
		var f float64
		f, err = segmentToFloatN(path, i, 32)
		*v = float32(f)

	case *float64:
		*v, err = segmentToFloatN(path, i, 64)

	case *int:
		var n int64
		n, err = segmentToIntN(path, i, 0)
		*v = int(n)

	case *int16:
		var n int64
		n, err = segmentToIntN(path, i, 16)
		*v = int16(n)

	case *int32:
		var n int64
		n, err = segmentToIntN(path, i, 32)
		*v = int32(n)

	case *int64:
		*v, err = segmentToIntN(path, i, 64)

	case *int8:
		var n int64
		n, err = segmentToIntN(path, i, 8)
		*v = int8(n)

	case *string:
		*v, err = segmentToString(path, i)

	case *uint:
		var n uint64
		n, err = segmentToUintN(path, i, 0)
		*v = uint(n)

	case *uint16:
		var n uint64
		n, err = segmentToUintN(path, i, 16)
		*v = uint16(n)

	case *uint32:
		var n uint64
		n, err = segmentToUintN(path, i, 32)
		*v = uint32(n)

	case *uint64:
		*v, err = segmentToUintN(path, i, 64)

	case *uint8:
		var n uint64
		n, err = segmentToUintN(path, i, 8)
		*v = uint8(n)

	case *time.Duration:
		var s string
		s, err = segmentToString(path, i)
		if err == nil {
			d, err := time.ParseDuration(s)
			if err == nil {
				*v = d
			}
		}

	case encoding.TextUnmarshaler:
		var s string
		s, err = segmentToString(path, i)
		if err == nil {
			err = v.UnmarshalText([]byte(s))
		}

	case flag.Value:
		var s string
		s, err = segmentToString(path, i)
		if err == nil {
			err = v.Set(s)
		}

	default:
		err = ErrUnknownType
	}

	return err
}

// Sequent is similar to [Segment], except that it locates the segment that is
// subsequent to the "key" segment.
func Sequent(v any, path, key string) error {
	return SubSeg(v, path, key, 0)
}

// Span returns the path segments between indexes i and j, including the segment
// indicated by index i. If an index is negative, the negative count begins with
// the last segment. Providing a 0 for index j is a special case which acts as
// an alias for the end of the path. If the first segment does not begin with a
// slash and it is part of the requested span, no slash will be added. Index i
// must not precede index j.
func Span(path string, i, j int) (string, error) {
	var f, l int
	var ok bool

	if i < 0 {
		f, ok = segStartIndexFromEnd(path, i)
	} else {
		f, ok = segStartIndexFromStart(path, i)
	}
	if !ok {
		return "", ErrFirstSegNotFound
	}

	if j > 0 {
		l, ok = segEndIndexFromStart(path, j)
	} else {
		l, ok = segEndIndexFromEnd(path, j)
	}
	if !ok {
		return "", ErrLastSegNotFound
	}

	if f == l {
		return "", nil
	}

	if f > l {
		return "", ErrSegOrderReversed
	}

	return path[f:l], nil
}

// SubSeg is similar to both [Sequent] and [Segment]. It first locates the
// "key", then uses index i to locate a segment. For example, to access the
// segment immediately after the "key", an index of 0 should be provided (which
// is how [Sequent] is implemented). Technically, a negative index is valid,
// but it is nonsensical in this function.
func SubSeg(v any, path, key string, i int) error {
	var err error

	switch v := v.(type) {
	case *bool:
		*v, err = subSegToBool(path, key, i)

	case *float32:
		var f float64
		f, err = subSegToFloatN(path, key, i, 32)
		*v = float32(f)

	case *float64:
		*v, err = subSegToFloatN(path, key, i, 64)

	case *int:
		var n int64
		n, err = subSegToIntN(path, key, i, 0)
		*v = int(n)

	case *int16:
		var n int64
		n, err = subSegToIntN(path, key, i, 16)
		*v = int16(n)

	case *int32:
		var n int64
		n, err = subSegToIntN(path, key, i, 32)
		*v = int32(n)

	case *int64:
		*v, err = subSegToIntN(path, key, i, 64)

	case *int8:
		var n int64
		n, err = subSegToIntN(path, key, i, 8)
		*v = int8(n)

	case *string:
		*v, err = subSegToString(path, key, i)

	case *uint:
		var n uint64
		n, err = subSegToUintN(path, key, i, 0)
		*v = uint(n)

	case *uint16:
		var n uint64
		n, err = subSegToUintN(path, key, i, 16)
		*v = uint16(n)

	case *uint32:
		var n uint64
		n, err = subSegToUintN(path, key, i, 32)
		*v = uint32(n)

	case *uint64:
		*v, err = subSegToUintN(path, key, i, 64)

	case *uint8:
		var n uint64
		n, err = subSegToUintN(path, key, i, 8)
		*v = uint8(n)

	case encoding.TextUnmarshaler:
		var s string
		s, err = subSegToString(path, key, i)
		if err == nil {
			err = v.UnmarshalText([]byte(s))
		}

	default:
		err = ErrUnknownType
	}

	return err
}

// SubSpan is similar to [Span], but only handles the portion of the path
// subsequent to the "key".
func SubSpan(path, key string, i, j int) (string, error) {
	si, ok := segIndexByKey(path, key)
	if !ok {
		return "", ErrKeySegNotFound
	}

	if i >= 0 {
		i++
	}
	if j > 0 {
		j++
	}

	s, err := Span(path[si:], i, j)
	if err != nil {
		return "", err
	}

	return s, nil
}

// Parth manages path and error data for processing a single path multiple
// times while handling errors only once. Only the first encountered error is
// stored since all subsequent calls to Parth methods will have no effect.
type Parth struct {
	path string
	err  error
}

// New constructs a pointer to an instance of [Parth] around the provided path.
func New(path string) *Parth {
	return &Parth{path: path}
}

// NewBySpan constructs a pointer to an instance of [Parth] after preprocessing
// the provided path with [Span].
func NewBySpan(path string, i, j int) *Parth {
	s, err := Span(path, i, j)
	return &Parth{s, err}
}

// NewBySubSpan constructs a pointer to an instance of [Parth] after
// preprocessing the provided path with [SubSpan].
func NewBySubSpan(path, key string, i, j int) *Parth {
	s, err := SubSpan(path, key, i, j)
	return &Parth{s, err}
}

// Err returns the first error encountered by the [*Parth] instance.
func (p *Parth) Err() error {
	return p.err
}

// Segment operates the same as the package-level function [Segment].
func (p *Parth) Segment(v any, i int) {
	if p.err != nil {
		return
	}

	p.err = Segment(v, p.path, i)
}

// Sequent operates the same as the package-level function [Sequent].
func (p *Parth) Sequent(v any, key string) {
	p.SubSeg(v, key, 0)
}

// Span operates the same as the package-level function [Span].
func (p *Parth) Span(i, j int) string {
	if p.err != nil {
		return ""
	}

	s, err := Span(p.path, i, j)
	p.err = err

	return s
}

// SubSeg operates the same as the package-level function [SubSeg].
func (p *Parth) SubSeg(v any, key string, i int) {
	if p.err != nil {
		return
	}

	p.err = SubSeg(v, p.path, key, i)
}

// SubSpan operates the same as the package-level function [SubSpan].
func (p *Parth) SubSpan(key string, i, j int) string {
	if p.err != nil {
		return ""
	}

	s, err := SubSpan(p.path, key, i, j)
	p.err = err

	return s
}
