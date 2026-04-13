
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	// Buckets: index = frequency
	buckets := make([][]int, len(nums)+1)

	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}

	// Traverse from high freq to low
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
