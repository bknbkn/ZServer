package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

/*
	DataPack
	|  DataLen |   ID   | Data |
	|--4bytes--|-4bytes-|------|
*/
type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}

}
func (d *DataPack) GetHeadLen() uint32 {
	// DataLen 4 bytes + ID 4bytes = 8 bytes
	return 8
}

func (d *DataPack) Pack(msg serverinterface.IMessage) ([]byte, error) {
	dataBuffer := bytes.NewBuffer(make([]byte, 0, 16))

	// Write DataLen
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	// Write DataID
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// Write Data
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuffer.Bytes(), nil
}

func (d *DataPack) UnPackHead(binData []byte) (serverinterface.IMessage, error) {
	dataBuffer := bytes.NewReader(binData)
	msg := &Message{}

	// 读DataLen
	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if msg.GetMsgLen() > utils.GlobalConfig.MaxPackageSize {
		return nil, errors.New(fmt.Sprintf("msg length is %v, too large msg data", msg.GetMsgLen()))
	}

	// 读DataID
	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 具体data需要下一次读取
	return msg, nil
}
