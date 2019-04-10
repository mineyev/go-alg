package go_alg

func BinarySearch(hay []int, needle int) (int, bool) {
	low := 0
	hight := len(hay) - 1

	for low <= hight {
		var current = (low + hight) / 2

		if hay[current] == needle {
			return current, true
		}

		if needle < hay[current] {
			hight = current - 1
		}

		if needle > hay[current] {
			low = current + 1
		}
	}

	return 0, false
}

func RecursiveBinarySearch(haystack []int, needle int) (int, bool) {
	current := len(haystack) / 2
	var ok bool
	var result int

	if haystack[current] == needle {
		return current, true
	}

	if current == 0 {
		return current, false
	}

	if needle < haystack[current] {
		result, ok = RecursiveBinarySearch(haystack[0:current], needle)
	} else {
		result, ok = RecursiveBinarySearch(haystack[current:len(haystack)], needle)
		if ok == true {
			result = result + current
		}
	}

	return result, ok
}
