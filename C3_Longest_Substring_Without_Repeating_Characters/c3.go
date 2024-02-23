package c3

func lengthOfLongestSubstring(s string) int {
	asciiArr := [128]int{}

	left, right, maxLen := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		asciiArr[c]++

		for asciiArr[c] > 1 {
			d := s[left]
			left++
			asciiArr[d]--
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}

	return maxLen
}

func lengthOfLongestSubstringMap(s string) int {
	result := 0
	// our "hashset" to keep track of characters weve seen
	seen := map[byte]bool{}

	// pointer for the "start" of our sliding window
	start := 0

	for end := 0; end < len(s); end++ {
		// keep moving start forward, as so long as
		// this current character at end has been seen before
		for seen[s[end]] {
			// remove whats at start, and "pinch" the sliding window
			seen[s[start]] = false
			start++
		}

		// we can put this one here now, its safe
		seen[s[end]] = true

		// now here, we are guaranteed to have a "window" with no repeaters
		if end-start+1 > result {
			result = end - start + 1
		}
	}

	return result
}
