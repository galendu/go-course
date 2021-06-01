package ifstmt

import (
	"fmt"
	"testing"
)

func TestIfOne(t *testing.T) {
	/* 局部变量定义 */
	var a int = 100

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func TestIfM(t *testing.T) {
	var age int = 18
	if age < 18 {
		fmt.Println("nice")
	} else if age < 28 {
		fmt.Println("beauty")
	} else if age < 38 {
		fmt.Println("sexy")
	} else {
		fmt.Println("next")
	}
}

func TestNested(t *testing.T) {
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}

func TestDeepNested(t *testing.T) {
	// 身高1.8m以上, 25 ~ 35岁, 男
	var (
		height float32
		age    uint
		gender string
		passed bool
	)

	height = 1.9
	age = 30
	gender = "male"

	if height > 1.8 {
		if age > 25 && age <= 35 {
			if gender == "male" {
				passed = true
			}
		}
	}

	if passed {
		fmt.Println("congratulations! your successed")
	} else {
		fmt.Println("not passed")
	}
}

func TestIfEnum(t *testing.T) {
	const (
		Unknown = iota
		Male
		Female
	)

	gender := 0

	if gender == Unknown {

	} else if gender == Male {

	} else if gender == Female {

	} else {
		fmt.Println()
	}
}
