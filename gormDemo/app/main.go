package main

import (
	"PracticeProjects/gormDemo/mysql/dataconv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type YwUser struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	fmt.Println("starting..")

	dsn := "root:Mylife1586@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}

	//// 查询
	//var resultList []YwUser
	//err = db.Find(&resultList).Error
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resultList)

	var resultList []dataconv.Expense
	err = db.Find(&resultList).Error
	if err != nil {
		panic(err)
	}
	for _, expense := range resultList {
		fmt.Println(expense.ExpenseId)
	}
}
