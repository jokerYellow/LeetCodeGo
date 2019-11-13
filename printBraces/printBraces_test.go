package printBraces

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test(t, 6)
}

func test(t *testing.T, input int) {
	fmt.Println(printBraces(input))
}
