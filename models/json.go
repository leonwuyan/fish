package models

type Result struct {
	State int         `json:"state"`
	Msg   string      `json:"msg,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Total int         `json:"total"`
}

type StatisticsPlay struct {
	GameId      int `orm:"column(GameId)" json:"game_id"`
	PlayTimes   int `orm:"column(PlayTimes)" json:"play_times"`
	PlayPlayers int `orm:"column(PlayPlayers)" json:"play_players"`
	WinOrLose   int `orm:"column(WinOrLose)" json:"win_or_lose"`
}

type ClientMessage struct {
	MsgType    int    `json:"msg_type"`
	MsgContent string `json:"msg_content"`
}
