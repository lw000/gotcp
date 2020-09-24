package main

func main() {
	// serve := network.TCPService{}
	// serve.SetMessageHandler(func(session *network.Session, msg *network.Message) {
	// 	var req Tmsg.ReqRegister
	// 	err := proto.Unmarshal(msg.GetData(), &req)
	// 	if err != nil {
	// 		log.Println("proto unmarshal err ", err)
	// 		return
	// 	}
	// 	log.Println("recv from client : ", req)
	//
	// 	var ack Tmsg.AckRegister
	// 	ack.Code = 1000
	// 	ack.Msg = "login"
	// 	ack.Token = "123123123213213121231231"
	// 	err = session.SendProtoMessage(msg.GetMsgId(), &ack)
	// 	if err != nil {
	// 		log.Println("error:", err)
	// 	}
	// })
	//
	// err := serve.Start(8888)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// time.AfterFunc(time.Second, func() {
	// 	for i := 0; i < 1; i++ {
	// 		go RunClient2("127.0.0.1", "8888")
	// 	}
	// })
	// serve.Serve()
	for i := 0; i < 1; i++ {
		go RunClient2("127.0.0.1", "8888")
	}

	select {}
}
