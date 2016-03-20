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

// SegmentToString receives an int representing a path segment, and returns both
// the specified segment as a string and a nil error. If any error is
// encountered, a zero value string and error are returned.
func SegmentToString(path string, i int) (string, error) {
	if i >= 0 {
		return posSegToString(path, i)
	}

	return negSegToString(path, i)
}

// SegmentToInt64 receives an int representing a path segment, and returns both
// the specified segment as an int64 and a nil error. If any error is
// encountered, a zero value int64 and error are returned.
func SegmentToInt64(path string, i int) (int64, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}

	if s, err = firstIntFromString(s); err != nil {
		return 0, err
	}

	var v int64

	if v, err = strconv.ParseInt(s, 10, 64); err != nil {
		return 0, ErrUnparsable
	}

	return v, nil
}

// SegmentToInt32 receives an int representing a path segment, and returns both
// the specified segment as an int32 and a nil error. If any error is
// encountered, a zero value int32 and error are returned.
func SegmentToInt32(path string, i int) (int32, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}

	if s, err = firstIntFromString(s); err != nil {
		return 0, err
	}

	var v int64

	if v, err = strconv.ParseInt(s, 10, 32); err != nil {
		return 0, ErrUnparsable
	}

	return int32(v), nil
}

// SegmentToInt16 receives an int representing a path segment, and returns both
// the specified segment as an int16 and a nil error. If any error is
// encountered, a zero value int16 and error are returned.
func SegmentToInt16(path string, i int) (int16, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}

	if s, err = firstIntFromString(s); err != nil {
		return 0, err
	}

	var v int64

	if v, err = strconv.ParseInt(s, 10, 16); err != nil {
		return 0, ErrUnparsable
	}

	return int16(v), nil
}

// SegmentToInt8 receives an int representing a path segment, and returns both
// the specified segment as an int8 and a nil error. If any error is
// encountered, a zero value int8 and error are returned.
func SegmentToInt8(path string, i int) (int8, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}

	if s, err = firstIntFromString(s); err != nil {
		return 0, err
	}

	var v int64

	if v, err = strconv.ParseInt(s, 10, 8); err != nil {
		return 0, ErrUnparsable
	}

	return int8(v), nil
}

// SegmentToInt receives an int representing a path segment, and returns both
// the specified segment as an int and a nil error. If any error is
// encountered, a zero value int and error are returned.
func SegmentToInt(path string, i int) (int, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}

	if s, err = firstIntFromString(s); err != nil {
		return 0, err
	}

	var v int64

	if v, err = strconv.ParseInt(s, 10, 0); err != nil {
		return 0, ErrUnparsable
	}

	return int(v), nil
}

// SegmentToBool receives an int representing a path segment, and returns both
// the specified segment as a bool and a nil error. If any error is
// encountered, a zero value bool and error are returned.
func SegmentToBool(path string, i int) (bool, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return false, err
	}

	var v bool

	if v, err = strconv.ParseBool(s); err != nil {
		return false, ErrUnparsable
	}

	return v, nil
}

// SegmentToFloat64 receives an int representing a path segment, and returns
// both the specified segment as a float64 and a nil error. If any error is
// encountered, a zero value float64 and error are returned.
func SegmentToFloat64(path string, i int) (float64, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0.0, err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return 0.0, err
	}

	var v float64

	if v, err = strconv.ParseFloat(s, 64); err != nil {
		return 0.0, ErrUnparsable
	}

	return v, nil
}

// SegmentToFloat32 receives an int representing a path segment, and returns
// both the specified segment as a float32 and a nil error. If any error is
// encountered, a zero value float32 and error are returned.
func SegmentToFloat32(path string, i int) (float32, error) {
	var s string
	var err error

	if s, err = SegmentToString(path, i); err != nil {
		return 0.0, err
	}

	if s, err = firstFloatFromString(s); err != nil {
		return 0.0, err
	}

	var v float64

	if v, err = strconv.ParseFloat(s, 32); err != nil {
		return 0.0, ErrUnparsable
	}

	return float32(v), nil
}

// SubSegToString receives a key which is used to search for the first matching
// path segment, and returns both the subsequent segment as a string and a nil
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
// distance from the matched segment, and returns the content between those
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

func posSegToString(path string, i int) (string, error) {
	c, ind0, ind1 := 0, 0, 0

	for n := 0; n < len(path); n++ {
		if path[n] == '/' {
			if c == i {
				if n+1 < len(path) && path[n+1] != '/' {
					ind0 = n + 1
				} else {
					break
				}
			}

			if c > i {
				ind1 = n
				break
			}
			c++
		} else if n == 0 {
			if c == i {
				ind0 = n
			}

			c++
		} else if n == len(path)-1 {
			if c > i {
				ind1 = n + 1
			}

			break
		}
	}

	if i < 0 || ind1 == 0 {
		return "", ErrSegNotExist
	}

	return path[ind0:ind1], nil
}

func negSegToString(path string, i int) (string, error) {
	i = i * -1
	c, ind0, ind1 := 1, 0, 0

	for n := len(path) - 1; n >= 0; n-- {
		if path[n] == '/' {
			if c == i {
				if n-1 >= 0 && path[n-1] != '/' {
					ind1 = n
				} else {
					break
				}
			}

			if c > i {
				ind0 = n + 1

				break
			}

			c++
		} else if n == len(path)-1 {
			if c == i {
				ind1 = n + 1
			}

			c++
		} else if n == 0 {
			if c > i {
				ind0 = n + 1
			}

			break
		}
	}

	if i < 1 || ind0 == 0 {
		return "", ErrSegNotExist
	}

	return path[ind0:ind1], nil
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
