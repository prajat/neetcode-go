package main

/*Given an array nums of unique integers, return all the possible permutations. You may return the answer in any order. */
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	currentPath := make([]int, 0)
	picked := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(currentPath) == len(nums) {
			path := make([]int, len(currentPath))
			copy(path, currentPath)
			result = append(result, path)
			return
		}

		for i := 0; i < len(nums); i++ {
			if picked[i] {
				continue
			}

			picked[i] = true
			currentPath = append(currentPath, nums[i])

			dfs()

			currentPath = currentPath[:len(currentPath)-1]
			picked[i] = false
		}
	}
	dfs()
	return result
}
