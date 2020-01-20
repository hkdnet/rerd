package main

import "testing"

func Test_endWith(t *testing.T) {
	for _, testcase := range []struct {
		s, target string
		expected  bool
	}{
		{"abc", "c", true},
		{"abc", "d", false},
		{"abc", "b", false},
	} {
		if endWith(testcase.s, testcase.target) != testcase.expected {
			t.Errorf("endWith(%v, %v) should return %v\n", testcase.s, testcase.target, testcase.expected)
		}
	}
}
