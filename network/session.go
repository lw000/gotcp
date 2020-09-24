package network

import (
	"github.com/gogo/protobuf/proto"
	"log"
	"net"
)

type MessageHandler func(session *Session, msg *Message)

type Session struct {
	serve     *TCPService
	conn      net.Conn
	sessionID uint32
	onHandler MessageHandler
	connected int32
}

func newSession(serve *TCPService, conn net.Conn, sessionID uint32, fn MessageHandler) *Session {
	return &Session{
		serve:     serve,
		conn:      conn,
		sessionID: sessionID,
		onHandler: fn,
	}
}

func (s *Session) SetSessionID(sessionID uint32) {
	s.sessionID = sessionID
}

func (s *Session) SessionID() uint32 {
	return s.sessionID
}

func (s *Session) SetConn(conn net.Conn) {
	s.conn = conn
}

func (s *Session) Conn() net.Conn {
	return s.conn
}

func (s *Session) Run() {
	defer func() {
		if p := recover(); p != nil {
			log.Println(p)
		}
		if s.serve != nil {
			s.serve.RemoveSession(s.sessionID)
		}
		log.Println("conn closed")
	}()

	for {
		pack := NewDataPack()
		// 先读出流中的head部分
		headData := make([]byte, pack.GetHeadLen())
		_, err := s.conn.Read(headData) // ReadFull 会把msg填充满为止
		if err != nil {
			log.Println("read head error")
			return
		}

		// 将headData字节流 拆包到msg中
		msgHead, err := pack.Unpack(headData)
		if err != nil {
			log.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			// msg 是有data数据的，需要再次读取data数据
			msg := msgHead
			msg.Data = make([]byte, msg.GetDataLen())

			// 根据dataLen从io中读取字节流
			_, err := s.conn.Read(msg.Data)
			if err != nil {
				log.Println("read data err:", err)
				return
			}

			if s.onHandler != nil {
				s.onHandler(s, msg)
			}
		}
	}
}

func (s *Session) SendMessage(clientId, msgId uint32, data []byte) error {
	// 将data封包，并且发送
	pack := NewDataPack()
	msgData, err := pack.Pack(NewMsgPackage(clientId, msgId, data))
	if err != nil {
		return err
	}
	_, err = s.conn.Write(msgData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) SendProtoMessage(clientId, msgId uint32, pb proto.Message) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return s.SendMessage(clientId, msgId, data)
}
