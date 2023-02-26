package _380_insert_delete_get_random_O_1_

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	nums := []int{}
	valToIndex := make(map[int]int)
	return RandomizedSet{nums: nums, valToIndex: valToIndex}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	index := this.valToIndex[val]
	last := len(this.nums) - 1
	this.valToIndex[this.nums[last]] = index
	this.nums[index], this.nums[last] = this.nums[last], this.nums[index]
	this.nums = this.nums[:last]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
