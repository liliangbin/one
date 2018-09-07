package main

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"bytes"
	"database/sql"
	"strconv"
)

type IString interface {
	SubString(index int, end int) string
	IndexOf(index string) int
}

type Stringer struct {
	Str string
}

type Total struct {
	OneInfo []OneInfo
}

type OneInfo struct {
	Text   []string `json:"text"`
	ImgUrl []string `json:"imgUrl"`
	Day    []string `json:"day"`
	Month  []string `json:"month"`
	Find   string     `json:"find"`
}

func init() {
	fmt.Println("开始准备录入数据,都是测试数据")
	fmt.Println(reflect.TypeOf("."))
	str := Stringer{"ex.index"}
	index := str.IndexOf(".")
	fmt.Println(index)
	strstr := str.SubString(0, index)
	fmt.Println(strstr)
	var buffer bytes.Buffer
	buffer.WriteString("hello world ")
	buffer.WriteString("heiheihhie ")
	stringgg := buffer.String()
	fmt.Println(stringgg)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("正式数据")
}

func (str Stringer) SubString(index int, end int) string {

	if index > end {
		panic(nil)
	}
	str1 := str.Str[index:end]
	return str1

}
func (str Stringer) IndexOf(index string) int {

	for head, str1 := range str.Str {

		if string(str1) == index {
			return head
		}
	}

	return -1
}

var m2 = map[string]int{
	"Jan": 1,
	"Feb": 2,
	"Mar": 3,
	"Apr": 4,
	"May": 5,
	"Jun": 6,
	"Jul": 7,
	"Aug": 8,
	"Sep": 9,
	"Oct": 10,
	"Nov": 11,
	"Dec": 12,
}

func main() {
	var err error
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/liliangbin?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	data, err := ioutil.ReadFile("../articleexport.json")
	if err != nil {
		log.Println(err)
	}
	var oneinfos []OneInfo
	errr := json.Unmarshal(data, &oneinfos)
	if errr != nil {
		log.Println(errr)
	}

	db.Exec("TRUNCATE `liliangbin`.`one`")
	for _, one := range oneinfos {
		insert(db, one)
	}
}

func insert(db *sql.DB, info OneInfo) {

	month_year := info.Month[0]

	month_year_O := Stringer{month_year}
	temp := month_year_O.SubString(0, month_year_O.IndexOf("."))
	mon := strconv.Itoa(m2[temp])
	fmt.Println(temp)
	year := month_year_O.SubString(month_year_O.IndexOf(".")+2, len(month_year_O.Str))

	var buffer bytes.Buffer
	buffer.WriteString(year)
	buffer.WriteString("-")
	buffer.WriteString(mon)
	buffer.WriteString("-")
	buffer.WriteString(info.Day[0])
	fmt.Println(buffer.String())
	fmt.Println(info.Find)
	db.Exec("insert into one(text,img_url,day,month,time,f_id )values (?,?,?,?,?,?)", info.Text[0], info.ImgUrl[0], info.Day[0], info.Month[0], buffer.String(),info.Find)
}
