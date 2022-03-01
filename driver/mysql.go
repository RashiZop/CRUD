package driver

import (
   "database/sql"
   "fmt"
)

func ConnectToSQL() *sql.DB {
   db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
   if err != nil {
      panic(err.Error())
   } else {
      fmt.Println("Connected")
   }
   return db
}