package parth

import "testing"

var (
	errFmtGotWant  = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr   = "Did not receive expected err for segment value %v"
	errFmtUnexpErr = "Received unexpected err for segment type %T: %v"
)

func TestUnitFirstIntFromString(t *testing.T) {
	var tests = []struct {
		s     string
		i     string
		isErr bool
	}{
		{"0.1", "0", false},
		{"0.2a", "0", false},
		{"aaaa1.3", "1", false},
		{"4", "4", false},
		{"5aaaa", "5", false},
		{"aaa6aa", "6", false},
		{".7.aaaa", "0", false},
		{".8aa", "0", false},
		{"-9", "-9", false},
		{"-9", "-9", false},
		{"10-", "10", false},
		{"3.14e+11", "3", false},
		{"3.14e.+12", "3", false},
		{"3.14e+.13", "3", false},
		{"3.14e+.13", "3", false},
		{".", "", true},
		{"error", "", true},
	}

	for _, v := range tests {
		i, err := firstIntFromString(v.s)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitFirstFloatFromString(t *testing.T) {
	tests := []struct {
		s     string
		f     string
		isErr bool
	}{
		{"/0.1", "0.1", false},
		{"/0.2a", "0.2", false},
		{"/aaaa1.3", "1.3", false},
		{"/4", "4", false},
		{"/5aaaa", "5", false},
		{"/aaa6aa", "6", false},
		{"/.7.aaaa", ".7", false},
		{"/.8aa", ".8", false},
		{"/-9", "-9", false},
		{"/10-", "10", false},
		{"/3.14e+11", "3.14e+11", false},
		{"/3.14e.+12", "3.14", false},
		{"/3.14e+.13", "3.14", false},
		{"/3.14e+.13", "3.14", false},
		{"/error", "", true},
		{"/.", "", true},
	}

	for _, v := range tests {
		f, err := firstFloatFromString(v.s)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, f, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.f
		got := f
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitSegStartIndexFromStart(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{0, "/test1", 0, false},
		{1, "/test1/test-2", 6, false},
		{2, "/t1-2/fd", 0, true},
		{2, "/test1/test-2/test_3", 13, false},
		{0, "test3/t3/", 0, false},
		{1, "test4/t4/", 5, false},
		{6, "/t5/f/fiv/55/5/fi/ve", 17, false},
		{0, "/", 0, false},
		{1, "/", 0, true},
		{2, "/", 0, true},
		{4, "/test/out", 0, true},
		{-1, "/test/out", 0, true},
		{2, "/0/1//", 4, false},
		{3, "/0/1//", 5, false},
	}

	for _, v := range tests {
		i, err := segStartIndexFromStart(v.s, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitSegStartIndexFromEnd(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{-1, "/test1", 0, false},
		{-1, "/test1/test-2", 6, false},
		{-1, "/test1/test-2/test_3", 13, false},
		{-3, "test3/t3/", 0, false},
		{-1, "test4/t4/", 8, false},
		{-2, "/t5/f/fiv/55/5/fi/ve", 14, false},
		{-1, "/", 0, false},
		{-2, "/", 0, true},
		{-4, "/test/out", 0, true},
		{1, "/test/out", 0, true},
	}

	for _, v := range tests {
		i, err := segStartIndexFromEnd(v.s, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitSegEndIndexFromStart(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{1, "/test1", 6, false},
		{2, "/test1/test-2", 13, false},
		{2, "/test1/test-2/test_3", 13, false},
		{1, "test3/t3/", 5, false},
		{2, "test4/t4/", 8, false},
		{5, "/t5/f/fiv/55/5/fi/ve", 14, false},
		{1, "/", 1, false},
		{-4, "/test/out", 0, true},
		{4, "/test/out", 0, true},
	}

	for _, v := range tests {
		i, err := segEndIndexFromStart(v.s, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitSegEndIndexFromEnd(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{0, "/test1", 6, false},
		{-1, "/t1", 0, false},
		{-1, "/test1/test-2", 6, false},
		{-2, "/test1/t-2/t_3", 6, false},
		{-3, "test3/t3/", 0, false},
		{-1, "test4/t4/", 8, false},
		{-2, "/t5/f/fiv/55/5/fi/ve", 14, false},
		{-1, "/", 0, false},
		{-4, "/test/out", 0, true},
		{4, "/test/out", 0, true},
	}

	for _, v := range tests {
		i, err := segEndIndexFromEnd(v.s, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestUnitSegIndexByKey(t *testing.T) {
	tests := []struct {
		k     string
		s     string
		i     int
		isErr bool
	}{
		{"test", "/1/test/3", 2, false},
		{"2", "/2/t/3", 0, false},
		{"3", "/1/test/3", 7, false},
		{"4", "/44/44/33", 0, true},
		{"best", "12/best/3", 2, false},
		{"6", "6/tt/66", 0, false},
		{"7", "1/test/7", 6, false},
		{"first", "first/2/three", 0, false},
		{"bad", "/ba/d/", 0, true},
		{"11", "/4/56/11/", 5, false},
		{"", "/4/56/11/", 0, true},
		{"t", "", 0, true},
	}

	for _, v := range tests {
		i, err := segIndexByKey(v.s, v.k)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.s)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}
