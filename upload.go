package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)


func main() {
	db, err := sql.Open("mysql","dev:Bandu12#$@tcp(172.16.1.98:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接出错了")
	}
	defer db.Close()
	rows,_ := db.Query("select id,name from s_user")
	for rows.Next(){
		var id int
		var name string
		err = rows.Scan(&id,&name)
		if err != nil{
			fmt.Println("取数失败了")
		}
		fmt.Println(id)
		fmt.Println(name)
	}

	xlsx, err := excelize.OpenFile("./test1/hxb.xlsx")
	if err != nil{
		fmt.Println(err)
		return
	}
	excelRows := xlsx.GetRows("试题汇总")
	excelRows = excelRows[1:]
	for _, row := range excelRows{
		for _, coCell := range row{
			fmt.Println(coCell,"\t")
		}
		fmt.Println()
	}
}
