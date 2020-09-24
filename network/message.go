package network

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const MaxPacketSize = 1024 * 10

type Message struct {
	DataLen  uint32 // 消息的长度
	MsgId    uint32 // 消息的ID
	ClientId uint32 // 消息客户端ID
	Data     []byte // 消息的内容
}

// 封包拆包类实例，暂时不需要成员
type DataPack struct{}

// 创建一个Message消息包
func NewMsgPackage(ClientId uint32, id uint32, data []byte) *Message {
	return &Message{
		DataLen:  uint32(len(data)),
		MsgId:    id,
		ClientId: ClientId,
		Data:     data,
	}
}

// 封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// 获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.MsgId
}

// 消息客户端ID
func (msg *Message) GetClientId() uint32 {
	return msg.ClientId
}

// 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// 设计消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.MsgId = msgId
}

// 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

// 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) +  DataLen uint32(4字节) + ClientId uint32(4字节)
	return 12
}

// 封包方法(压缩数据)
func (dp *DataPack) Pack(msg *Message) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 写clientID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetClientId()); err != nil {
		return nil, err
	}

	// 写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (*Message, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	// 只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}

	// 读clientID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ClientId); err != nil {
		return nil, err
	}

	// 判断dataLen的长度是否超出我们允许的最大包长度
	if msg.DataLen > MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	// 这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
