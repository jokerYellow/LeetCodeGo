package printBraces

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test(t, 6)
	test(t, 2)
	test(t, 3)
	test(t, 0)
}

func test(t *testing.T, input int) {
	fmt.Println(printBraces(input))
}
