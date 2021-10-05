package dp

//REference for implementation and learning - https://www.youtube.com/watch?v=73r3KWiEvyk
//https://leetcode.com/problems/house-robber/submissions/
func rob(nums []int) int {
	rob1 := 0
	rob2 := 0
	temp := 0
	// if n is current pos,  rob1 is max till n-2; rob2 is max till  n-1
	for i := 0; i < len(nums); i++ {
		if rob1+nums[i] > rob2 {
			temp = rob1 + nums[i]
		} else {
			temp = rob2
		}

		rob1 = rob2
		rob2 = temp

	}
	return rob2

}
