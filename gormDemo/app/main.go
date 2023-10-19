package main

import (
	"PracticeProjects/gormDemo/mysql/dataconv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("starting..")
	fmt.Println("starting")

	dsn := "root:Mylife1586@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 查询
	var expenseList []dataconv.Expense
	err = db.Find(&expenseList).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(expenseList)
}
