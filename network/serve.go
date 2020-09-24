package network

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

type TCPService struct {
	sessionLock sync.RWMutex
	listener    *net.TCPListener
	onHandler   MessageHandler
	sessions    map[uint32]*Session
}

func NewTCPService() *TCPService {
	return &TCPService{
		sessions: make(map[uint32]*Session),
	}
}

func (s *TCPService) SetMessageHandler(fn MessageHandler) {
	s.onHandler = fn
}

func (s *TCPService) Start(port int) error {
	addr, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort("0.0.0.0", fmt.Sprintf("%d", port)))
	if err != nil {
		return err
	}

	s.listener, err = net.ListenTCP("tcp4", addr)
	if err != nil {
		return err
	}

	return nil
}

func (s *TCPService) Stop() {
	_ = s.listener.Close()
}

func (s *TCPService) AddSession(cid uint32, session *Session) {
	s.sessionLock.Lock()
	s.sessions[cid] = session
	s.sessionLock.Unlock()
}

func (s *TCPService) RemoveSession(cid uint32) {
	s.sessionLock.Lock()
	delete(s.sessions, cid)
	s.sessionLock.Unlock()
}

func (s *TCPService) Serve() {
	var cid uint32 = 1
	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("accept conn", conn.RemoteAddr())

		session := newSession(s, conn, cid, s.onHandler)
		s.AddSession(cid, session)
		cid++
		go session.Run()
	}
}

func (s *TCPService) SendMessage(clientID uint32, msgID uint32, data []byte) error {
	s.sessionLock.RLock()
	defer s.sessionLock.RUnlock()
	sess, ok := s.sessions[clientID]
	if ok {
		return sess.SendMessage(clientID, msgID, data)
	}
	return errors.New("session not found")
}
