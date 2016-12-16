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
type user_ID struct {
	id    string
}

//タスク情報の構造体
type task_data struct {
  id        string
  user_id   string
  title     string
  sub_task  string
  year      string
  month     string
  day       string
}

//クエリから情報取得
//ユーザー情報の初期化
func userID_init(c echo.Context) user_ID {
	return user_ID{c.QueryParam("id")}
}

//クエリからの取得情報でのイベント情報初期化
func task_init(c echo.Context) task_data{
  return task_data{c.QueryParam("id"),c.QueryParam("user_id"),c.QueryParam("title"),c.QueryParam("sub_task"),c.QueryParam("year"),c.QueryParam("month"),c.QueryParam("day")}
}

func (task task_data) regist_task(db *sql.DB) json {
  query := "insert into Task (user_id,title,sub_task,year,month,day) values ('"+task.user_id+"','"+task.title+"','"+task.sub_task+"','"+task.year+"','"+task.month+"','"+task.day+")"
  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    fal := json_task{0,"0","0",0,0,0}
    res := json{}
    res.Status = false
    res.Data = append(res.Data,fal)
    return res
  }
  fal := json_task{0,"0","0",0,0,0}
  res := json{}
  res.Status = true
  res.Data = append(res.Data,fal)
  return res
}

//データベースからユーザーの登録したイベント情報を抽出
func (user user_data) extract_taskdata_from_db(db *sql.DB) []string {
	//sqlクエリ
	query := "select id,title,sub_task,year,mont,day from Task where user_id=" + user.id + "and year=" + user.year + "and month=" + user.month+ "order by dtend"
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
func (user user_data) get_task(db *sql.DB) json {
	//イベントデータを連想配列で取得
	data := user.extract_taskdata_from_db(db)
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

func (task task_data) regist_task_to_db(db *sql.DB) json {
  query := "insert into Task (user_id,title,sub_task,year,month,day) values ('"+task.user_id+"','"+task.title+"','"+task.sub_task+"','"+task.year+"','"+task.month+"','"+task.day+")"
  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    fal := json_task{0,"0","0",0,0,0}
    res := json{}
    res.Status = false
    res.Data = append(res.Data,fal)
    return res
  }
  fal := json_task{0,"0","0",0,0,0}
  res := json{}
  res.Status = true
  res.Data = append(res.Data,fal)
  return res
}

func Echo_task(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user_init(c)  //ユーザー情報を取得
		json_res := user.get_task(db) //イベント情報を取得
		return c.JSON(http.StatusOK, json_res)
	}
}

func Echo_task_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    task := task_init(c)
    status := task.regist_task_to_db(db)
    return c.JSON(http.StatusOK,status)
  }
}
