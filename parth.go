// Package parth provides functions for accessing path segments.
//
// When returning an int/uint of any size, the first whole number within the
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
	ErrUnknownType = errors.New("unknown type provided")

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

func segmentToBool(path string, i int) (bool, error) {
	s, err := segmentToString(path, i)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, ErrUnparsable
	}

	return v, nil
}

func segmentToFloatN(path string, i, size int) (float64, error) {
	s, err := segToStrFloat(path, i)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, size)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return v, nil
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

func segmentToString(path string, i int) (string, error) {
	j := i + 1
	if i < 0 {
		i--
	}

	s, err := Span(path, i, j)
	if err != nil {
		return "", err
	}

	if s[0] == '/' {
		s = s[1:]
	}

	return s, nil
}

func segmentToUintN(path string, i, size int) (uint64, error) {
	s, err := segToStrUint(path, i)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(s, 10, size)
	if err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
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
	var err error

	if i < 0 {
		f, err = segStartIndexFromEnd(path, i)
	} else {
		f, err = segStartIndexFromStart(path, i)
	}
	if err != nil {
		return "", ErrFirstSegNotExist
	}

	if j > 0 {
		l, err = segEndIndexFromStart(path, j)
	} else {
		l, err = segEndIndexFromEnd(path, j)
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

func subSegToBool(path, key string, i int) (bool, error) {
	s, err := subSegToString(path, key, i)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, ErrUnparsable
	}

	return v, nil
}

func subSegToFloatN(path, key string, i, size int) (float64, error) {
	s, err := subSegToStrFloat(path, key, i)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, size)
	if err != nil {
		return 0.0, ErrUnparsable
	}

	return v, nil
}

func subSegToIntN(path, key string, i, size int) (int64, error) {
	s, err := subSegToStrInt(path, key, i)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(s, 10, size)
	if err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
}

func subSegToString(path, key string, i int) (string, error) {
	ki, err := segIndexByKey(path, key)
	if err != nil {
		return "", err
	}

	i++

	s, err := segmentToString(path[ki:], i)
	if err != nil {
		return "", err
	}

	return s, nil
}

func subSegToUintN(path, key string, i, size int) (uint64, error) {
	s, err := subSegToStrUint(path, key, i)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(s, 10, size)
	if err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
}

// SubSpan receives a key which is used to search for the first
// matching path segment and an int value representing a second segment by it's
// distance from the matched segment, then returns the content between those
// segments as a string and a nil error. If any error is encountered, a zero
// value string and error are returned. The int representing a segment can be
// of negative values. Providing a 0 int is a special case which indicates the
// end of the path.
func SubSpan(path, key string, i, j int) (string, error) {
	si, err := segIndexByKey(path, key)
	if err != nil {
		return "", err
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

func segToStrUint(path string, i int) (string, error) {
	s, err := segmentToString(path, i)
	if err != nil {
		return "", err
	}

	if s, err = firstUintFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func segToStrInt(path string, i int) (string, error) {
	s, err := segmentToString(path, i)
	if err != nil {
		return "", err
	}

	if s, err = firstIntFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func segToStrFloat(path string, i int) (string, error) {
	s, err := segmentToString(path, i)
	if err != nil {
		return "", err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func subSegToStrUint(path, key string, i int) (string, error) {
	s, err := subSegToString(path, key, i)
	if err != nil {
		return "", err
	}

	if s, err = firstUintFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func subSegToStrInt(path, key string, i int) (string, error) {
	s, err := subSegToString(path, key, i)
	if err != nil {
		return "", err
	}

	if s, err = firstIntFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func subSegToStrFloat(path, key string, i int) (string, error) {
	s, err := subSegToString(path, key, i)
	if err != nil {
		return "", err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return "", err
	}

	return s, nil
}

func firstUintFromString(s string) (string, error) {
	ind, l := 0, 0

	for n := 0; n < len(s); n++ {
		if unicode.IsDigit(rune(s[n])) {
			if l == 0 {
				ind = n
			}

			l++
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
