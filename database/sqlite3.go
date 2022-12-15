package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func Database() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
             	name STRING,
				age  INT)`
	// 返り値はいらないのでExecを使い、"_"で受け取る
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	//cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// 結果がいらない場合Exec
	//_, err = DbConnection.Exec(cmd, "Nancy", 20)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//cmd = "UPDATE person SET age = ? WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, 26, "Mike")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//マルチセレクト
	//cmd = "SELECT * FROM person"
	//rows, _ := DbConnection.Query(cmd)
	//defer DbConnection.Close()
	//var pp []Person
	//// パターンとして覚えておく
	//for rows.Next() {
	//	var p Person
	//	err := rows.Scan(&p.Name, &p.Age)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	pp = append(pp, p)
	//}
	//// 上記は1個1個エラーを見ているが、まとめて書くこともできる
	//err = rows.Err()
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//for _, p := range pp {
	//	fmt.Println(p.Name, p.Age)
	//}

	//// シングルセレクト
	//cmd = "SELECT * FROM person WHERE age = ?"
	//row := DbConnection.QueryRow(cmd, 20)
	//var p Person
	//err = row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row")
	//	} else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	// 削除
	//cmd = "DELETE FROM person WHERE name = ?"
	//// 結果がいらない場合Exec
	//_, err = DbConnection.Exec(cmd, "Nancy")
	//if err != nil {
	//	log.Println(err)
	//}

	// SQLインジェクションの例（ダイナミックな値に対して?を使うのは余計なSQL文を取り除いてくれるから）
	tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.XXX', 100)"
	// tableの場合は"?"でダイナミックに代入はできない
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer DbConnection.Close()
	var pp []Person
	// パターンとして覚えておく
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	// 上記は1個1個エラーを見ているが、まとめて書くこともできる
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

}
