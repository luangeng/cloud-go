package mysql

import (
	"database/sql"
	"fmt"
)

// DBDriver aa
//const DBDriver = "root:root@tcp(127.0.0.1:3306)/medex?charset=utf8"
//db, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8")

var DBConmap = make(map[string] DBCon)

type DBCon struct {
	name string
	url string
	db *sql.DB
}

func open(name, url string){
	db, err := sql.Open("mysql", url)
	CheckErr(err)
	DBConmap[name] = DBCon{name, url, db}
}

func Get(name string)*sql.DB{
	return DBConmap[name].db
}

func Return(db *sql.DB){
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

