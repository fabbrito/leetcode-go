package c6zigzagconversion

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	res := make([]uint8, len(s), len(s))
	colStep := numRows*2 - 2
	diagStep := colStep - 2
	idx := 0
	for row := 0; row < numRows; row += 1 {
		for i := row; i < len(s); i += colStep {
			res[idx] = s[i]
			idx += 1
			if row > 0 && row < numRows-1 && i+diagStep < len(s) {
				res[idx] = s[i+diagStep]
				idx += 1
			}
		}
		if row > 0 && row < numRows-1 {
			diagStep -= 2
		}
	}

	return string(res)
}
