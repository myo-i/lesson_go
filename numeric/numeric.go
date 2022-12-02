package numeric

import "fmt"

func Numeric() {
	//var (
	//	u8  uint    = 255
	//	i8  int     = 127
	//	f32 float32 = 0.2
	//	// 複素数
	//	c64 complex64 = -5 + 12i
	//)
	//fmt.Println(u8, i8, f32, c64)
	//fmt.Printf("%T, %v", u8, u8)

	// シフト
	fmt.Println(1 << 0) // 0001 0001　output→1
	fmt.Println(1 << 1) // 0001 0010　output→2
	fmt.Println(1 << 2) // 0001 0100　output→4
	fmt.Println(1 << 3) // 0001 1000　output→8
}
