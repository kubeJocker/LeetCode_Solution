package problem0142

import (
	"Leetcode_Solution/Kit"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: Kit.MakeListNodeListWithCycle([]int{1, 2}, -1)}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
