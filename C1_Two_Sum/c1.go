package c1

func twoSum(nums []int, target int) []int {
	diffToIndex := make(map[int]int)
	for i, num := range nums {
		if val, ok := diffToIndex[target-num]; ok {
			return []int{val, i}
		} else {
			diffToIndex[num] = i
		}
	}
	return nil
}
