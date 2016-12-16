package main

import (
  "fmt"
  "time"
  "reflect"
)

func main() {
  dat := time.Now()
  const layout = "2016-1-1"
  fmt.Println(reflect.TypeOf(dat.Format(layout)))
}
