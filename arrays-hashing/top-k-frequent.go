package main

import "sort"

// Buckets Method O(n)
func TopKFrequent(nums []int, k int) []int {
	output := make([]int, 0)
	//make map
	m := make(map[int]int, 0)
	for i := range nums {
		m[nums[i]] += 1
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range m {
		buckets[count] = append(buckets[count], num)
	}

	//Iterate from highest to lowest
	for count := len(buckets) - 1; count >= 0; count-- {
		output = append(output, buckets[count]...)
	}

	return output[:k]
}

// sorting method O(nlogn)
type pair struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)

	// Count frequency of each number.
	for _, num := range nums {
		frequency[num]++
	}

	// Convert the map into a slice.
	items := make([]pair, 0, len(frequency))

	for num, count := range frequency {
		items = append(items, pair{
			num:   num,
			count: count,
		})
	}

	// Sort by frequency in descending order.
	sort.Slice(items, func(i, j int) bool {
		return items[i].count > items[j].count
	})

	// Take the first k numbers.
	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, items[i].num)
	}

	return result
}
