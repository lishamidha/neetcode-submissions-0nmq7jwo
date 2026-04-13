
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	// Count frequencies
	for _, num := range nums {
		freqMap[num]++
	}

	// Convert map to slice
	type pair struct {
		num  int
		freq int
	}

	var pairs []pair
	for num, freq := range freqMap {
		pairs = append(pairs, pair{num, freq})
	}

	// Sort by frequency descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	// Collect top k
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}
