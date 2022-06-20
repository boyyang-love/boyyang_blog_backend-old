/**
 * @Author: boyyang
 * @Date: 2022-06-14 10:42:56
 * @LastEditTime: 2022-06-14 13:29:30
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\socket\socker.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Socket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := models.User{}
	global.DB.Find(&user)
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		fmt.Println(message)
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
			ws.WriteMessage(mt, message)
		}
		if err == nil {
			msg, _ := json.Marshal(&user)
			//写入ws数据
			err = ws.WriteMessage(mt, msg)
			if err != nil {
				break
			}
		}
	}
}
