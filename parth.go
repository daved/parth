// Package parth provides a simple API for accessing path segments.
//
// When returning an int of any size, the first whole number within the
// specified segment will be returned.  When returning a float of any size,
// the first decimal number within the specified segment will be returned.
package parth

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Parth holds the path for parsing.
type Parth struct {
	path string
}

// New receives a path, and returns a new Parth object.
func New(path string) *Parth {
	return &Parth{path: path}
}

// String receives an int representing a path segment, and returns both the
// specified segment as a string and a nil error.  If any error is encountered,
// a zero value string and error are returned.
func (p *Parth) String(i int) (string, error) {
	c, ind0, ind1 := 0, 0, 0
	for n := 0; n < len(p.path); n++ {
		if p.path[n] == '/' {
			if c == i {
				if n+1 < len(p.path) && p.path[n+1] != '/' {
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
		} else if n == len(p.path)-1 {
			if c > i {
				ind1 = n + 1
			}
			break
		}
	}
	if i < 0 || ind1 == 0 {
		return "", fmt.Errorf("path segment index %d does not exist", i)
	}
	return p.path[ind0:ind1], nil
}

// Int64 receives an int representing a path segment, and returns both the
// specified segment as an int64 and a nil error.  If any error is encountered,
// a zero value int64 and error are returned.
func (p *Parth) Int64(i int) (int64, error) {
	var s string
	var v int64
	var err error
	if s, err = p.String(i); err != nil {
		return 0, err
	}
	if s, err = getIntString(s); err != nil {
		return 0, err
	}
	if v, err = strconv.ParseInt(s, 10, 64); err != nil {
		return 0, err
	}
	return v, nil
}

// Int32 receives an int representing a path segment, and returns both the
// specified segment as an int32 and a nil error.  If any error is encountered,
// a zero value int32 and error are returned.
func (p *Parth) Int32(i int) (int32, error) {
	var s string
	var v int64
	var err error
	if s, err = p.String(i); err != nil {
		return 0, err
	}
	if s, err = getIntString(s); err != nil {
		return 0, err
	}
	if v, err = strconv.ParseInt(s, 10, 32); err != nil {
		return 0, err
	}
	return int32(v), nil
}

// Int16 receives an int representing a path segment, and returns both the
// specified segment as an int16 and a nil error.  If any error is encountered,
// a zero value int16 and error are returned.
func (p *Parth) Int16(i int) (int16, error) {
	var s string
	var v int64
	var err error
	if s, err = p.String(i); err != nil {
		return 0, err
	}
	if s, err = getIntString(s); err != nil {
		return 0, err
	}
	if v, err = strconv.ParseInt(s, 10, 16); err != nil {
		return 0, err
	}
	return int16(v), nil
}

// Int8 receives an int representing a path segment, and returns both the
// specified segment as an int8 and a nil error.  If any error is encountered,
// a zero value int8 and error are returned.
func (p *Parth) Int8(i int) (int8, error) {
	var s string
	var v int64
	var err error
	if s, err = p.String(i); err != nil {
		return 0, err
	}
	if s, err = getIntString(s); err != nil {
		return 0, err
	}
	if v, err = strconv.ParseInt(s, 10, 8); err != nil {
		return 0, err
	}
	return int8(v), nil
}

// Int receives an int representing a path segment, and returns both the
// specified segment as an int and a nil error.  If any error is encountered,
// a zero value int and error are returned.
func (p *Parth) Int(i int) (int, error) {
	var s string
	var v int64
	var err error
	if s, err = p.String(i); err != nil {
		return 0, err
	}
	if s, err = getIntString(s); err != nil {
		return 0, err
	}
	if v, err = strconv.ParseInt(s, 10, 0); err != nil {
		return 0, err
	}
	return int(v), nil
}

// Bool receives an int representing a path segment, and returns both the
// specified segment as a bool and a nil error.  If any error is encountered,
// a zero value bool and error are returned.
func (p *Parth) Bool(i int) (bool, error) {
	var s string
	var v bool
	var err error
	if s, err = p.String(i); err != nil {
		return false, err
	}
	if v, err = strconv.ParseBool(s); err != nil {
		return false, err
	}
	return v, nil
}

// Float64 receives an int representing a path segment, and returns both the
// specified segment as a float64 and a nil error.  If any error is encountered,
// a zero value float64 and error are returned.
func (p *Parth) Float64(i int) (float64, error) {
	var s string
	var v float64
	var err error
	if s, err = p.String(i); err != nil {
		return 0.0, err
	}
	if s, err = getFloatString(s); err != nil {
		return 0.0, err
	}
	if v, err = strconv.ParseFloat(s, 64); err != nil {
		return 0.0, err
	}
	return v, nil
}

// Float32 receives an int representing a path segment, and returns both the
// specified segment as a float32 and a nil error.  If any error is encountered,
// a zero value float32 and error are returned.
func (p *Parth) Float32(i int) (float32, error) {
	var s string
	var v float64
	var err error
	if s, err = p.String(i); err != nil {
		return 0.0, err
	}
	if s, err = getFloatString(s); err != nil {
		return 0.0, err
	}
	if v, err = strconv.ParseFloat(s, 32); err != nil {
		return 0.0, err
	}
	return float32(v), nil
}

func getIntString(s string) (string, error) {
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
			if l == 0 && s[n] == '.' && unicode.IsDigit(rune(s[n+1])) {
				return "0", nil
			}
			if l > 0 {
				break
			}
		}
	}

	if l == 0 {
		return "", errors.New("path segment does not contain int")
	}
	return s[ind : ind+l], nil
}

func getFloatString(s string) (string, error) {
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
		} else {
			if l > 0 {
				break
			}
		}
	}

	if l == 0 || s[ind:ind+l] == "." {
		return "", errors.New("path segment does not contain float")
	}
	return s[ind : ind+l], nil
}
