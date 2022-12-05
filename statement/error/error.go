package error

import (
	"fmt"
	"log"
	"os"
)

func Error() {
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, "\n", string(data))

	// 上記のcount, err := ではイニシャライズしているが、今回は=で上書きしていて意味合いが違うので注意
	// 返り値が1つの場合は短縮して書ける
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
