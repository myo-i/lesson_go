package play

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var filedata []string

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

// os.Openで読み込むとline := linesScan.Text()で一行読めたけど、os.Createだと読み込めない...
// それはOpenはファイルを読み込むだけなのに対して、Createはあくまで作成。
// ないものを読み込めば新規作成してくれるし、あるものを読み込めば中身を空にする
// Openで開いたファイルは書き込めず、Createで開いたファイルは書き込める
func readWriteFIle(filepath string) {
	// Openは読み込み権限のみ
	file, err := os.Open(filepath)
	// Createは書き込み権限も追加
	if err != nil {
		fmt.Println(err, "\n", "Error can't read file")
	}
	defer file.Close()

	linesScan := bufio.NewScanner(file)
	for linesScan.Scan() {
		line := linesScan.Text()
		if strings.Contains(line, "python") {
			line = strings.Replace(line, "python", "Golang", 1)
		}
		fmt.Println(line)
		filedata = append(filedata, line)
	}

	// いったんスライスにデータを保持して書き込んでみる
	createfile, _ := os.Create("demo.py")
	// 読み込む対象のファイルに書き込む場合は下記を使う
	//createfile, _ := os.Create(filepath)
	for _, data := range filedata {
		_, err := createfile.WriteString(data + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err := linesScan.Err(); err != nil {
		fmt.Println(err)
	}
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
	readWriteFIle("play/play.py")
	//readByte("play/play.py")
	//readRecord("play/play.py")
	fmt.Println(filedata)
}
