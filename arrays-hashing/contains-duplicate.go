package main

func hasDuplicate(nums []int) bool {
	hashMap := make(map[int]bool)
	for i, _ := range nums {
		_, ok := hashMap[nums[i]]
		if ok {
			return true
		}
		hashMap[nums[i]] = true
	}
	return false
}
