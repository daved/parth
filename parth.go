// Package parth provides functions for accessing path segments.
//
// When returning an int of any size, the first whole number within the
// specified segment will be returned. When returning a float of any size,
// the first decimal number within the specified segment will be returned.
package parth

import (
	"errors"
	"strconv"
	"unicode"
)

// Err{Name} values are for error identification.
var (
	ErrFirstSegNotExist = errors.New("first segment index does not exist")
	ErrLastSegNotExist  = errors.New("last segment index does not exist")
	ErrSegOrderReversed = errors.New("first segment must precede last segment")
	ErrSegNotExist      = errors.New("segment index does not exist")
	ErrIntNotFound      = errors.New("segment does not contain int")
	ErrFloatNotFound    = errors.New("segment does not contain float")
	ErrUnparsable       = errors.New("unable to parse segment")

	ErrIndexBad      = errors.New("index cannot be found")
	ErrIndexNotFound = errors.New("no index found")
)

// SegmentToString receives an int representing a path segment, then returns
// both the specified segment as a string and a nil error. If any error is
// encountered, a zero value string and error are returned.
func SegmentToString(path string, i int) (string, error) {
	s, err := SpanToString(path, i, i+1)
	if err != nil {
		return "", err
	}

	if s[0] == '/' {
		return s[1:], nil
	}

	return s, nil
}

func segmentToIntN(path string, i, size int) (int64, error) {
	s, err := segToStrInt(path, i)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(s, 10, size)
	if err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
}

// SegmentToInt64 receives an int representing a path segment, then returns
// both the specified segment as an int64 and a nil error. If any error is
// encountered, a zero value int64 and error are returned.
func SegmentToInt64(path string, i int) (int64, error) {
	return segmentToIntN(path, i, 64)
}

