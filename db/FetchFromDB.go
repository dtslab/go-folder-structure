package db

import "fmt"

// these can be called inside services
func FetchFromDB() {
	fmt.Println("QUery DB to fetch data")
}