package fishServer

import (
	"bytes"
	"encoding/binary"
)

type CMD_BASE struct {
	CMD_SIZE     uint16
	SUB_CMD_TYPE byte
	CMD_TYPE     byte
}

func (this *CMD_BASE) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

//Check Client Info
type CS_CHECK_CLIENT struct {
	CMD_BASE
	ClientValue uint32
	MachineCode [18]byte
}
type SC_CHECK_CLIENT struct {
	CMD_BASE
	Result bool
}

func (this *SC_CHECK_CLIENT) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

//Server Info
type SC_GAME_INFO struct {
	CMD_BASE
	Id           byte
	Online       uint32
	Tables       uint32
	LogonPlayers uint32
}

func (this *SC_GAME_INFO) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type SC_LOGON_INFO struct {
	CMD_BASE
	Id      byte
	Players uint32
}

func (this *SC_LOGON_INFO) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type SC_DB_INFO struct {
	CMD_BASE
	Id byte
}

func (this *SC_DB_INFO) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type SC_OPERATE_INFO struct {
	CMD_BASE
	Id uint32
}

func (this *SC_OPERATE_INFO) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

//freeze player
type CS_FREEZE_PLAYER struct {
	CMD_BASE
	PlayerId  uint32
	ClientId  uint32
	FreezeMin uint32
}
type SC_FREEZE_PLAYER struct {
	CMD_BASE
	PlayerId uint32
	ClientId uint32
	Result   byte
}

func (this *SC_FREEZE_PLAYER) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type CS_AGENT_INFO_CHANGE struct {
	CMD_BASE
	ChannelID uint32
	IsAdded   bool
}
type SC_AGENT_INFO_CHANGE struct {
	CMD_BASE
	Result bool
}

func (this *SC_AGENT_INFO_CHANGE) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type CS_CHANNEL_INFO_CHANGE struct {
	CMD_BASE
	ChannelID uint32
	IsAdded   bool
}
type SC_CHANNEL_INFO_CHANGE struct {
	CMD_BASE
}

func (this *SC_CHANNEL_INFO_CHANGE) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type CS_NOTICE_CHANGE struct {
	CMD_BASE
}
type SC_NOTICE_CHANGE struct {
	CMD_BASE
}

func (this *SC_NOTICE_CHANGE) FromBytes(buff []byte) (err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, buff); err != nil {
		return
	}
	if err = binary.Read(buf, binary.LittleEndian, this); err != nil {
		return
	}
	return nil
}

type CS_CUSTOM_SERVICE_MSG struct {
	CMD_BASE
	UserID uint32
	MsgID  uint32
}

func CmdToBytes(cmd interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, cmd)
	if err != nil {
		println(err.Error())
	}
	size := buf.Len()
	//println("package size:%d", size)
	buf.Bytes()[0] = byte(size)
	buf.Bytes()[1] = byte(size >> 8)
	return buf.Bytes()
}
