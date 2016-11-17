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
	_ "strconv"
)

//ユーザーの構造体
type user_data struct {
	id    string
	year  string
	month string
}

//イベントの構造体(json形式)
type event struct {
	Id          string `json:Id`
	Summary     string `json:"Summary"`
	Dtstart     string `json:"dtstart"`
	Dtend       string `json:"dtend"`
	Description string `json:"dtstart"`
	Day         string `json:day`
}

type json_event struct {
	Id          string `json:id`
	Summary     string `json:"summary"`
	Dtstart     string `json:"dtstart"`
	Dtend       string `json:"dtend"`
	Description string `json:"description"`
}

type json_all struct {
  Status  bool `json:status`
  Data   []json_event
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
			//データが取得できているかのエラー県ん出
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
func (user user_data) get_event(db *sql.DB) json_all {
	//イベントデータを連想配列で取得
	data := user.extract_eventdata_from_db(db)
	//取得したカラム数
	num_colmu := 6

	//各日のイベント格納用連想配列の初期化
	schev := json_all{}
  sche := make([]string,0)
	var st string
	//var index int

	//returnするjson
  json := "{'status':'true','data':"
  schev.Status = true
  json_map := map[string]string{}
  json_map["status"] = "true"
	//json := "{'status':'true','data':{"
	//充分なデータを取得できていなかったらstatus:falseでreturn
	if data[0] == "false"{
    fal := json_event{"0","0","0","0","0"}
    res := json_all{}
    res.Status = true
    res.Data = append(res.Data,fal)
		return res
	}
  
	for i := 0; i < len(data); i = i + num_colmu {
		st = "['id':" + data[0+i] + ",'Summary':'" + data[1+i] + "','dtstart':'" + data[2+i] + "','dtend':'" + data[3+i] + "','description':'" + data[4+i] + "']"
    code := json_event{data[0+i],data[1+i],data[2+i],data[3+i],data[4+i]}

		//index, _ = strconv.Atoi(data[5+i])
		sche = append(sche, st)
    schev.Data = append(schev.Data,code)
	}

		//json += strconv.Itoa(set+1) + ":"
	for cont := 0; cont < len(sche); cont += 1 {
		json += sche[cont]
		if len(sche)-1 != cont {
			json += ","
		}
	}
	json += "}"
	return schev
}

func Echo_event(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user_initation(c)  //ユーザー情報を取得
		json := user.get_event(db) //イベント情報を取得
		return c.JSON(http.StatusOK, json)
	}
}
