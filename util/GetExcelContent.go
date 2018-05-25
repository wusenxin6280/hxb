package util

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func GetExcelContent(file ,sheet string)  ([]string,error){
	xlsx, err := excelize.OpenFile(file);
	if err != nil{
		return nil,fmt.Errorf("Excel文件不存在！")
	}
	rows := xlsx.GetRows(sheet)
	rows = rows[1:]
	unit := make(map[string]interface{})
	for _, row := range rows{
		unit["id"] = row[0]
		unit["name"] = row[1]
	}

}
