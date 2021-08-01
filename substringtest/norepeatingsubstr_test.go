package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},
		//Chinese support
		{"这里风景如画", 6},
		{"一二三三二一", 3},
		{"黑化黑灰化肥黑灰会挥发发灰黑化肥黑灰化肥挥发", 7},
	}

	for _, tt := range tests {
		actual := lengthOfLongestSubString(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s, expected %d\n", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化黑灰化肥黑灰会挥发发灰黑化肥黑灰化肥挥发", 7
	for i := 0; i < b.N; i++ {
		actual := lengthOfLongestSubString(s)
		if actual != ans {
			b.Errorf("got %d for input %s, expected %d\n", actual, s, ans)
		}
	}
}
