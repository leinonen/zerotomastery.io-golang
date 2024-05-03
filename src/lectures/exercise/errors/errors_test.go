package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{time: "14:07:33", ok: true},
		{time: "0:59:59", ok: true},
		{time: "bad", ok: false},
		{time: "11:22", ok: false},
		{time: "", ok: false},
		{time: "1:-4:4", ok: false},
	}
	for _, data := range table {

		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v error should be nil", data.time, err)
		}
	}
}
