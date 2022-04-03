/**
 * @Author: boyyang
 * @Date: 2022-02-16 15:14:21
 * @LastEditTime: 2022-04-03 12:02:50
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\utils\returnData.go
 */

package utils

type Code struct {
	Code  int
	Msg   string
	Count int
}
type ReturnMsg struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count,omitempty"`
	Data  interface{} `json:"data"`
}

func RetunMsgFunc(code Code, data interface{}) *ReturnMsg {
	rm := new(ReturnMsg)
	rm.Code = code.Code
	rm.Count = code.Count
	rm.Msg = code.Msg
	rm.Data = data
	return rm
}
