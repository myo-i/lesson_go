package panic_and_recover

import "fmt"

// 基本的にはpanicではなく、エラーハンドリングを行うことが推奨されている！！

func thirdPartyConnectDB() {
	// panicは呼び出されるとエラーのスタックトレースを吐き、コードを強制終了する
	panic("Unable to connect database!!")
}

func save() {
	defer func() {
		// recoverでpanicを起こした後の処理を記述できる
		s := recover()
		fmt.Println(s)
	}()
	// deferの上に記述するとpanicのみ発生する
	thirdPartyConnectDB()
}

func PanicAneRecover() {
	save()
	fmt.Println("OK?")
}
