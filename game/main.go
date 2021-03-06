package main

import (
	"demo/gotcp/network"
	"demo/gotcp/protos/msg"
	"github.com/gogo/protobuf/proto"
	"log"
)

var (
	serve *network.TCPService
)

func main() {
	serve = network.NewTCPService()
	serve.SetMessageHandler(func(session *network.Session, msg *network.Message) {
		var req Tmsg.ReqEnterGame
		err := proto.Unmarshal(msg.GetData(), &req)
		if err != nil {
			log.Println("proto unmarshal err ", err)
			return
		}
		log.Println("recv from client : ", req)

		var ack Tmsg.AckEnterGame
		ack.Code = 0
		ack.Msg = "enterGame"
		err = session.SendProtoMessage(msg.GetClientId(), msg.GetMsgId(), &ack)
		if err != nil {
			log.Println("error:", err)
		}
	})
	err := serve.Start(8890)
	if err != nil {
		log.Fatal(err)
	}
	serve.Serve()
}
