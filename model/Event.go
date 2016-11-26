package model

import (
	//apiパッケージ
	"github.com/labstack/echo"
	"net/http"

	//mysqlパッケージ
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	//標準パッケージ
	"fmt"
	"strconv"
  "strings"
)

//ユーザーの構造体
type user_data struct {
	id    string
	year  string
	month string
}

//イベント情報の構造体
type event_data struct {
  id string
  user_id string
  summary string
  dtstart string
  dtend string
  description string
  year string
  month string
  day string
}

//クエリから情報取得
//ユーザー情報の初期化
func user_init(c echo.Context) user_data {
	return user_data{c.QueryParam("id"), c.QueryParam("year"), c.QueryParam("month")}
}

//クエリからの取得情報でのイベント情報初期化
func event_init(c echo.Context) event_data{
  return event_data{c.QueryParam("id"),c.QueryParam("user_id"),c.QueryParam("summary"),c.QueryParam("dtstart"),c.QueryParam("dtend"),c.QueryParam("description"),c.QueryParam("year"),c.QueryParam("month"),c.QueryParam("day")}
}

//日付データをyear,month,dayにパース
func  parse_timedata(event event_data) event_data{
  fmt.Print(event)
  dtstart := event.dtstart

  //時間と日にちに分解
  date_time := strings.Split(dtstart," ")
  date := strings.Split(date_time[0],"-")

  event.year = date[0]
  event.month = date[1]
  event.day = date[2]

  return event
}

//データベースからユーザーの登録したイベント情報を抽出
func (user user_data) extract_eventdata_from_db(db *sql.DB) []string {
	//sqlクエリ
	query := "select id,summary,dtstart,dtend,description,day from Event where user_id=" + user.id + " and year=" + user.year + " and month=" + user.month + " order by day"
	//データ取得
	rows, err := db.Query(query)

	var data_extracted_from_db []string
	//データが抽出できているかのエラー検出
	if err != nil {
		data_extracted_from_db = append(data_extracted_from_db, "false")
		fmt.Println(err)
		return data_extracted_from_db
	}
	colum, err := rows.Columns()

	values := make([]sql.RawBytes, len(colum))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for _, col := range values {
			//データが取得できているかのエラー検出
			if col == nil {
				data_extracted_from_db = append(data_extracted_from_db, "false")
			} else {
				data_extracted_from_db = append(data_extracted_from_db, string(col))
			}
		}
	}
	return data_extracted_from_db
}

//ユーザーのイベント情報を返す
func (user user_data) get_event(db *sql.DB) json_comp {
	//イベントデータを連想配列で取得
	data := user.extract_eventdata_from_db(db)
	//取得したカラム数
	num_colmu := 6

	//返すjsonデータの初期化
  sche  := json_comp{}

	//returnするjson
  sche.Status = true
	//充分なデータを取得できていなかったらstatus:falseでreturn
	if data[0] == "false"{
    status_response(false)
	}

  day := json_day{}
  for k := 0; k < 31;k = k + 1 {
    sche.Data = append(sche.Data,day)
  }
  for i := 0; i < len(data); i = i + num_colmu {
    id , _ := strconv.Atoi(data[0+i])
		code := json_event{id,data[1+i],data[2+i],data[3+i],data[4+i]}

    d , _ := strconv.Atoi(data[5+i])
    sche.Data[d-1].Day = append(sche.Data[d-1].Day,code)
	}
	return sche
}

func status_response (status bool) json_comp {
    fal := json_event{0,"0","0","0","0"}
    res := json_comp{}
    gg := json_day{}
    res.Status = status
    res.Data = append(res.Data,gg)
    res.Data[0].Day = append(res.Data[0].Day,fal)
    return res
}

//取得データからクエリを生成
//DBに投げてtrue,falseを返す
func (event event_data) regist_event_to_db(db *sql.DB) json_comp{
  full_event := parse_timedata(event)

  query := "insert into Event (user_id,summary,dtstart,dtend,description,year,month,day) values ('"+event.user_id+"','"+event.summary+"','"+event.dtstart+"','"+event.dtend+"','"+event.description+"','"+full_event.year+"','"+full_event.month+"','"+full_event.day+"')"

  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    return status_response(false)
  }
  return status_response(true)
}


//dbのイベント情報の更新
func (event event_data) update_event_to_db(db *sql.DB) json_comp{
  event = parse_timedata(event)

  query := "update Event set summary='"+event.summary+"',dtstart='"+event.dtstart+"',dtend='"+event.dtend+"',description='"+event.description+"',year='"+event.year+"',month='"+event.month+"',day='"+event.day+"' where id="+event.id+" and user_id="+event.user_id

  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    return status_response(false)
  }
  return status_response(true)
}

func (event event_data) delete_event_at_db(db *sql.DB) json_comp {
  query := "DELETE FROM Event WHERE id="+event.id+" and user_id="+event.user_id

  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    return status_response(false)
  }
  return status_response(true)
}

func Echo_event_detail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user_init(c)
		json := user.get_event(db) 
		return c.JSON(http.StatusOK, json)
	}
}

func Echo_event_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := event_init(c)
    status := event.regist_event_to_db(db)
    return c.JSON(http.StatusOK,status)
  }
}

func Echo_event_update(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := event_init(c)
    status := event.update_event_to_db(db)
    return c.JSON(http.StatusOK,status)
  }
}

func Echo_event_delete(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := event_init(c)
    status := event.delete_event_at_db(db)
    return c.JSON(http.StatusOK,status)
  }
}
