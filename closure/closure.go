package closure

import "fmt"

// () func() intは引数なし、返り値は関数で、関数の返り値はintという意味
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}

}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func Closure() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}
