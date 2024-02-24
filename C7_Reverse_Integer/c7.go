package c7reverseinteger

import "math"

func reverse(x int) int {
	res := 0
	for val := x; val >= 1 || val <= -1; val /= 10 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		res = res*10 + val%10
	}
	return res
}
