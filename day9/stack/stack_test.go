package stack_test

import (
	"testing"

	"gitee.com/infraboard/go-course/day9/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack.NewStack()
	s.Push(1)
	t.Log(s.Pop())
}

func TestStackOrder(t *testing.T) {
	should := assert.New(t)

	s := stack.NewStack()
	s.Push(9)
	s.Push(1)
	s.Push(0)
	s.Push(2)

	s.Order()
	should.Equal(s.Pop(), 9)
	should.Equal(s.Pop(), 2)
	should.Equal(s.Pop(), 1)
	should.Equal(s.Pop(), 0)
}
