
package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    dbUser := os.Getenv("MYSQL_USER")
    dbPass := os.Getenv("MYSQL_PASSWORD")
    dbHost := os.Getenv("MYSQL_HOST")
    dbName := os.Getenv("MYSQL_DB")
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic("データベースへの接続に失敗しました")
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Docker!")
    })

    http.ListenAndServe(":8080", nil)
}
