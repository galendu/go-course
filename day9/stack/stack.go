package stack

import "fmt"

func NewNumberStack(numbers []int) *Stack {
	items := make([]Item, 0, len(numbers))
	for i := range numbers {
		items = append(items, numbers[i])
	}
	return &Stack{
		items: items,
	}
}

func NewStack() *Stack {
	return &Stack{
		items: []Item{},
	}
}

type Item interface{}

type Stack struct {
	items []Item
}

// Push adds an Item to the top of the stack
func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Clear() {
	s.items = []Item{}
}

func (s *Stack) Search(item Item) (pos int, err error) {
	for i := range s.items {
		if item == s.items[i] {
			return i, nil
		}
	}
	return 0, fmt.Errorf("item %s not found", item)
}

func (s *Stack) ForEach(fn func(Item)) {
	for i := range s.items {
		fn(i)
	}
}

// 把stack的自己完成排序
func (s *Stack) Sort() {
	// 准备一个辅助的stack, 另一个书堆容器
	orderdStack := NewStack()

	for !s.IsEmpty() {
		// 然后开始我们的排序流程
		current := s.Pop()

		// 当前元素大于右边, 应该把右边的罗过左边, 直到右边再无小于左边的元素
		for !orderdStack.IsEmpty() && current.(int) > orderdStack.Peek().(int) {
			s.Push(orderdStack.Pop())
		}

		// 此时 当前值 一定是 <= 右边的
		orderdStack.Push(current)
	}

	for !orderdStack.IsEmpty() {
		s.Push(orderdStack.Pop())
	}
}
