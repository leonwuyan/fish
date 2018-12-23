package models

type Result struct {
	State int         `json:"state"`
	Msg   string      `json:"msg,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Total int         `json:"total"`
}
