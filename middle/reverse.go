package middle

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	MaxInt := int(math.Pow(2, 31))
	MinInt := -int(math.Pow(2, 31))
	if x > MaxInt-1 || x < MinInt {
		return res
	}
	for {
		lastNum := x % 10
		if res > MaxInt/10 || (res == MaxInt/10 && lastNum > 7) {
			return 0
		}
		if res < MinInt/10 || (res == MinInt/10 && lastNum < -8) {
			return 0
		}
		res = res*10 + lastNum
		x = x / 10
		if x == 0 {
			break
		}
	}

	return res
}

func main() {
	a := int(math.Pow(2, 31)) - 10
	fmt.Println(a)

	fmt.Println(reverse(a))
}
