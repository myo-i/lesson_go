package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ここら辺知っとけば十分
func loggingSettings(logFile string) {
	// フラグと読み書きできるグループを設定
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// Stdoutは画面上に出るエラーでlogfileに書き込む
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	// SetFlagsで日時、時間、ファイル名（long ver）のフラグを設定
	// SetOutputで実際に書き込むエラーと書き込む場所を設定？
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func Logging() {
	loggingSettings("test.log")

	_, err := os.Open("asdfg")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	// Fatalは実行された時点でプログラムが終了する
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!!!")
	fmt.Println("ok")
}
