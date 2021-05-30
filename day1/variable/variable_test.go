package variable

import (
	"testing"
)

func TestConst(t *testing.T) {
	const (
		a = "sdf"
	)
	a = "bbb"
	t.Log(a)
}

func TestShort(t *testing.T) {
	a, b, c := "value1", 10, 0.01
	t.Log(a, b, c)
}

func TestMergeMulti(t *testing.T) {
	var (
		a string  = "value1"
		b int     = 10
		c float32 = 0.01
	)
	t.Log(a, b, c)
}

func TestMerge(t *testing.T) {
	var a string = "one line"
	t.Log(a)
}

func TestMulti(t *testing.T) {
	var (
		a string
		b int
	)
	a, b = "string", 10
	t.Log(a, b)
}

func TestVariable(t *testing.T) {
	var a string
	a = "test"
	t.Log(a)
}

func TestFloat64(t *testing.T) {
	var a float64
	t.Log(a)
}

func TestStringNull(t *testing.T) {
	var a string
	t.Log(a)
}

func TestCharNull(t *testing.T) {
	var a rune
	t.Log(a)
}
