package _876

import (
	"Leetcode_Solution/Kit"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{Kit.MakeListNode([]int{1, 2, 3, 4, 5})}, want: &ListNode{Val: 3}},
		{name: "test2", args: args{Kit.MakeListNode([]int{1, 2, 3, 4, 5, 6})}, want: &ListNode{Val: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
