package c7reverseinteger

func reverse(x int) int {
	res := 0
	for val := x; val >= 1 || val <= -1; val /= 10 {
		res = res*10 + val%10
	}
	return res
}
