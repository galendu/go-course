package selectio

import (
	"fmt"
	"testing"
	"time"
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

func TestTimeAfter(t *testing.T) {
	fmt.Println(time.Now())

	a := time.After(1 * time.Second)
	fmt.Println(<-a)
}

func TestSelectTimeout(t *testing.T) {
	SelectTimeout()
}
