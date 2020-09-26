package main

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age    int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	if pass, _ := regexp.MatchString(
		`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, input,
	); pass {
		return true
	}
	return false
}

func validate(v interface{}) (bool, string) {
	validateResult := true
	errmsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vv.NumField(); i++ {
		fieldValue := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		k := fieldValue.Kind()
		switch k {
		case reflect.Int:
			val := fieldValue.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errmsg = "validate int failed, tag is: " + strconv.FormatInt(
					tagVal, 10,
				)
				validateResult = false
			}
		case reflect.String:
		case reflect.Struct:

		}
	}

	return validateResult, errmsg
}

func main() {
	var a = T{Age: 10, Nested: Nested{Email: "abc@abc.com"}}
	validateResult, errmsg := validate(a)
	log.Println(validateResult, errmsg)

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
