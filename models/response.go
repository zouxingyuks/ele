package models

type Response struct {
	Msg  string      `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}