// SegmentToInt32 receives an int representing a path segment, then returns
// both the specified segment as an int32 and a nil error. If any error is
// encountered, a zero value int32 and error are returned.
func SegmentToInt32(path string, i int) (int32, error) {
	v, err := segmentToIntN(path, i, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}

// SegmentToInt16 receives an int representing a path segment, then returns
// both the specified segment as an int16 and a nil error. If any error is
// encountered, a zero value int16 and error are returned.
func SegmentToInt16(path string, i int) (int16, error) {
	v, err := segmentToIntN(path, i, 16)
	if err != nil {
		return 0, err
	}

	return int16(v), nil
}

// SegmentToInt8 receives an int representing a path segment, then returns both
// the specified segment as an int8 and a nil error. If any error is
// encountered, a zero value int8 and error are returned.
func SegmentToInt8(path string, i int) (int8, error) {
	v, err := segmentToIntN(path, i, 8)
	if err != nil {
		return 0, err
	}

	return int8(v), nil
}

// SegmentToInt receives an int representing a path segment, then returns both
// the specified segment as an int and a nil error. If any error is
// encountered, a zero value int and error are returned.
func SegmentToInt(path string, i int) (int, error) {
	v, err := segmentToIntN(path, i, 0)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

// SegmentToBool receives an int representing a path segment, then returns both
// the specified segment as a bool and a nil error. If any error is
// encountered, a zero value bool and error are returned.
func SegmentToBool(path string, i int) (bool, error) {
	s, err := SegmentToString(path, i)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, ErrUnparsable
	}

	return v, nil
}

// SegmentToFloat64 receives an int representing a path segment, then returns
// both the specified segment as a float64 and a nil error. If any error is
// encountered, a zero value float64 and error are returned.
func SegmentToFloat64(path string, i int) (float64, error) {
	s, err := segToStrFloat(path, i)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return v, nil
}

// SegmentToFloat32 receives an int representing a path segment, then returns
// both the specified segment as a float32 and a nil error. If any error is
// encountered, a zero value float32 and error are returned.
func SegmentToFloat32(path string, i int) (float32, error) {
	s, err := segToStrFloat(path, i)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return float32(v), nil
}

// SubSegToString receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as a string and a nil
// error. If any error is encountered, a zero value string and error are
// returned.
func SubSegToString(path, key string) (string, error) {
	ki, err := segIndexByKey(path, key)
	if err != nil {
		return "", err
	}

	s, err := SegmentToString(path[ki:], 1)
	if err != nil {
		return "", err
	}

	return s, nil
}

func subSegToIntN(path, key string, size int) (int64, error) {
	s, err := subSegToStrInt(path, key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(s, 10, size)
	if err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
}

// SubSegToInt64 receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as an int64 and a nil
// error. If any error is encountered, a zero value int64 and error are
// returned.
func SubSegToInt64(path, key string) (int64, error) {
	return subSegToIntN(path, key, 64)
}

// SubSegToInt32 receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as an int32 and a nil
// error. If any error is encountered, a zero value int32 and error are
// returned.
func SubSegToInt32(path, key string) (int32, error) {
	v, err := subSegToIntN(path, key, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}

// SubSegToInt16 receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as an int16 and a nil
// error. If any error is encountered, a zero value int16 and error are
// returned.
func SubSegToInt16(path, key string) (int16, error) {
	v, err := subSegToIntN(path, key, 16)
	if err != nil {
		return 0, err
	}

	return int16(v), nil
}

// SubSegToInt8 receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as an int8 and a nil
// error. If any error is encountered, a zero value int8 and error are
// returned.
func SubSegToInt8(path, key string) (int8, error) {
	v, err := subSegToIntN(path, key, 8)
	if err != nil {
		return 0, err
	}

	return int8(v), nil
}

// SubSegToInt receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as an int and a nil
// error. If any error is encountered, a zero value int and error are
// returned.
func SubSegToInt(path, key string) (int, error) {
	v, err := subSegToIntN(path, key, 0)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

// SubSegToBool receives a key which is used to search for the first matching
// path segment, then returns both the subsequent segment as a bool and a nil
// error. If any error is encountered, a zero value bool and error are
// returned.
func SubSegToBool(path, key string) (bool, error) {
	s, err := SubSegToString(path, key)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, ErrUnparsable
	}

	return v, nil
}

// SubSegToFloat64 receives a key which is used to search for the first
// matching path segment, then returns both the subsequent segment as a float64
// and a nil error. If any error is encountered, a zero value float64 and error
// are returned.
func SubSegToFloat64(path, key string) (float64, error) {
	s, err := subSegToStrFloat(path, key)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return v, nil
}

// SubSegToFloat32 receives a key which is used to search for the first
// matching path segment, then returns both the subsequent segment as a float32
// and a nil error. If any error is encountered, a zero value float32 and error
// are returned.
func SubSegToFloat32(path, key string) (float32, error) {
	s, err := subSegToStrFloat(path, key)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return float32(v), nil
}

// SpanToString receives two int values representing path segments, and
// returns the content between those segments, including the first segment, as
// a string and a nil error. If any error is encountered, a zero value string
// and error are returned. The segments can be of negative values, but the
// first segment must come before the last segment. Providing a 0 int for the
// second int is a special case which indicates the end of the path.
func SpanToString(path string, firstSeg, lastSeg int) (string, error) {
	var f, l int
	var err error

	if firstSeg < 0 {
		f, err = segStartIndexFromEnd(path, firstSeg)
	} else {
		f, err = segStartIndexFromStart(path, firstSeg)
	}
	if err != nil {
		return "", ErrFirstSegNotExist
	}

	if lastSeg > 0 {
		l, err = segEndIndexFromStart(path, lastSeg)
	} else {
		l, err = segEndIndexFromEnd(path, lastSeg)
	}
	if err != nil {
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

// SubSpanToString receives a key which is used to search for the first
// matching path segment and an int value representing a second segment by it's
// distance from the matched segment, then returns the content between those
// segments as a string and a nil error. If any error is encountered, a zero
// value string and error are returned. The int representing a segment can be
// of negative values. Providing a 0 int is a special case which indicates the
// end of the path.
func SubSpanToString(path, key string, lastSeg int) (string, error) {
	ki, err := segIndexByKey(path, key)
	if err != nil {
		return "", err
	}

	if lastSeg > 0 {
		lastSeg++
	}

	s, err := SpanToString(path[ki:], 1, lastSeg)
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
// segments, and any error encountered. See SpanToString for more info.
func NewFromSpan(path string, firstSeg, lastSeg int) *Parth {
	s, err := SpanToString(path, firstSeg, lastSeg)
	return &Parth{s, err}
}

// NewFromSubSpan receives a path as a string, a key which is used to search
// for the first matching path segment, and an int value representing a second
// segment by it's distance from the matched segment, then returns a new Parth
// set with the content between those segments, and any error encountered. See
// SubSpanToString for more info.
func NewFromSubSpan(path, key string, lastSeg int) *Parth {
	s, err := SubSpanToString(path, key, lastSeg)
	return &Parth{s, err}
}

// Err returns the first error encountered by the Parth instance.
func (p *Parth) Err() error {
	return p.err
}

// SegmentToString receives an int representing a path segment, then returns
// the specified segment as a string. If any error is encountered, a zero value
// string is returned, and the Parth instance's err value is set. If an error
// has already been set, a zero value string is returned.
func (p *Parth) SegmentToString(i int) string {
	if p.err != nil {
		return ""
	}

	s, err := SegmentToString(p.path, i)

	p.err = err
	return s
}

// SegmentToInt64 receives an int representing a path segment, then returns the
// specified segment as an int64. If any error is encountered, a zero value
// int64 is returned, and the Parth instance's err value is set. If an error
// has already been set, a zero value int64 is returned.
func (p *Parth) SegmentToInt64(i int) int64 {
	if p.err != nil {
		return 0
	}

	n, err := SegmentToInt64(p.path, i)

	p.err = err
	return n
}

// SegmentToInt32 receives an int representing a path segment, then returns the
// specified segment as an int32. If any error is encountered, a zero value
// int32 is returned, and the Parth instance's err value is set. If an error
// has already been set, a zero value int32 is returned.
func (p *Parth) SegmentToInt32(i int) int32 {
	if p.err != nil {
		return 0
	}

	n, err := SegmentToInt32(p.path, i)

	p.err = err
	return n
}

// SegmentToInt16 receives an int representing a path segment, then returns the
// specified segment as an int16. If any error is encountered, a zero value
// int16 is returned, and the Parth instance's err value is set. If an error
// has already been set, a zero value int16 is returned.
func (p *Parth) SegmentToInt16(i int) int16 {
	if p.err != nil {
		return 0
	}

	n, err := SegmentToInt16(p.path, i)

	p.err = err
	return n
}

// SegmentToInt8 receives an int representing a path segment, then returns the
// specified segment as an int8. If any error is encountered, a zero value int8
// is returned, and the Parth instance's err value is set. If an error has
// already been set, a zero value int8 is returned.
func (p *Parth) SegmentToInt8(i int) int8 {
	if p.err != nil {
		return 0
	}

	n, err := SegmentToInt8(p.path, i)

	p.err = err
	return n
}

// SegmentToInt receives an int representing a path segment, then returns the
// specified segment as an int. If any error is encountered, a zero value int
// is returned, and the Parth instance's err value is set. If an error has
// already been set, a zero value int is returned.
func (p *Parth) SegmentToInt(i int) int {
	if p.err != nil {
		return 0
	}

	n, err := SegmentToInt(p.path, i)

	p.err = err
	return n
}

// SegmentToBool receives an int representing a path segment, then returns the
// specified segment as a bool. If any error is encountered, a zero value bool
// is returned, and the Parth instance's err value is set. If an error has
// already been set, a zero value bool is returned.
func (p *Parth) SegmentToBool(i int) bool {
	if p.err != nil {
		return false
	}

	b, err := SegmentToBool(p.path, i)

	p.err = err
	return b
}

// SegmentToFloat64 receives an int representing a path segment, then returns
// the specified segment as a float64. If any error is encountered, a zero
// value float64 is returned, and the Parth instance's err value is set. If an
// error has already been set, a zero value float64 is returned.
func (p *Parth) SegmentToFloat64(i int) float64 {
	if p.err != nil {
		return 0
	}

	f, err := SegmentToFloat64(p.path, i)

	p.err = err
	return f
}

// SegmentToFloat32 receives an int representing a path segment, then returns
// the specified segment as a float32. If any error is encountered, a zero
// value float32 is returned, and the Parth instance's err value is set. If an
// error has already been set, a zero value float32 is returned.
func (p *Parth) SegmentToFloat32(i int) float32 {
	if p.err != nil {
		return 0
	}

	f, err := SegmentToFloat32(p.path, i)

	p.err = err
	return f
}

// SubSegToString receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as a string. If any error
// is encountered, a zero value string is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value string is
// returned.
func (p *Parth) SubSegToString(key string) string {
	if p.err != nil {
		return ""
	}

	s, err := SubSegToString(p.path, key)

	p.err = err
	return s
}

// SubSegToInt64 receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an int64. If any error
// is encountered, a zero value int64 is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value int64 is
// returned.
func (p *Parth) SubSegToInt64(key string) int64 {
	if p.err != nil {
		return 0
	}

	i, err := SubSegToInt64(p.path, key)

	p.err = err
	return i
}

// SubSegToInt32 receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an int32. If any error
// is encountered, a zero value int32 is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value int32 is
// returned.
func (p *Parth) SubSegToInt32(key string) int32 {
	if p.err != nil {
		return 0
	}

	i, err := SubSegToInt32(p.path, key)

	p.err = err
	return i
}

// SubSegToInt16 receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an int16. If any error
// is encountered, a zero value int16 is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value int16 is
// returned.
func (p *Parth) SubSegToInt16(key string) int16 {
	if p.err != nil {
		return 0
	}

	i, err := SubSegToInt16(p.path, key)

	p.err = err
	return i
}

// SubSegToInt8 receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an int8. If any error
// is encountered, a zero value int8 is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value int8 is
// returned.
func (p *Parth) SubSegToInt8(key string) int8 {
	if p.err != nil {
		return 0
	}

	i, err := SubSegToInt8(p.path, key)

	p.err = err
	return i
}

// SubSegToInt receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an int. If any error is
// encountered, a zero value int8 is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value int8 is
// returned.
func (p *Parth) SubSegToInt(key string) int {
	if p.err != nil {
		return 0
	}

	i, err := SubSegToInt(p.path, key)

	p.err = err
	return i
}

// SubSegToBool receives a key which is used to search for the first matching
// path segment, then returns the subsequent segment as an bool. If any error
// is encountered, a zero value bool is returned, and the Parth instances err
// value is set. If an error has already been set, a zero value bool is
// returned.
func (p *Parth) SubSegToBool(key string) bool {
	if p.err != nil {
		return false
	}

	b, err := SubSegToBool(p.path, key)

	p.err = err
	return b
}

// SubSegToFloat64 receives a key which is used to search for the first
// matching path segment, then returns the subsequent segment as an float64. If
// any error is encountered, a zero value float64 is returned, and the Parth
// instance's err value is set. If an error has already been set, a zero value
// float64 is returned.
func (p *Parth) SubSegToFloat64(key string) float64 {
	if p.err != nil {
		return 0
	}

	f, err := SubSegToFloat64(p.path, key)

	p.err = err
	return f
}

// SubSegToFloat32 receives a key which is used to search for the first
// matching path segment, then returns the subsequent segment as an float32. If
// any error is encountered, a zero value float32 is returned, and the Parth
// instance's err value is set. If an error has already been set, a zero value
// float32 is returned.
func (p *Parth) SubSegToFloat32(key string) float32 {
	if p.err != nil {
		return 0
	}

	f, err := SubSegToFloat32(p.path, key)

	p.err = err
	return f
}

// SpanToString receives two int values representing path segments, then
// returns the content between those segments, including the first segment, as
// a string. If any error is encountered, a zero value string is returned, and
// the Parth instance's err value is set. If an error has already been set, a
// zero value string is returned. See SpanToString for more info.
func (p *Parth) SpanToString(firstSeg, lastSeg int) string {
	if p.err != nil {
		return ""
	}

	s, err := SpanToString(p.path, firstSeg, lastSeg)

	p.err = err
	return s
}

// SubSpanToString receives a key which is used to search for the first
// matching path segment, and an int value representing a second segment by
// it's distance from the matched segment, then returns the content between
// those segments as a string. If an error has already been set, a zero value
// string is returned. See SubSpanToString for more info.
func (p *Parth) SubSpanToString(key string, lastSeg int) string {
	if p.err != nil {
		return ""
	}

	s, err := SubSpanToString(p.path, key, lastSeg)

	p.err = err
	return s
}

func segStartIndexFromStart(path string, seg int) (int, error) {
	if seg < 0 {
		return 0, ErrIndexBad
	}

	for n, ct := 0, 0; n < len(path); n++ {
		if n > 0 && path[n] == '/' {
			ct++
		}

		if ct == seg {
			return n, nil
		}
	}

	return 0, ErrIndexNotFound
}

func segStartIndexFromEnd(path string, seg int) (int, error) {
	if seg > -1 {
		return 0, ErrIndexBad
	}

	for n, ct := len(path)-1, 0; n >= 0; n-- {
		if path[n] == '/' || n == 0 {
			ct--
		}

		if ct == seg {
			return n, nil
		}
	}

	return 0, ErrIndexNotFound
}

func segEndIndexFromStart(path string, seg int) (int, error) {
	if seg < 1 {
		return 0, ErrIndexBad
	}

	for n, ct := 0, 0; n < len(path); n++ {
		if path[n] == '/' && n > 0 {
			ct++
		}

		if ct == seg {
			return n, nil
		}

		if n+1 == len(path) && ct+1 == seg {
			return n + 1, nil
		}
	}

	return 0, ErrIndexNotFound
}

func segEndIndexFromEnd(path string, seg int) (int, error) {
	if seg > 0 {
		return 0, ErrIndexBad
	}

	if seg == 0 {
		return len(path), nil
	}

	if len(path) == 1 && path[0] == '/' {
		return 0, nil
	}

	for n, ct := len(path)-1, 0; n >= 0; n-- {
		if n == 0 || path[n] == '/' {
			ct--
		}

		if ct == seg {
			return n, nil
		}

	}

	return 0, ErrIndexNotFound
}

func segIndexByKey(path, key string) (int, error) {
	if path == "" || key == "" {
		return 0, ErrUnparsable
	}

	for n := 0; n < len(path); n++ {
		si, err := segStartIndexFromStart(path, n)
		if err != nil {
			return 0, ErrSegNotExist
		}

		if len(path[si:]) == len(key)+1 {
			if path[si+1:] == key {
				return si, nil
			}

			return 0, ErrSegNotExist
		}

		tmpEI, err := segStartIndexFromStart(path[si:], 1)
		if err != nil {
			return 0, ErrSegNotExist
		}

		if path[si+1:tmpEI+si] == key || n == 0 && path[0] != '/' && path[si:tmpEI+si] == key {
			return si, nil
		}
	}

	return 0, nil
}

func segToStrInt(path string, i int) (string, error) {
	s, err := SegmentToString(path, i)
	if err != nil {
		return "", err
	}

	if s, err = firstIntFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func segToStrFloat(path string, i int) (string, error) {
	s, err := SegmentToString(path, i)
	if err != nil {
		return "", err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func subSegToStrInt(path, key string) (string, error) {
	s, err := SubSegToString(path, key)
	if err != nil {
		return "", err
	}

	if s, err = firstIntFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func subSegToStrFloat(path, key string) (string, error) {
	s, err := SubSegToString(path, key)
	if err != nil {
		return "", err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func firstIntFromString(s string) (string, error) {
	ind, l := 0, 0

	for n := 0; n < len(s); n++ {
		if unicode.IsDigit(rune(s[n])) {
			if l == 0 {
				ind = n
			}

			l++
		} else if s[n] == '-' {
			if l == 0 {
				ind = n
				l++
			} else {
				break
			}
		} else {
			if l == 0 && s[n] == '.' {
				if n+1 < len(s) && unicode.IsDigit(rune(s[n+1])) {
					return "0", nil
				}

				break
			}

			if l > 0 {
				break
			}
		}
	}

	if l == 0 {
		return "", ErrIntNotFound
	}

	return s[ind : ind+l], nil
}

func firstFloatFromString(s string) (string, error) {
	c, ind, l := 0, 0, 0

	for n := 0; n < len(s); n++ {
		if unicode.IsDigit(rune(s[n])) {
			if l == 0 {
				ind = n
			}

			l++
		} else if s[n] == '-' {
			if l == 0 {
				ind = n
				l++
			} else {
				break
			}
		} else if s[n] == '.' {
			if l == 0 {
				ind = n
			}

			if c > 0 {
				break
			}

			l++
			c++
		} else if s[n] == 'e' && l > 0 && n+1 < len(s) && s[n+1] == '+' {
			l++
		} else if s[n] == '+' && l > 0 && s[n-1] == 'e' {
			if n+1 < len(s) && unicode.IsDigit(rune(s[n+1])) {
				l++
				continue
			}

			l--
			break
		} else {
			if l > 0 {
				break
			}
		}
	}

	if l == 0 || s[ind:ind+l] == "." {
		return "", ErrFloatNotFound
	}

	return s[ind : ind+l], nil
}
