package selectio

import (
	"fmt"
	"testing"
)

func TestBasic1(t *testing.T) {
	fmt.Println("don't cache1234124")
	Basic1()
}

func TestBasic2(t *testing.T) {
	Basic2()
}

func TestSelectOrder(t *testing.T) {
	SelectOrder()
}
