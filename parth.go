// Package parth provides functions for accessing path segments.
//
// When returning an int/uint of any size, the first whole number within the
// specified segment will be returned. When returning a float of any size,
// the first decimal number within the specified segment will be returned.
package parth

import (
	"errors"
)

// Err{Name} values are for error identification.
var (
	ErrUnknownType = errors.New("unknown type provided")

	ErrSegNotExist      = errors.New("segment not found by index")
	ErrFirstSegNotExist = errors.New("first segment not found by index")
	ErrLastSegNotExist  = errors.New("last segment not found by index")
	ErrKeyNotExist      = errors.New("segment not found by key")
	ErrSegOrderReversed = errors.New("first segment must precede last segment")
	ErrUnparsable       = errors.New("unable to parse segment")
)

// Segment receives an int representing a path segment, then returns
// both the specified segment as a string and a nil error. If any error is
// encountered, a zero value string and error are returned.
func Segment(path string, i int, v interface{}) error {
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

	default:
		err = ErrUnknownType
	}

	return err
}

// Sequent ...
func Sequent(path, key string, v interface{}) error {
	return SubSeg(path, key, 0, v)
}

// Span receives two int values representing path segments, and
// returns the content between those segments, including the first segment, as
// a string and a nil error. If any error is encountered, a zero value string
// and error are returned. The segments can be of negative values, but the
// first segment must come before the last segment. Providing a 0 int for the
// second int is a special case which indicates the end of the path.
func Span(path string, i, j int) (string, error) {
	var f, l int
	var ok bool

	if i < 0 {
		f, ok = segStartIndexFromEnd(path, i)
	} else {
		f, ok = segStartIndexFromStart(path, i)
	}
	if !ok {
		return "", ErrFirstSegNotExist
	}

	if j > 0 {
		l, ok = segEndIndexFromStart(path, j)
	} else {
		l, ok = segEndIndexFromEnd(path, j)
	}
	if !ok {
		return "", ErrLastSegNotExist
	}

	if f == l {
		return "", nil
	}

	if f > l {
		return "", ErrSegOrderReversed
	}

	return path[f:l], nil
}

// SubSeg receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as a string and a nil
// error. If any error is encountered, a zero value string and error are
// returned.
func SubSeg(path, key string, i int, v interface{}) error {
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

	default:
		err = ErrUnknownType
	}

	return err
}

// SubSpan receives a key which is used to search for the first
// matching path segment and an int value representing a second segment by it's
// distance from the matched segment, then returns the content between those
// segments as a string and a nil error. If any error is encountered, a zero
// value string and error are returned. The int representing a segment can be
// of negative values. Providing a 0 int is a special case which indicates the
// end of the path.
func SubSpan(path, key string, i, j int) (string, error) {
	si, ok := segIndexByKey(path, key)
	if !ok {
		return "", ErrKeyNotExist
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

// Parth holds path and error data for processing paths multiple times, and
// then checking for errors only once. Only the first encountered error will be
// returned.
type Parth struct {
	path string
	err  error
}

// New receives a path as a string, then returns a new Parth set with the
// provided path.
func New(path string) *Parth {
	return &Parth{path: path}
}

// NewFromSpan receives a path as a string, and two int values representing
// path segments, then returns a new Parth set with the content between those
// segments, and any error encountered. See Span for more info.
func NewFromSpan(path string, i, j int) *Parth {
	s, err := Span(path, i, j)
	return &Parth{s, err}
}

// NewFromSubSpan receives a path as a string, a key which is used to search
// for the first matching path segment, and an int value representing a second
// segment by it's distance from the matched segment, then returns a new Parth
// set with the content between those segments, and any error encountered. See
// SubSpan for more info.
func NewFromSubSpan(path, key string, i, j int) *Parth {
	s, err := SubSpan(path, key, i, j)
	return &Parth{s, err}
}

// Err returns the first error encountered by the Parth instance.
func (p *Parth) Err() error {
	return p.err
}

// Segment receives an int representing a path segment, then returns
// the specified segment as a string. If any error is encountered, a zero value
// string is returned, and the Parth instance's err value is set. If an error
// has already been set, a zero value string is returned.
func (p *Parth) Segment(i int, v interface{}) {
	if p.err != nil {
		return
	}

	p.err = Segment(p.path, i, v)
}

// Sequent ...
func (p *Parth) Sequent(key string, v interface{}) {
	p.SubSeg(key, 0, v)
}

// Span receives two int values representing path segments, then
// returns the content between those segments, including the first segment, as
// a string. If any error is encountered, a zero value string is returned, and
// the Parth instance's err value is set. If an error has already been set, a
// zero value string is returned. See Span for more info.
func (p *Parth) Span(i, j int) string {
	if p.err != nil {
		return ""
	}

	s, err := Span(p.path, i, j)
	p.err = err

	return s
}

// SubSeg receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as a string. If any error
// is encountered, a zero value string is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value string is
// returned.
func (p *Parth) SubSeg(key string, i int, v interface{}) {
	if p.err != nil {
		return
	}

	p.err = SubSeg(p.path, key, i, v)
}

// SubSpan receives a key which is used to search for the first
// matching path segment, and an int value representing a second segment by
// it's distance from the matched segment, then returns the content between
// those segments as a string. If an error has already been set, a zero value
// string is returned. See SubSpan for more info.
func (p *Parth) SubSpan(key string, i, j int) string {
	if p.err != nil {
		return ""
	}

	s, err := SubSpan(p.path, key, i, j)
	p.err = err

	return s
}
