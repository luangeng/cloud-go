package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect(name, url string){
	open(name, url)

}

func databases(){

}

func tables(name string) [][]string{
	db := Get(name)
	rows, err := db.Query("show tables")
	defer rows.Close()
	defer Return(db)
	CheckErr(err)
	var data = make([][]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		CheckErr(err)
	}
	return data
}

func insert(name,sql string)string{
	db := Get(name)
	defer Return(db)
	stmt, err := db.Prepare(sql)
	CheckErr(err)
	res, err := stmt.Exec()
	CheckErr(err)
	fmt.Printf("%v", res)
	return "res"
}

func query(name, sql string) [][]string{
	db := Get(name)
	rows, err := db.Query("SELECT * FROM "+ sql)
	defer rows.Close()
	defer Return(db)
	CheckErr(err)
	var data = make([][]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		CheckErr(err)
	}
	return data
}

func delete(sql string){

}


