package problem0002

import (
	. "Leetcode_Solution/Kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func Test0002(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: MakeListNode([]int{2, 4, 3}),
				two: MakeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: MakeListNode([]int{7, 0, 8}),
			},
		},
		question{
			p: para{
				one: MakeListNode([]int{9, 8, 7, 6, 5}),
				two: MakeListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans{
				one: MakeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		question{
			p: para{
				one: MakeListNode([]int{0}),
				two: MakeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: MakeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}
}
