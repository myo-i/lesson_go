package play

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readByte(filepath string) {
	// バイト単位での書き込み
	data, err := os.Create(filepath)
	defer data.Close()
	if err != nil {
		log.Fatalln(err)
	}
	byte := make([]byte, 1000)
	data.Read(byte)
	fmt.Println(string(byte))
}

func readRecord(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err, "\n", "Error can't read file")
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		line := lines.Text()
		fmt.Println(line)
	}
	if err := lines.Err(); err != nil {
		return err
	}
	return nil
}

func Play() {
	//readByte("play/play.py")
	readRecord("play/play.py")
}
