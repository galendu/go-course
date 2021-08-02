package sort

import (
	"sort"
)

func BubbleSort(numbers []int) []int {
	for range numbers {
		for j := 0; j < len(numbers)-1; j++ {
			// 当前值 numbers[i], 后一个值是多少 numbers[j+1]
			// fmt.Printf("数据: 当前: %d, 比对: %d\n", numbers[j], numbers[j+1])

			// 比较2个数, 交换顺序, 大数沉底, 小数冒出
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}

			// fmt.Printf("第%d趟: %v\n", i+1, numbers)
	}
	return numbers
}

func SelectSort(numbers []int) []int {
	for i := range numbers {
		// 拿到第一个的数, 就是numbers[i], 比如 3
		// fmt.Printf("第%d趟: %d\n", i+1, numbers[i])

		// 依次和后面相邻的数比较
		for j := i + 1; j < len(numbers); j++ {
			// fmt.Printf("数据 -->  当前数据: %d, 比对数据: %d\n", numbers[i], numbers[j])
			if numbers[i] > numbers[j] {
				// 如果当前数 > 后面的数据, 则交换位置
				numbers[i], numbers[j] = numbers[j], numbers[i]
				// fmt.Printf("交换 -->  当前数据: %d, 比对数据: %d\n", numbers[i], numbers[j])
			}
		}

		// fmt.Println("结果: ", numbers)
	}

	// fmt.Println("最终结果", numbers)
	return numbers
}

func BuildInSort(numbers []int) []int {
	sort.Sort(IntSlice(numbers))
	return numbers
}

func NewIntSlice(numbers []int) IntSlice {
	return IntSlice(numbers)
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
