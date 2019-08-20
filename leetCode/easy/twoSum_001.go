package easy

//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]



// 方法一：暴力法，O(n) = n^2
func TwoSum(nums []int, target int)  []int{
	for i, v := range nums{
		for j := i + 1; j < len(nums); j++ {
			if v + nums[j] == target {
				return []int{i,j}
			}
		}
	}

	return nil
}

// 方法二: hash
func TwoSum2(nums []int, target int) []int  {
	m := make(map[int]int, len(nums))
	for i, v := range nums{
		sub := target - v
		if j,ok := m[sub];ok {
			return []int{j,i}
		}else {
			m[v] = i
		}
	}
	return nil
}