package cast

import (
	"fmt"
	"strconv"
)

func Cast() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v, %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v, %d\n", yy, yy, yy)

	// goの型変換は厳密なためint("14")のようなキャストができない
	var s string = "14"
	// ASCII To Integer
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)
}
