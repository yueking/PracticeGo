package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
)

func main() {
	fmt.Println("starting...")
	var driverName = "oci8"
	var dataSourceName = "oa/oa@192.168.1.110:1521/database"

	var str = "select R_ID from YW_REQUISITION"
	db, err := sql.Open(driverName, dataSourceName)
	defer db.Close()
	//
	if err != nil {
		log.Fatal(err)
	}
	println("Connection success !")
	rows, err := db.Query(str)

	if err != nil {
		log.Fatal(err)
	}
	var (
		R_ID string
	)
	for rows.Next() {
		if err = rows.Scan(&R_ID); err != nil {
			log.Fatalln("error fetching", err)
		}
		log.Println(R_ID)
	}

}
