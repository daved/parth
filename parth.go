// Package parth provides a simple API for accessing path segments.
// Accessing multiple segments may produce errors at any stage, so be mindful
// of when they are checked.
package parth

import (
	"fmt"
	"strconv"
	"strings"
)

// Parser holds path segment info, and err for accumulated errors.
type Parser struct {
	segments []string
	len		int
	err      error
}

// New receives a path, and returns a new Parser.
func New(path string) *Parser {
	s := strings.Split(strings.TrimLeft(path, "/"), "/")
	return &Parser{segments: s, len:len(s)}
}

// Err allows errors to be checked more flexibly.
func (p *Parser) Err() error {
	return p.err
}

// String receives an int representing a segment, and returns the specified
// segment as a string, or returns empty and sets p.err upon any failure.
// Because paths start as strings, this method is used by other methods.
func (p *Parser) String(i int) string {
	if p.err != nil {
		return ""
	}
	if i < 0 || i >= p.len {
		p.err = fmt.Errorf("%d is out of bounds", i)
		return ""
	}
	return p.segments[i]
}

// Int64 receives an int representing a segment, and returns the specified
// segment as an int64, or returns 0 and sets p.err upon any failure.
func (p *Parser) Int64(i int) int64 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		p.err = err
		return 0
	}
	return v
}

// Int32 receives an int representing a segment, and returns the specified
// segment as an int32, or returns 0 and sets p.err upon any failure.
func (p *Parser) Int32(i int) int32 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		p.err = err
		return 0
	}
	return int32(v)
}

// Int16 receives an int representing a segment, and returns the specified
// segment as an int16, or returns 0 and sets p.err upon any failure.
func (p *Parser) Int16(i int) int16 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		p.err = err
		return 0
	}
	return int16(v)
}

// Int8 receives an int representing a segment, and returns the specified
// segment as an int8, or returns 0 and sets p.err upon any failure.
func (p *Parser) Int8(i int) int8 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		p.err = err
		return 0
	}
	return int8(v)
}

// Int receives an int representing a segment, and returns the specified
// segment as an int, or returns 0 and sets p.err upon any failure.
func (p *Parser) Int(i int) int {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		p.err = err
		return 0
	}
	return int(v)
}

// Bool receives an int representing a segment, and returns the specified
// segment as a bool, or returns 0 and sets p.err upon any failure.
func (p *Parser) Bool(i int) bool {
	s := p.String(i)
	if p.err != nil {
		return false
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		p.err = err
		return false
	}
	return v
}

// Float64 receives an int representing a segment, and returns the specified
// segment as an float64, or returns 0 and sets p.err upon any failure.
func (p *Parser) Float64(i int) float64 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		p.err = err
		return 0
	}
	return v
}

// Float32 receives an int representing a segment, and returns the specified
// segment as a float32, or returns 0 and sets p.err upon any failure.
func (p *Parser) Float32(i int) float32 {
	s := p.String(i)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		p.err = err
		return 0
	}
	return float32(v)
}
