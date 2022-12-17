package play

import (
	"fmt"
	"log"
	"os"
)

func read() {
	// バイト単位での書き込み
	data, err := os.Create("play/play.py")
	defer data.Close()
	if err != nil {
		log.Fatalln(err)
	}
	byte := make([]byte, 1000)
	data.Read(byte)
	fmt.Println(string(byte))
}

func Play() {
	read()
}
