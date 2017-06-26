/*
协议格式：package + length + message
message = messageId + messageType + body
*/

package protocal

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

// 消息类型
const (
	MSG_TYPE_REQUEST  = 0
	MSG_TYPE_NOTIFY   = 1
	MSG_TYPE_RESPONSE = 2
	MSG_TYPE_PUSH     = 3
)

// 消息类型，消息分成两种，一种是有messageId的，一种是没有messageId的
const (
	PACKAGE_TYPE_NONE_MESSAGE_ID = 0 // 无消息id类型的消息
	PACKAGE_TYPE_HAVE_MESSAGE_ID = 1 // 有消息id类型的消息
)

// 包长度定义
const (
	// MAGIC_NUM = uint16(1314)

	// MAGIC_NUM_SIZE = 2
	PACKAGE_SIZE = 1
	LENGTH_SIZE  = 3
	HEADER_SIZE  = 4 // PACKAGE_SIZE + LENGTH_SIZE

	MESSAGE_ID_SIZE     = 2
	MESSAGE_TYPE_SIZE   = 2
	MESSAGE_NUMBER_SIZE = 4
)

// 包类型定义
const (
	PACKAGE_TYPE_HANDSHAKE     = uint8(1)   // 握手
	PACKAGE_TYPE_HANDSHAKE_ACK = uint8(2)   // 握手回复
	PACKAGE_TYPE_HEARTBEAT     = uint8(3)   // 心跳
	PACKAGE_TYPE_DATA          = uint8(4)   // 数据包
	PACKAGE_TYPE_KICK          = uint8(5)   // 退出、踢出
	PACKAGE_TYPE_SYSTEM        = uint8(100) // 系统消息
)

type ImPacket struct {
	buff  []byte
	IType int
}

func (this *ImPacket) Serialize() []byte {
	return this.buff
}

// 生成一个message，有messageId和messageType的
func NewImMessage(mId uint16, mType uint16, mNumber uint32, body []byte) []byte {
	mBuff := make([]byte, MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE+MESSAGE_NUMBER_SIZE+len(body))

	// 写入messageId
	binary.BigEndian.PutUint16(mBuff[0:MESSAGE_ID_SIZE], mId)
	// 写入messageType
	binary.BigEndian.PutUint16(mBuff[MESSAGE_ID_SIZE:MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE], mType)
	// 写入messageNumber
	binary.BigEndian.PutUint32(mBuff[MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE:MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE+MESSAGE_NUMBER_SIZE], mNumber)

	// 写入body
	copy(mBuff[MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE+MESSAGE_NUMBER_SIZE:], body)

	return mBuff
}

// 生成一条消息
func NewImPacket(packageId uint8, message []byte) *ImPacket {
	p := &ImPacket{}

	// 判断消息类型，是否携带messageId
	if packageId == PACKAGE_TYPE_DATA {
		p.IType = PACKAGE_TYPE_HAVE_MESSAGE_ID
	} else {
		p.IType = PACKAGE_TYPE_NONE_MESSAGE_ID
	}

	p.buff = make([]byte, PACKAGE_SIZE+LENGTH_SIZE+len(message))

	// 写入packageId
	// binary.BigEndian.PutUint16(p.buff[0:1], uint16(MAGIC_NUM))
	p.buff[0] = byte(packageId)

	// 写入包长
	putLength(p.buff[PACKAGE_SIZE:PACKAGE_SIZE+LENGTH_SIZE], uint32(len(message)))

	// 写入包内容
	copy(p.buff[PACKAGE_SIZE+LENGTH_SIZE:], message)

	return p
}

// 从字节流中读出包类型
func (this *ImPacket) GetPackage() uint8 {
	return uint8(this.buff[0])
}

// 从字节流中读出长度
func (this *ImPacket) GetLength() uint32 {
	return length(this.buff[PACKAGE_SIZE : PACKAGE_SIZE+LENGTH_SIZE])
}

// 读取消息内容
func (this *ImPacket) GetMessage() []byte {
	return this.buff[HEADER_SIZE:]
}

// 读出messageId
func (this *ImPacket) GetMessageId() uint16 {
	if this.IType == PACKAGE_TYPE_HAVE_MESSAGE_ID {
		messageBytes := this.GetMessage()
		return binary.BigEndian.Uint16(messageBytes[0:MESSAGE_ID_SIZE])
	} else {
		return uint16(0)
	}
}

// 读出messageType
func (this *ImPacket) GetMessageType() uint16 {
	if this.IType == PACKAGE_TYPE_HAVE_MESSAGE_ID {
		messageBytes := this.GetMessage()
		return binary.BigEndian.Uint16(messageBytes[MESSAGE_ID_SIZE : MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE])
	} else {
		return uint16(0)
	}
}

// 读出messageNumber
func (this *ImPacket) GetMessageNumber() uint32 {
	if this.IType == PACKAGE_TYPE_HAVE_MESSAGE_ID {
		messageBytes := this.GetMessage()
		return binary.BigEndian.Uint32(messageBytes[MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE : MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE+MESSAGE_NUMBER_SIZE])
	} else {
		return uint32(0)
	}
}

// 读出消息正文
func (this *ImPacket) GetBody() []byte {
	if this.IType == PACKAGE_TYPE_HAVE_MESSAGE_ID {
		messageBytes := this.GetMessage()
		return messageBytes[MESSAGE_ID_SIZE+MESSAGE_TYPE_SIZE+MESSAGE_NUMBER_SIZE:]
	} else {
		return this.GetMessage()
	}
}

// 读取一条消息
func ReadPacket(conn *net.TCPConn) (*ImPacket, error) {
	var (
		packageBytes []byte = make([]byte, PACKAGE_SIZE)
		lengthBytes  []byte = make([]byte, LENGTH_SIZE)

		packageId uint8
	)

	// 读取package
	if _, err := io.ReadFull(conn, packageBytes); err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		} else {
			return nil, errors.New(fmt.Sprintf("pakageId read error: %s.", err.Error()))
		}
	}
	// 转成uint8
	packageId = packageBytes[0]

	// 读取lengthBytes
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		} else {
			return nil, errors.New(fmt.Sprintf("packet length read error: %s", err.Error()))
		}
	}
	// 内容长度
	mLength := length(lengthBytes)

	// 读取message
	messageBytes := make([]byte, mLength)
	if _, err := io.ReadFull(conn, messageBytes); err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		} else {
			return nil, errors.New(fmt.Sprintf("read packet message error: %s", err.Error()))
		}
	}

	return NewImPacket(packageId, messageBytes), nil
}

// 写入长度
func putLength(b []byte, v uint32) {
	_ = b[2] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

// 获取长度
func length(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[2]) | uint32(b[1])<<8 | uint32(b[0])<<16
}

// 发送消息
func (this *ImPacket) Send(conn *net.TCPConn) {
	conn.Write(this.Serialize())
}
