package dto

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"../model"
	)

type Mysql_db struct {
	db *sql.DB
}

func (f *Mysql_db) Mysql_open() {

	Odb, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/liliangbin?charset=utf8")
	if err != nil {

		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")
	f.db = Odb
}

func (f *Mysql_db) Mysql_close() {
	defer f.db.Close()
	fmt.Println("数据库连接成功   ---======>  关闭")
}

func (f *Mysql_db) Mysql_select_Max( sql_data string) model.One {

	rows, err := f.db.Query(sql_data)
	if err != nil {
		panic(nil)
	}
	defer rows.Close()
	var one model.One
	for rows.Next() {

		err = rows.Scan(&one.Id, &one.Text, &one.ImgUrl, &one.Day, &one.Month, &one.Time, &one.FullId)

		fmt.Println(one.Day)
	}

	return one

}

func (f *Mysql_db) Mysql_select_id(id int ) model.One  {

	var one model.One
	err:=f.db.QueryRow("select * from one  where f_id=?",id).Scan(&one.Id, &one.Text, &one.ImgUrl, &one.Day, &one.Month, &one.Time, &one.FullId)

	if err!=nil {
		panic(err)
	}
	fmt.Println(one.Day)
	return one

}

func (f *Mysql_db) Insert() {

	result, err := f.db.Exec(
		"INSERT INTO one (text,time) VALUES (?, ?)",
		"gopher",
		"dfdsfs",
	)

	if err != nil {

		fmt.Println("dfff")
	}

	fmt.Println(result)

}
