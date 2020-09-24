package main

import (
	"demo/gotcp/network"
	"log"
)

var (
	serve      *network.TCPService
	loginProxy *network.TCPClient
	gameProxy  *network.TCPClient
)

func runLoginProxy(host, port string) {
	loginProxy = network.NewTCPClient()
	loginProxy.SetMessageHandler(func(session *network.Session, msg *network.Message) {
		err := serve.SendMessage(msg.GetClientId(), msg.GetMsgId(), msg.GetData())
		if err != nil {
			log.Println("error:", err)
		}
	})
	err := loginProxy.Connect(host, port)
	if err != nil {
		log.Fatal(err)
	}
	loginProxy.Loop()
}

func runGameProxy(host, port string) {
	gameProxy = network.NewTCPClient()
	gameProxy.SetMessageHandler(func(session *network.Session, msg *network.Message) {
		err := serve.SendMessage(msg.GetClientId(), msg.GetMsgId(), msg.GetData())
		if err != nil {
			log.Println("error:", err)
		}
	})
	err := gameProxy.Connect(host, port)
	if err != nil {
		log.Fatal(err)
	}
	gameProxy.Loop()
}

func main() {
	go runLoginProxy("127.0.0.1", "8889")
	go runGameProxy("127.0.0.1", "8890")

	serve = network.NewTCPService()
	serve.SetMessageHandler(func(session *network.Session, msg *network.Message) {
		msgID := msg.GetMsgId()
		if msgID > 0 && msgID < 1000 {

		} else if msgID > 1000 && msgID < 2000 {
			err := loginProxy.Session().SendMessage(session.SessionID(), msg.GetMsgId(), msg.GetData())
			if err != nil {
				log.Println("send login proxy message err ", err)
				return
			}
		} else if msgID > 2000 && msgID < 3000 {
			err := gameProxy.Session().SendMessage(session.SessionID(), msg.GetMsgId(), msg.GetData())
			if err != nil {
				log.Println("send game proxy message err ", err)
				return
			}
		} else {

		}
	})
	err := serve.Start(8888)
	if err != nil {
		log.Fatal(err)
	}
	serve.Serve()
}
