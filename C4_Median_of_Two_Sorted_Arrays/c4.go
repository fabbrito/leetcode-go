package c4

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n1, n2 := len(nums1), len(nums2)
	low, high := 0, n1

	for low <= high {
		partitionX := low + (high-low)/2
		partitionY := (n1+n2+1)/2 - partitionX

		maxX := math.MinInt64
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}

		minX := math.MaxInt64
		if partitionX < n1 {
			minX = nums1[partitionX]
		}

		maxY := math.MinInt64
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}

		minY := math.MaxInt64
		if partitionY < n2 {
			minY = nums2[partitionY]
		}

		if maxX <= minY && maxY <= minX {
			if (n1+n2)%2 == 0 {
				return (float64(max(maxX, maxY) + min(minX, minY))) / 2.0
			}
			return float64(max(maxX, maxY))
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}
