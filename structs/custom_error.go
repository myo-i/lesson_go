package _structs

import "fmt"

type NotFound struct {
	UserName string
}

// カスタムエラーも原理はStringerと同じ
// カスタムエラーを作成する際はポインタとアドレスで実装したほうが良いとされている
func (e *NotFound) Error() string {
	return fmt.Sprintf("Not found: %v", e.UserName)
}

func myFuncError() error {
	// 何か知らのエラー判定を行う
	ok := false
	if ok {
		return nil
	}
	return &NotFound{"bob"}
}

func CustomError() {
	if err := myFuncError(); err != nil {
		fmt.Println(err)
	}

	// カスタムエラーをポインタとアドレスで実装したほうが良い理由
	// 下記のようにe := NotFound{"bob"}として定義したものを比較すると同一のエラーとして処理されてしまう
	//e1 := NotFound{"bob"}
	//e2 := NotFound{"bob"}　// e1 == e2 → true 同一のエラーとして処理される
	e1 := &NotFound{"bob"}
	e2 := &NotFound{"bob"} // e1 == e2 → false 異なるエラーとして処理される
	fmt.Println(e1 == e2)
	if err := myFuncError(); err != nil {
		fmt.Println(err)
		// エラーを比較して処理を分岐させる場合もあるので正確な比較が必要
		if err == e1 {

		}
		if err != e2 {

		}
	}
}

//なぜカスタムエラーを作成した際は値ではなくアドレスで比較するのか？
//
//例えばkey1:"Mike" と key2:"Mike" として、key1のエラーがUserNotFound{"Mike"}、key2のエラーがUserNotFound{"Mike"}となったとする。
//
//どちらも違うユーザーなのでエラー判定をする際、
//
//UserNotFound{"Mike"}(key1) == UserNotFound{"Mike"}(key2)となるようなことがあっても異なるエラーとして判定されなければならないが、値の比較となると同一のエラーとして処理されてしまう。
//
//そのため同一の値であっても異なるものとして判定するために、エラーそのものをアドレス参照している。
