package main

import (
  "database/sql"
  "fmt"
  "reflect"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:tomohi6@tcp(192.168.60.55:3306)/social-app")
  fmt.Println(reflect.TypeOf(db))
  if err != nil {
    panic(err.Error())
  }
  defer db.Close() // 関数がリターンする直前に呼び出される
  rows, err := db.Query("SELECT * FROM test")
  fmt.Println(rows)
  if err != nil {
    panic(err.Error())
  }
  columns, err := rows.Columns() // カラム名を取得
  if err != nil {
    panic(err.Error())
  }
  values := make([]sql.RawBytes, len(columns))
  fmt.Println(values)
  //  rows.Scan は引数に `[]interface{}`が必要
  scanArgs := make([]interface{}, len(values))
  fmt.Println(scanArgs)
  for i := range values {
    scanArgs[i] = &values[i]
  }
  fmt.Println(values)
  fmt.Println(scanArgs)
  for rows.Next() {
    fmt.Println(values)
    err = rows.Scan(scanArgs...)
    fmt.Println(err)
    fmt.Println(values)
    if err != nil {
      panic(err.Error())
    }
    var value string
    fmt.Println(values)
    for i, col := range values {
      // Here we can check if the value is nil (NULL value)
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
  }
}
