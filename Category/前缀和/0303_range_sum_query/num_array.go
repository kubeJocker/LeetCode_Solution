package _303_range_sum_query

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{preSum: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.preSum[right] - this.preSum[left-1]
	}
	return this.preSum[right]
}
