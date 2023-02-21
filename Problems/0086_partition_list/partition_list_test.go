package problem0086

import (
	"Leetcode_Solution/Kit"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: Kit.MakeListNode([]int{1, 4, 3, 2, 5, 2}), x: 3}, want: Kit.MakeListNode([]int{1, 2, 2, 4, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
