package _switch

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "asdg hs"
}

func Switch() {
	// switch文のみの変数を使う場合は一行で書ける
	switch os := getOsName(); os {
	case "windows":
		fmt.Println("Windows!!")
	case "mac":
		fmt.Println("Mac!!")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")

	}
}
