package parth

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"testing"
)

var (
	x interface{}
)

func stdSegmentInt(p string, i int) (int, error) {
	ss := strings.Split(strings.TrimLeft(p, "/"), "/")

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

func BenchmarkStdSegmentInt(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = stdSegmentInt(p, 1)
	}

	x = r
}

func BenchmarkSegmentToIntN(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t, _ := segmentToIntN(p, 1, 0)
		r = int(t)
	}

	x = r
}

func BenchmarkSegmentToIntNNeg(b *testing.B) {
	p := "/zero/1"
	var r int

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t, _ := segmentToIntN(p, -1, 0)
		r = int(t)
	}

	x = r
}

func BenchmarkSegmentToString(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r, _ = segmentToString(p, 1)
	}

	x = r
}
func stdSpan(p string, i, j int) string {
	cs := strings.Split(p, "/")

	if p[0] == '/' {
		cs[1] = "/" + cs[1]
	}

	return path.Join(cs[i:j]...)
}

func BenchmarkStdSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r = stdSpan(p, 0, 1)
	}

	x = r
}

func BenchmarkSpan(b *testing.B) {
	p := "/zero/1/2"
	var r string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r, _ = Span(p, 0, 1)
	}

	x = r
}

/*
func BenchmarkSubSpan(b *testing.B) {
	p := "/zero/1/2"
	k := "zero"
	i := 2
	var r string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r, _ = SubSpanToString(p, k, i)
	}

	bmrs = r
}

func BenchmarkVsCtxParthString2x(b *testing.B) {
	p := "/thing/123"
	var r0, r1 string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0, _ = SegmentToString(p, 1)
		r1, _ = SegmentToString(p, 1)
	}

	bmrs = r0
	bmrs = r1
}

func BenchmarkVsCtxParthString3x(b *testing.B) {
	p := "/thing/123"
	var r0, r1, r2 string

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r0, _ = SegmentToString(p, 1)
		r1, _ = SegmentToString(p, 1)
		r2, _ = SegmentToString(p, 1)
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
}*/
