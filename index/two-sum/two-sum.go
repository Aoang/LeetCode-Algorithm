package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := tmp[target-v]; ok {
			return []int{j, i}
		}

		tmp[v] = i
	}

	return nil
}
