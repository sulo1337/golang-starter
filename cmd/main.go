package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:changeme@tcp(127.0.0.1:3306)/cleanarch"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}