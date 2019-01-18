package fishServer

var FishInstance = newFishServer()

type FishServer struct {
}

func newFishServer() *FishServer {
	return new(FishServer)
}

func (this *FishServer) CheckClient(clientValue int, machineCode [18]byte) {
	CS_CHAN <- CS_CHECK_CLIENT{
		makeBaseCmd(CMD_MAIN_CONTROL, CS_CMD_CHECK_CLIENT),
		uint32(clientValue),
		machineCode,
	}
}
func (this *FishServer) FreezePlayerById(id int, min int) {
	CS_CHAN <- CS_FREEZE_PLAYER{
		makeBaseCmd(CMD_MAIN_CONTROL, CS_CMD_FREEZE_PLAYER),
		uint32(id),
		uint32(0),
		uint32(min),
	}
}
func (this *FishServer) ChangeAgentConfig(id int, isAdd bool) {
	CS_CHAN <- CS_AGENT_INFO_CHANGE{
		makeBaseCmd(CMD_MAIN_CONTROL, CS_CMD_AGENT_INFO_CHANGE),
		uint32(id),
		isAdd,
	}
}
func (this *FishServer) ChangeChannelConfig(id int, isAdd bool) {
	CS_CHAN <- CS_CHANNEL_INFO_CHANGE{
		makeBaseCmd(CMD_MAIN_CONTROL, CS_CMD_CHANNEL_INFO_CHANGE),
		uint32(id),
		isAdd,
	}
}
func (this *FishServer) ChangeNoticeConfig() {
	CS_CHAN <- CS_NOTICE_CHANGE{
		makeBaseCmd(CMD_MAIN_CONTROL, CS_CMD_NOTICE_CHANGE),
	}
}
func (this *FishServer) SendServiceMsg(userId, msgId int) {
	CS_CHAN <- CS_CUSTOM_SERVICE_MSG{
		makeBaseCmd(CMD_MAIN_CONTROL, CS__CMD_CUSTOM_SERVICE_MSG),
		uint32(userId),
		uint32(msgId),
	}
}
