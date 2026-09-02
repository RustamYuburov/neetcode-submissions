func majorityElement(nums []int) int {
    counter, maxNum := 0, 0

	for _, n := range nums {
		if counter == 0 {
			maxNum = n
		}

		if n == maxNum {
			counter++
		} else {
			counter--
		}
	}

	return maxNum
}
