package c5

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 0", args: args{s: "babad"}, want: "bab"},
		{name: "Example 1", args: args{s: "cbbd"}, want: "bb"},
		{name: "Example 2", args: args{s: "aasdabccbaaad"}, want: "abccba"},
		{name: "Example 3", args: args{s: "a"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromeManacher(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 0", args: args{s: "babad"}, want: "bab"},
		{name: "Example 1", args: args{s: "cbbd"}, want: "bb"},
		{name: "Example 2", args: args{s: "aasdabccbaaad"}, want: "abccba"},
		{name: "Example 3", args: args{s: "a"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeManacher(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeManacher() = %v, want %v", got, tt.want)
			}
		})
	}
}
