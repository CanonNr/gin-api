package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Run() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := Upgrade.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// 读取来自浏览器的消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// 在控制台输出
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// 返回给浏览器
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	http.ListenAndServe(":8080", nil)
}
