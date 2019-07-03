package main

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		rules   []fbRule
		numbers []int
		want    string
	}{
		{[]fbRule{{2, "Two"}, {3, "Three"}}, []int{1, 2, 3, 4, 5, 6}, "1 Two Three Two 5 TwoThree"},
	}
	for _, c := range cases {
		got := answer(c.rules, c.numbers)
		if got != c.want {
			t.Errorf("answer(%v, %v) == %v, want %v", c.rules, c.numbers, got, c.want)
		}
	}
}
