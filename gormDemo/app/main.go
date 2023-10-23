package main

import (
	"PracticeProjects/gormDemo/mysql/dataconv/model/lhmis"
	"encoding/json"
	"fmt"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func openMySQL() (db *gorm.DB, err error) {
	dsn := "root:Mylife1586@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	return db, err
}

func openOracle() (db *gorm.DB, err error) {
	url := oracle.BuildUrl("192.168.1.110", 1521, "database", "oa", "oa", nil)
	db, err = gorm.Open(oracle.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	return db, err
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	fmt.Println("starting..")

	db, err := openOracle()
	s2, err := db.DB()
	defer s2.Close()
	if err != nil {
		log.Fatal(err)
	}
	//var resultList []lhmis.YwRequisition
	//date, _ := time.Parse("2006-01-02", "2023-01-01")
	//db.Where("CREATEDATE > ?", date).Find(&resultList)
	//db.Where("R_ID = ?", "FY20231017002").Find(&resultList)
	//
	//fmt.Println(len(resultList))
	//
	//for _, requisition := range resultList {
	//	fmt.Println(requisition.R_ID)
	//	marshal, _ := json.Marshal(requisition)
	//	fmt.Println(marshal)
	//
	//}

	var request lhmis.YwRequisition
	request.R_ID = "F230101"
	request.CREATEDATE = time.Now()

	marshal, _ := json.Marshal(request)
	fmt.Printf("%s\n", marshal)

	var movie Movie
	movie.Title = "哪吒脑海"
	json, _ := json.Marshal(movie)
	fmt.Printf("%s\n", json)

}
