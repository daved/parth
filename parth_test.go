package parth_test

import (
	"strings"
	"testing"

	"github.com/codemodus/parth"
)

func TestStringComplex(t *testing.T) {
	tests := [][]string{
		[]string{"/test", "/nest", "/best", "/rest", "/zest", "/jest", ""},
		[]string{"/abc-def", "/uvw_xyz/"},
		[]string{"/", "/test"},
		[]string{"test", "/best"},
	}

	for _, v := range tests {
		var path string
		for _, y := range v {
			path += y
		}
		po := parth.New(path)

		for k, v := range v {
			seg, err := po.String(k)
			if err != nil && v != "" && v != "/" {
				t.Fatal(err)
			}

			want := strings.TrimSuffix(strings.TrimPrefix(v, "/"), "/")
			got := seg
			if got != want {
				t.Errorf("Segment Value = %v, want %v", got, want)
			}
		}
	}
}

func TestInts(t *testing.T) {
	ePath := "/error"
	tests := []string{"/0", "/1", "/2", ePath}

	var path string
	for _, v := range tests {
		path += v
	}
	po := parth.New(path)

	for k, v := range tests {
		seg, err := po.Int(k)
		if err != nil && v != ePath {
			t.Fatal(err)
		}

		want := int(k)
		if v == ePath {
			want = int(0)
		}
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}

	for k, v := range tests {
		seg, err := po.Int8(k)
		if err != nil && v != ePath {
			t.Fatal(err)
		}

		want := int8(k)
		if v == ePath {
			want = int8(0)
		}
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}

	for k, v := range tests {
		seg, err := po.Int16(k)
		if err != nil && v != ePath {
			t.Fatal(err)
		}

		want := int16(k)
		if v == ePath {
			want = int16(0)
		}
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}

	for k, v := range tests {
		seg, err := po.Int32(k)
		if err != nil && v != ePath {
			t.Fatal(err)
		}

		want := int32(k)
		if v == ePath {
			want = int32(0)
		}
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}

	for k, v := range tests {
		seg, err := po.Int64(k)
		if err != nil && v != ePath {
			t.Fatal(err)
		}

		want := int64(k)
		if v == ePath {
			want = int64(0)
		}
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}
}

func TestBool(t *testing.T) {
	ePath := "/error"
	var tests = []struct {
		p string
		b bool
	}{
		{"/true", true},
		{ePath, false},
		{"/false", false},
		{"/0", false},
		{"/1", true},
	}

	var path string
	for _, v := range tests {
		path += v.p
	}
	po := parth.New(path)

	for k, v := range tests {
		seg, err := po.Bool(k)
		if err != nil && v.p != ePath {
			t.Fatal(err)
		}

		want := v.b
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}
}

func TestFloats(t *testing.T) {
	ePath := "/error"
	var tests = []struct {
		p   string
		f32 float32
		f64 float64
	}{
		{"/0.0", 0.0, 0.0},
		{"/1.2", 1.2, 1.2},
		{"/2", 2.0, 2.0},
		{ePath, 0.0, 0.0},
	}

	var path string
	for _, v := range tests {
		path += v.p
	}
	po := parth.New(path)

	for k, v := range tests {
		seg, err := po.Float32(k)
		if err != nil && v.p != ePath {
			t.Fatal(err)
		}

		want := v.f32
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}

	for k, v := range tests {
		seg, err := po.Float64(k)
		if err != nil && v.p != ePath {
			t.Fatal(err)
		}

		want := v.f64
		got := seg
		if got != want {
			t.Errorf("Segment Value = %v, want %v", got, want)
		}
	}
}
