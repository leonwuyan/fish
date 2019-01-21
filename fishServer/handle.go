package fishServer

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var CS_CHAN = make(chan interface{})
var m_GameServer_Conn net.Conn
var m_ServerIp string
var m_clientValue int
var heartBeat = CMD_BASE{65535, 255, 255}
var m_IsConnected = false
var m_IsStop = false
var m_ConnDeadTime = time.Now()
var m_TimeOut = 5

func Start(ip string, port int, checkValue int) {
	m_ServerIp = ip + ":" + strconv.Itoa(port)
	m_clientValue = checkValue
	if Connection() == nil {
		go handlerReceive()
		go handlerSend()
		//发送客户端验证消息
		FishInstance.CheckClient(m_clientValue, [18]byte{0})
		for !m_IsStop {
			time.Sleep(time.Second * 10)
			continue
		}
	} else {
		println("服务器启动失败")
	}
}
func Stop() {
	m_IsStop = true
}
func Disconnection() {
	m_GameServer_Conn.Close()
	m_IsConnected = false
}
func Connection() (err error) {
	m_GameServer_Conn, err = net.Dial("tcp", m_ServerIp)
	if err != nil {
		println("连接游戏服务器失败")
		m_IsConnected = false
	} else {
		println("连接游戏服务器成功")
		m_ConnDeadTime = time.Now().Add(time.Duration(m_TimeOut) * time.Second)
		m_GameServer_Conn.SetDeadline(m_ConnDeadTime)
		m_IsConnected = true
	}
	return
}
func ReConnection() {
	if Connection() == nil {
		FishInstance.CheckClient(m_clientValue, [18]byte{0})
	}
}
func handlerSend() {
	for {
		if m_IsConnected {
			cmd := <-CS_CHAN
			_, err := m_GameServer_Conn.Write(CmdToBytes(cmd))
			if err != nil {
				//println(fmt.Sprintf("发送数据失败，%+v，%s", cmd, err.Error()))
				//发送失败，重发一下
				CS_CHAN <- cmd
			} else {
				//println(fmt.Sprintf("发送成功，%+v", cmd))
			}
		}
	}
}
func handlerReceive() {
	for {
		if m_IsConnected {
			if time.Now().After(m_ConnDeadTime) {
				Disconnection()
				continue
			}
			var buffer = make([]byte, 1024*8)
			_, err := m_GameServer_Conn.Read(buffer)
			if err != nil {
				println(fmt.Sprintf("读取消息失败,%s", err.Error()))
				continue
			}
			var cmd CMD_BASE
			cmd.FromBytes(buffer)
			handlerServerMsg(cmd, buffer)
			m_ConnDeadTime = time.Now().Add(time.Duration(m_TimeOut) * time.Second)
			m_GameServer_Conn.SetDeadline(m_ConnDeadTime)
		} else {
			ReConnection()
		}
	}
	time.Sleep(100)
}
func handlerServerMsg(cmd CMD_BASE, buffer []byte) {
	switch CMD_TYPE(cmd.CMD_TYPE) {
	case CMD_MAIN_GAME:
		switch SUB_CMD_TYPE(cmd.SUB_CMD_TYPE) {
		case SC_CMD_HEARTBEAT: //服务器心跳消息
			onServerHeartBeat()
			break
		}
		break
	case CMD_MAIN_CONTROL:
		switch SUB_CMD_TYPE(cmd.SUB_CMD_TYPE) {
		case SC_CMD_CHECK_CLIENT:
			var data SC_CHECK_CLIENT
			data.FromBytes(buffer)
			onReceiveCheckClient(data)
			break
		case SC_CMD_GAME_INFO:
			var data SC_GAME_INFO
			data.FromBytes(buffer)
			onReceiveGameInfo(data)
			break
		case SC_CMD_LOGON_INFO:
			var data SC_LOGON_INFO
			data.FromBytes(buffer)
			onReceiveLogonInfo(data)
			break
		case SC_CMD_DB_INFO:
			var data SC_DB_INFO
			data.FromBytes(buffer)
			onReceiveDbInfo(data)
			break
		case SC_CMD_OPERATE_INFO:
			var data SC_OPERATE_INFO
			data.FromBytes(buffer)
			onReceiveOperateInfo(data)
			break
		case SC_CMD_FREEZE_PLAYER:
			var data SC_FREEZE_PLAYER
			data.FromBytes(buffer)
			onReceiveFreezePlayer(data)
			break
		case SC_CMD_AGENT_INFO_CHANGE, SC_CMD_CHANNEL_INFO_CHANGE, SC_CMD_NOTICE_CHANGE:
			println(fmt.Sprintf("收到回执消息：%+v,%d", cmd, len(buffer)))
			break
		default:
			println(fmt.Sprintf("收到消息：%+v,%d", cmd, len(buffer)))
			return
		}
		break
	}
}
func onServerHeartBeat() {
	CS_CHAN <- heartBeat
}
func onReceiveCheckClient(data SC_CHECK_CLIENT) {
	//println(fmt.Sprintf("收到消息：%+v", data))
	if data.Result {
		println("验证成功")
	}
}
func onReceiveGameInfo(data SC_GAME_INFO) {
	//println(fmt.Sprintf("收到消息：%+v", data))
}
func onReceiveLogonInfo(data SC_LOGON_INFO) {
	//println(fmt.Sprintf("收到消息：%+v", data))
}
func onReceiveDbInfo(data SC_DB_INFO) {
	//println(fmt.Sprintf("收到消息：%+v", data))
}
func onReceiveOperateInfo(data SC_OPERATE_INFO) {
	//println(fmt.Sprintf("收到消息：%+v", data))
}
func onReceiveFreezePlayer(data SC_FREEZE_PLAYER) {
	//println(fmt.Sprintf("收到消息：%+v", data))
	resultStr := ""
	switch data.Result {
	case 0:
		resultStr = "错误 玩家不存在"
		break
	case 1:
		resultStr = "玩家不在线 冻结数据写入数据库"
		break
	case 2:
		resultStr = "玩家处于AFK 状态 冻结成功"
		break
	case 3:
		resultStr = "玩家处于Exit 状态 冻结成功"
		break
	case 4:
		resultStr = "玩家在线 剔除 冻结成功"
		break
	}
	println(resultStr)
}
func makeBaseCmd(cmdType CMD_TYPE, subType SUB_CMD_TYPE) (base CMD_BASE) {
	base = CMD_BASE{
		CMD_TYPE:     byte(cmdType),
		SUB_CMD_TYPE: byte(subType),
	}
	return
}
