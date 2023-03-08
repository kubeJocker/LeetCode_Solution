package problem0023

import (
	. "Leetcode_Solution/Kit"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{lists: MakeListNodeList([][]int{
			[]int{1, 4, 5},
			[]int{1, 3, 4},
			[]int{2, 6},
		})}, want: MakeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKListsByQueue(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{lists: MakeListNodeList([][]int{
			[]int{1, 4, 5},
			[]int{1, 3, 4},
			[]int{2, 6},
		})}, want: MakeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsByQueue(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsByQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
