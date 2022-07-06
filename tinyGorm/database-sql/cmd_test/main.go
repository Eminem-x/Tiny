package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"tinyGorm"
)

func main() {
	engine, _ := tinyGorm.NewEngine("sqlite3", "tinyGorm.db")
	defer engine.Close()
	s := engine.NewSession()

	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec() // 为了显示 error, create 两次

	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)

}
