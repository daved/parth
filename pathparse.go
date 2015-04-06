package pathparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	segments []string
	err      error
}

func New(path string) *Parser {
	return &Parser{
		segments: strings.Split(strings.TrimLeft(path, "/"), "/"),
	}
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) String(i int) string {
	if p.err != nil {
		return ""
	}
	if 0 > i || i >= len(p.segments) {
		p.err = fmt.Errorf("%d is out of bounds", i)
		return ""
	}
	return p.segments[i]
}

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
