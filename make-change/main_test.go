package main

import (
	"fmt"
	"testing"
)

func TestMakeChange(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{1, 1},  // 1
		{6, 2},  // 5 + 1
		{47, 5}, // 25 + 10 + 10 + 1 + 1
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("makeChangeNaive(%d)", tt.arg), func(t *testing.T) {
			if got := makeChangeNaive(tt.arg); got != tt.want {
				t.Errorf("makeChangeNaive(%d) = %d, want %d", tt.arg, got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("makeChangeMemo(%d)", tt.arg), func(t *testing.T) {
			if got := makeChangeMemo(tt.arg); got != tt.want {
				t.Errorf("makeChangeMemo(%d) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func BenchmarkMakeChange(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func(int) int
		args []int
	}{
		{"makeChangeNaive", makeChangeNaive, []int{1, 6, 47}},
		{"makeChangeMemo", makeChangeMemo, []int{1, 6, 47}},
	}
	for _, bb := range benchmarks {
		for _, arg := range bb.args {
			b.Run(fmt.Sprintf("%s(%d)", bb.name, arg), func(b *testing.B) {
				// b.ResetTimer()
				// b.ReportAllocs()
				for n := 0; n < b.N; n++ {
					bb.f(arg)
				}
			})
		}
	}
}

var coins = []int{25, 10, 5, 1}

// brute-force solution
// time complexity: O(n^c), where
//   c - amount of change (branching factor),
//   n - number of coin denominations (height of the tree)
// space complexity: O(n)
func makeChangeNaive(n int) int {
	if n == 0 {
		return 0
	}
	minCoins := n
	for _, c := range coins {
		if n-c < 0 {
			continue
		}
		min := makeChangeNaive(n - c)
		if min < minCoins {
			minCoins = min
		}
	}
	return minCoins + 1
}

// top-down solution
// time complexity: O(c*n)
// space complexity: O(n)
func makeChangeMemo(n int) int {
	cache := make([]int, n+1)
	return makeChangeMemoRec(n, cache)
}

func makeChangeMemoRec(n int, cache []int) int {
	if cache[n] == 0 {
		if n == 0 {
			return 0
		}
		minCoins := n
		for _, c := range coins {
			if n-c < 0 {
				continue
			}
			min := makeChangeMemoRec(n-c, cache)
			if min < minCoins {
				minCoins = min
			}
		}
		cache[n] = minCoins + 1
	}
	return cache[n]
}
