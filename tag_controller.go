package main

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"

  // ORM
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "net/http"
  "bytes"
  _ "net/url"
  "io/ioutil"
  "fmt"
  "strconv"
  _ "time"
  "strings"
  "regexp"
  _ "reflect"
  "encoding/json"

  "./model"
)

type short struct {
  Kind            string `json:kind`
  Id              string `json:id`
  LongUrl         string `json:longUrl`
}

type resNews struct {
  ResponseData    status `json:response`
  ResponseDetails string `json:responseDetails`
  ResponseStatus  int    `json:responseStatus`
}

type status struct {
  Feed            statusNews
}

type statusNews struct {
  FeedUrl         string `json:feedUrl`
  Title           string `json:title`
  Link            string `json:link`
  Author          string `json:author`
  Description     string `json:description`
  Type            string `json:type`
  Entries         []News
}

type News struct {
  Title          string `json:title`
  Link           string `json:link`
  Author         string `json:author`
  PublishDate    string `json:publishDate`
  ContentSnippet string `json:contentsnippet`
  Content        string `json:content`
  Categories     []string
}

type response struct {
  NewsData       []resNewsData
}

type resNewsData struct {
  Title          string `json:title`
  Link           string `json:link`
  PublishDate    string `json:publishDate`
}

func tagInit(db *gorm.DB, c echo.Context) model.Tag{
  id, _ := strconv.Atoi(c.QueryParam("id"))
  return model.Tag{id, c.QueryParam("title")}
}

func TagRegist(db *gorm.DB) echo.HandlerFunc{
  return func(c echo.Context) error {
    var count = 0
    tag := tagInit(db, c)
    db.First(&tag).Count(&count)
    if count == 0 {
      db.Create(&tag)
    }
    db.First(&tag)
    taskid, _  := strconv.Atoi(c.QueryParam("task_id"))
    task_tag := model.Task_Tag{0, taskid, tag.Id}
    db.Create(&task_tag)
    return c.JSON(http.StatusOK, Res{Status:true})
  }
}

func getShortURL(topic string, query string) *http.Response {
  GOOGLESHORT := "https://www.googleapis.com/urlshortener/v1/url?key=AIzaSyCGHN5KRMcHq5CO3F2LU3BXLbgepCx0HOw"
  NEWSURL := "http://news.google.com/news?hl=ja&ned=us&ie=UTF-8&oe=UTF-8&output=rss"
  if topic != "" {
    etc := "&topic=" + topic
    NEWSURL += etc
  }
  if query != "" {
    etc := "&q=" + query
    NEWSURL += etc
  }

  client := &http.Client{}
  jsonStr := `{"longUrl":"`+NEWSURL+`"}`
  req, _ := http.NewRequest(
    "POST",
    GOOGLESHORT,
    bytes.NewBuffer([]byte(jsonStr)),
  )
  req.Header.Add("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
  return resp
}

func NewsURLShow(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    topic := c.QueryParam("topic")
    query := c.QueryParam("query")
    resp := getShortURL(topic, query)
    defer resp.Body.Close()
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println(err)
    }
    jsonBytes := ([]byte)(string(b))
    data := short{}
    err = json.Unmarshal(jsonBytes,&data)

    URL := "https://ajax.googleapis.com/ajax/services/feed/load?v=1.0&q="+data.Id+"&num=3"
    re, _ := http.NewRequest("GET", URL, nil)
    client := new(http.Client)
    res, _ := client.Do(re)
    defer res.Body.Close()
    byteArray, _ := ioutil.ReadAll(res.Body)
    jsonBytes = ([]byte)(string(byteArray))

    news := resNews{}
    err = json.Unmarshal(jsonBytes,&news)
    response := response{}
    for _, newsdata := range news.ResponseData.Feed.Entries {
      newsData := resNewsData{}
      newsData.Title = newsdata.Title
      r := regexp.MustCompile(`url=.*`)
      link := strings.Split(r.FindAllStringSubmatch(newsdata.Link, -1)[0][0],"=")
      newsData.Link  = link[1]
      newsData.PublishDate = newsdata.PublishDate
      response.NewsData = append(response.NewsData, newsData)
    }
    return c.JSON(http.StatusOK, response)
  }
}
