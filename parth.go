// Package parth provides a simple API for accessing path segments.
package parth

import (
	"fmt"
	"strconv"
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
	c, i1, i2 := 0, 0, 0
	for n := 0; n < len(p.path); n++ {
		if p.path[n] == '/' {
			if c == i {
				if n+1 < len(p.path) && p.path[n+1] != '/' {
					i1 = n + 1
				} else {
					break
				}
			}
			if c > i {
				i2 = n
				break
			}
			c++
		} else if n == 0 {
			if c == i {
				i1 = n
			}
			c++
		} else if n == len(p.path)-1 {
			if c > i {
				i2 = n + 1
			}
			break
		}
	}
	if i < 0 || i2 == 0 {
		return "", fmt.Errorf("path segment index %d does not exist", i)
	}
	return p.path[i1:i2], nil
}

// Int64 receives an int representing a path segment, and returns both the
// specified segment as an int64 and a nil error.  If any error is encountered,
// a zero value int64 and error are returned.
func (p *Parth) Int64(i int) (v int64, err error) {
	if s, err := p.String(i); err == nil {
		if v, err = strconv.ParseInt(s, 10, 64); err == nil {
			return v, nil
		}
	}

	return 0, err
}

// Int32 receives an int representing a path segment, and returns both the
// specified segment as an int32 and a nil error.  If any error is encountered,
// a zero value int32 and error are returned.
func (p *Parth) Int32(i int) (_ int32, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseInt(s, 10, 32); err == nil {
			return int32(v), nil
		}
	}

	return 0, err
}

// Int16 receives an int representing a path segment, and returns both the
// specified segment as an int16 and a nil error.  If any error is encountered,
// a zero value int16 and error are returned.
func (p *Parth) Int16(i int) (_ int16, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseInt(s, 10, 16); err == nil {
			return int16(v), nil
		}
	}

	return 0, err
}

// Int8 receives an int representing a path segment, and returns both the
// specified segment as an int8 and a nil error.  If any error is encountered,
// a zero value int8 and error are returned.
func (p *Parth) Int8(i int) (_ int8, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseInt(s, 10, 8); err == nil {
			return int8(v), nil
		}
	}

	return 0, err
}

// Int receives an int representing a path segment, and returns both the
// specified segment as an int and a nil error.  If any error is encountered,
// a zero value int and error are returned.
func (p *Parth) Int(i int) (_ int, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseInt(s, 10, 0); err == nil {
			return int(v), nil
		}
	}

	return 0, err
}

// Bool receives an int representing a path segment, and returns both the
// specified segment as a bool and a nil error.  If any error is encountered,
// a zero value bool and error are returned.
func (p *Parth) Bool(i int) (v bool, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseBool(s); err == nil {
			return v, nil
		}
	}

	return false, err
}

// Float64 receives an int representing a path segment, and returns both the
// specified segment as a float64 and a nil error.  If any error is encountered,
// a zero value float64 and error are returned.
func (p *Parth) Float64(i int) (v float64, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			return v, nil
		}
	}

	return 0.0, err
}

// Float32 receives an int representing a path segment, and returns both the
// specified segment as a float32 and a nil error.  If any error is encountered,
// a zero value float32 and error are returned.
func (p *Parth) Float32(i int) (_ float32, err error) {
	if s, err := p.String(i); err == nil {
		if v, err := strconv.ParseFloat(s, 32); err == nil {
			return float32(v), nil
		}
	}

	return 0.0, err
}
