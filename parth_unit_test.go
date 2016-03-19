package parth

import "testing"

var (
	errFmtGotWant  = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr   = "Did not receive expected err for segment value %v"
	errFmtUnexpErr = "Received unexpected err for segment type %T: %v"
)

func TestUnitPosSegToString(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		s     string
		isErr bool
	}{
		{0, "/test1", "test1", false},
		{1, "/test1/test-2", "test-2", false},
		{2, "/test1/test-2/test_3/", "test_3", false},
		{0, "test4/t4", "test4", false},
		{1, "//test5", "test5", false},
		{1, "/test6//", "", true},
		{3, "/test7", "", true},
		{0, "//test8", "", true},
		{0, "/", "", true},
	}

	for _, v := range tests {
		s, err := posSegToString(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, want)
		}
	}
}

func TestUnitNegSegToString(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		s     string
		isErr bool
	}{
		{-1, "/test1", "test1", false},
		{-1, "/test1/test-2", "test-2", false},
		{-2, "/test1/test-2", "test1", false},
		{-3, "/test1/test-2/test_3", "test1", false},
		{-1, "test4/t4/", "t4", false},
		{-1, "//test5", "test5", false},
		{-1, "/test6//", "", true},
		{-3, "/test7", "", true},
		{-2, "//test8", "", true},
		{-1, "/", "", true},
	}

	for _, v := range tests {
		s, err := negSegToString(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, want)
		}
	}
}

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

func TestUnitPosSegAsIndex(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{0, "/test1", 0, false},
		{1, "/test1/test-2", 6, false},
		{2, "/test1/test-2/test_3", 13, false},
		{0, "test3/t3/", 0, false},
		{1, "test4/t4/", 5, false},
		{6, "/t5/f/fiv/55/5/fi/ve", 17, false},
		{0, "/", 0, false},
		{1, "/", 1, false},
		{2, "/", 0, true},
		{4, "/test/out", 0, true},
	}

	for _, v := range tests {
		i, err := posSegAsIndex(v.s, v.ind)
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

func TestUnitNegSegAsIndex(t *testing.T) {
	tests := []struct {
		ind   int
		s     string
		i     int
		isErr bool
	}{
		{-1, "/test1", 6, false},
		{-1, "/test1/test-2", 13, false},
		{-2, "/test1/test-2/test_3", 13, false},
		{0, "test3/t3/", 0, true},
		{-1, "test4/t4/", 8, false},
		{-6, "/t5/f/fiv/55/5/fi/ve", 5, false},
		{-1, "/", 1, false},
		{-2, "/", 0, false},
		{-4, "/test/out", 0, true},
	}

	for _, v := range tests {
		i, err := negSegAsIndex(v.s, v.ind)
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
