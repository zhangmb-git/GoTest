package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ExcuteSql(strSql string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	result, err := db.Exec(strSql) //"INSERT INTO user_info(username,sex,email)VALUES (?,?,?)",
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(result)
	return
}

func InsertAskLog(userID int64, question string, answer string, flag int) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	strSql := "INSERT INTO t_ask(userID,question,answer,flag)VALUES (?,?,?,?)"
	_, err = db.Exec(strSql, userID, question, answer, flag)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.Close()
	//fmt.Println(result)
	return
}

func InsertAskQuestions(userID int64, question string, answer string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	strSql := "INSERT INTO t_question(userID,question,answer)VALUES (?,?,?)"
	_, err = db.Exec(strSql, userID, question, answer)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.Close()
	//fmt.Println(result)
	return
}
