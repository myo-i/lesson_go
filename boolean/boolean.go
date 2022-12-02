package boolean

import "fmt"

func Boolean() {
	t, f := true, false

	fmt.Printf("%T %v %t %t\n", t, t, t, 1)
	fmt.Printf("%T %v %t %t\n", f, f, f, 1)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
}
