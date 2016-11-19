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
)

//ユーザーの構造体
type user_data struct {
	id    string
	year  string
	month string
}

type json_comp struct {
  Status  bool  `json:status`
  Data    []json_day
}

type json_day struct {
  Day   []json_event
}

//クエリから情報取得
//ユーザー情報の初期化
func user_initation(c echo.Context) user_data {
	return user_data{c.QueryParam("id"), c.QueryParam("year"), c.QueryParam("month")}
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
	schev := json_all{}
  sche  := json_comp{}

	//returnするjson
  sche.Status = true
	//充分なデータを取得できていなかったらstatus:falseでreturn
	if data[0] == "false"{
    //fal := json_event{0,"0","0","0","0"}
    res := json_comp{}
    gg := json_day{}
    res.Status = false
    res.Data = append(res.Data,gg)
    //res.Data[0] = append(res.Data[0],fal)
		return res
	}

  day := json_day{}
  for k := 0; k < 31;k = k + 1 {
    sche.Data = append(sche.Data,day)
  }
  for i := 0; i < len(data); i = i + num_colmu {
    id , _ := strconv.Atoi(data[0+i])
		code := json_event{id,data[1+i],data[2+i],data[3+i],data[4+i]}
    d , _ := strconv.Atoi(data[5+i])
    fmt.Println(d)
    fmt.Println(sche)
    fmt.Println(sche.Data[d])
    sche.Data[d-1].Day = append(sche.Data[d-1].Day,code)
    schev.Data = append(schev.Data,code)
	}
  fmt.Println(sche.Data[0].Day)
	return sche
}

func Echo_event(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user_initation(c)  //ユーザー情報を取得
		json := user.get_event(db) //イベント情報を取得
		return c.JSON(http.StatusOK, json)
	}
}
