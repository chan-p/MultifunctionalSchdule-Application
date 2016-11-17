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
type user_id struct {
	id    string
}



//クエリから情報取得
//ユーザー情報の初期化
func user_init(c echo.Context) user_id {
	return user_id{c.QueryParam("id")}
}

//データベースからユーザーの登録したイベント情報を抽出
func (user user_id) extract_eventdata_from_db(db *sql.DB) []string {
	//sqlクエリ
	query := "select id,title,sub_task,year,mont,day from Task where user_id=" + user.id + " order by dtend"
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
func (user user_id) get_task(db *sql.DB) json {
	//イベントデータを連想配列で取得
	data := user.extract_eventdata_from_db(db)
	//取得したカラム数
	num_colmu := 6

	//各日のイベント格納用連想配列の初期化
	task := json{}

	//returnするjson
  task.Status = true
	//充分なデータを取得できていなかったらstatus:falseでreturn
	if data[0] == "false"{
    fal := json_task{0,"0","0",0,0,0}
    task := json{}
    task.Status = false
    task.Data = append(task.Data,fal)
		return task
	}

	for i := 0; i < len(data); i = i + num_colmu {
    id    , _ := strconv.Atoi(data[0+i])
    year  , _ := strconv.Atoi(data[3+i])
    month , _ := strconv.Atoi(data[4+i])
    day  , _  := strconv.Atoi(data[5+i])
		code := json_task{id,data[1+i],data[2+i],year,month,day}
    task.Data = append(task.Data,code)
	}

	return task
}

func Echo_task(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user_init(c)  //ユーザー情報を取得
		json_res := user.get_task(db) //イベント情報を取得
		return c.JSON(http.StatusOK, json_res)
	}
}
