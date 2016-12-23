package model

type Tag struct {
  Id    int    `json:id sql:AUTO_INCREMENT`
  Title string `json:title`
}
