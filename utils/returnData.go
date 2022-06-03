/**
 * @Author: boyyang
 * @Date: 2022-02-16 15:14:21
 * @LastEditTime: 2022-06-03 10:57:15
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\utils\returnData.go
 */

package utils

type Message struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Error error       `json:"error,omitempty"`
	Count int         `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func Msg(msg Message) *Message {
	if msg.Data == nil {
		return &Message{
			Code:  msg.Code,
			Msg:   msg.Msg,
			Error: msg.Error,
			Count: msg.Count,
		}
	} else {
		return &Message{
			Code:  msg.Code,
			Msg:   msg.Msg,
			Error: msg.Error,
			Data:  msg.Data,
			Count: msg.Count,
		}
	}
}
