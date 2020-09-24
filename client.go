package main

import (
	"demo/gotcp/network"
	"demo/gotcp/protos/msg"
	"github.com/gogo/protobuf/proto"
	"log"
	"time"
)

func RunClient2(host, port string) {
	client := network.NewTCPClient()
	client.SetMessageHandler(func(session *network.Session, msg *network.Message) {
		msgId := msg.GetMsgId()
		switch msgId {
		case 1001:
			var ack Tmsg.AckRegister
			err := proto.Unmarshal(msg.GetData(), &ack)
			if err != nil {
				log.Println("proto unmarshal err ", err)
				return
			}
			log.Println("Tmsg.AckRegister :", msg.GetMsgId(), ack.Code, ack.Msg, ack.Token)
		case 2001:
			var ack Tmsg.AckEnterGame
			err := proto.Unmarshal(msg.GetData(), &ack)
			if err != nil {
				log.Println("proto unmarshal err ", err)
				return
			}
			log.Println("Tmsg.AckEnterGame :", msg.GetMsgId(), ack.Code, ack.Msg)
		}

	})
	err := client.Connect(host, port, time.Second*15)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for i := 0; i < 10000; i++ {
			{
				var req Tmsg.ReqRegister
				req.Account = "1313123123"
				req.Password = "12311212312"
				err = client.Session().SendProtoMessage(0, 1001, &req)
				if err != nil {
					log.Println("proto unmarshal err ", err)
					return
				}
			}

			{
				var req Tmsg.ReqEnterGame
				req.RoomId = 1000
				req.GameId = 2000
				err = client.Session().SendProtoMessage(0, 2001, &req)
				if err != nil {
					log.Println("proto unmarshal err ", err)
					return
				}
			}
		}
	}()
	client.Loop()
}
