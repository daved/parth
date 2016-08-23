package parth_test

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/codemodus/parth"
)

var (
	bmri int
	bmrs string
)

func standardSegment(path string, i int) (int, error) {
	ss := strings.Split(strings.TrimLeft(path, "/"), "/")

	if len(ss) == 0 || i > len(ss) {
		err := fmt.Errorf("segment out of bounds")
		return 0, err
	}

	v, err := strconv.ParseInt(ss[i], 10, 0)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func BenchmarkStandardInt(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = standardSegment(p, 1)
	}

	bmri = r
}

func BenchmarkParthInt(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SegmentToInt(p, 1)
	}

	bmri = r
}

func BenchmarkParthIntNeg(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SegmentToInt(p, -1)
	}

	bmri = r
}

func BenchmarkParthSubSeg(b *testing.B) {
	p := "/zero/1/2"
	k := "1"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SubSegToString(p, k)
	}

	bmrs = r
}

func BenchmarkStandardSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		cs := strings.Split(p, "/")
		if p[0] == '/' {
			cs[1] = "/" + cs[1]
		}
		r = path.Join(cs[0:3]...)
	}

	bmrs = r
}

func BenchmarkParthSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SpanToString(p, 0, 1)
	}

	bmrs = r
}

func BenchmarkParthSubSpan(b *testing.B) {
	p := "/zero/1/2"
	k := "zero"
	i := 2
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = parth.SubSpanToString(p, k, i)
	}

	bmrs = r
}

func BenchmarkVsCtxParthString2x(b *testing.B) {
	p := "/thing/123"
	var r0, r1 string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0, _ = parth.SegmentToString(p, 1)
		r1, _ = parth.SegmentToString(p, 1)
	}

	bmrs = r0
	bmrs = r1
}

func BenchmarkVsCtxParthString3x(b *testing.B) {
	p := "/thing/123"
	var r0, r1, r2 string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0, _ = parth.SegmentToString(p, 1)
		r1, _ = parth.SegmentToString(p, 1)
		r2, _ = parth.SegmentToString(p, 1)
	}

	bmrs = r0
	bmrs = r1
	bmrs = r2
}

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

func newParams(val string) Params {
	return Params{
		Param{
			Key:   "id",
			Value: val,
		},
	}
}

func BenchmarkVsCtxContextGetSetGet(b *testing.B) {
	ps := newParams("123")
	var r0, r1 string
	req, _ := http.NewRequest("GET", "", nil)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0 = ps.ByName("id")

		ctx := context.WithValue(req.Context(), "id", r0)
		req = req.WithContext(ctx)

		r1 = req.Context().Value("id").(string)
	}

	bmrs = r0
	bmrs = r1
}

func BenchmarkVsCtxContextGetSetGetGet(b *testing.B) {
	ps := newParams("123")
	var r0, r1, r2 string
	req, _ := http.NewRequest("GET", "", nil)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0 = ps.ByName("id")

		ctx := context.WithValue(req.Context(), "id", r0)
		req = req.WithContext(ctx)

		r1 = req.Context().Value("id").(string)
		r2 = req.Context().Value("id").(string)
	}

	bmrs = r0
	bmrs = r1
	bmrs = r2
}
