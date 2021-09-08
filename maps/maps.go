package maps

import "math"

func MapsBasics() {

	// initialize a map using make. Key is string and value is int
	districtSchools := make(map[string]int)
	districtSchools["Bangalore Urban"] = 1000
	districtSchools["Bangalore Rural"] = 800
	districtSchools["Mysuru"] = 810
	districtSchools["Shivamogga"] = 200

	println(districtSchools)

	//
}

//https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	mapIntCnt := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := mapIntCnt[v]; ok {
			return true
		} else {
			mapIntCnt[v] = 1
		}
	}

	return false
}

//explanation of this algorithm is at https://www.youtube.com/watch?v=5WZl3MMT0Eg&t=222s
//https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {

	maxSum := float64(nums[0])
	curSum := 0
	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		maxSum = math.Max(float64(maxSum), float64(curSum))
	}

	return int(maxSum)
}
