package model

type json_event struct {
	Id          int `json:id`
	Summary     string `json:"summary"`
	Dtstart     string `json:"dtstart"`
	Dtend       string `json:"dtend"`
	Description string `json:"description"`
}

type json_all struct {
  Status  bool `json:status`
  Data    []json_event
}

type json struct {
  Status  bool `json:status`
  Data    []json_task
}

type json_task struct {
	Id        int    `json:id`
	Title     string `json:"summary"`
	Sub_task  string `json:"sub_task"`
  year      int    `json:"year"`
  month     int    `json:"month"`
  day       int    `json:"day"`
}
