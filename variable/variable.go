package variable

import "fmt"

// varは関数外でも宣言できる
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func Variable() {
	fmt.Println(i, f64, s, t, f)
	//var f bool = false

	// ショートの宣言は関数内からでしか宣言できない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)

	xi = 3
	fmt.Println(xi, xf64, xs, xt, xf)

	//defaultVariabl()
}

func defaultVariabl() {
	var (
		i    int
		f64  float64
		s    string
		t, f bool
	)
	fmt.Println(i, f64, s, t, f)

}
