package fishServer

type CMD_TYPE int

const (
	CMD_MAIN_GAME    CMD_TYPE = 0
	CMD_MAIN_CONTROL CMD_TYPE = 201
)

type CMD_CONTROL_TYPE int

const (
	CONTROL_CS_CHECK_CLIENT_INFO CMD_CONTROL_TYPE = 1
	CONTROL_SC_CHECK_CLIENT_INFO CMD_CONTROL_TYPE = 2
	CONTROL_CS_KICK_USER_BY_ID   CMD_CONTROL_TYPE = 7
	CONTROL_SC_KICK_USER_BY_ID   CMD_CONTROL_TYPE = 8
)

type SUB_CMD_TYPE byte

const (
	SC_CMD_HEARTBEAT    SUB_CMD_TYPE = 0
	CS_CMD_CHECK_CLIENT SUB_CMD_TYPE = 1
	SC_CMD_CHECK_CLIENT SUB_CMD_TYPE = 2
	//player action
	CS_CMD_FREEZE_PLAYER SUB_CMD_TYPE = 7
	SC_CMD_FREEZE_PLAYER SUB_CMD_TYPE = 8
	//Server Info
	SC_CMD_GAME_INFO    SUB_CMD_TYPE = 21
	SC_CMD_LOGON_INFO   SUB_CMD_TYPE = 22
	SC_CMD_DB_INFO      SUB_CMD_TYPE = 23
	SC_CMD_OPERATE_INFO SUB_CMD_TYPE = 26
	//system
	CS_CMD_AGENT_INFO_CHANGE   SUB_CMD_TYPE = 66
	SC_CMD_AGENT_INFO_CHANGE   SUB_CMD_TYPE = 67
	CS_CMD_CHANNEL_INFO_CHANGE SUB_CMD_TYPE = 68
	SC_CMD_CHANNEL_INFO_CHANGE SUB_CMD_TYPE = 69
	CS_CMD_NOTICE_CHANGE       SUB_CMD_TYPE = 76
	SC_CMD_NOTICE_CHANGE       SUB_CMD_TYPE = 77
	//
	CS__CMD_CUSTOM_SERVICE_MSG SUB_CMD_TYPE = 119 //客服消息新增回复
)
