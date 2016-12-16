package model

type User struct {
  Id    string `json:id`
  Email string `json:email`
  Pass  string `json:pass`
}
