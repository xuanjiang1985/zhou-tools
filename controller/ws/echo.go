package ws

import (
	"fmt"
	"net/http"
	scan "zhou/tools/controller"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	for {
		// 读取客户端的消息
		// msgType, msg, err := conn.ReadMessage()
		// if err != nil {
		// 	return
		// }

		// // 把消息打印到标准输出
		// fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		select {
		case s := <-scan.MsgChan:

			// 把消息写回客户端，完成回音
			if err = conn.WriteMessage(websocket.TextMessage, []byte(s)); err != nil {
				return
			}
		}

	}
}
