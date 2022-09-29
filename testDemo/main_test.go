package main

import "testing"

func TestFib(t *testing.T) {
	testTables := []struct {
		in     int64
		expect int64
	}{
		{
			in:     1,
			expect: 1,
		}, {
			in:     2,
			expect: 2,
		}, {
			in:     3,
			expect: 3,
		}, {
			in:     4,
			expect: 5,
		},
		{
			in:     -1,
			expect: 0,
		},
	}
	//sdsdsdsdsds
	for _, e := range testTables {
		res := fib(e.in)
		if res != e.expect {
			t.Errorf("n: %d, actual: %d, expect: %d", e.in, res, e.expect)
		}
	}
}
