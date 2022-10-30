package controller

import "goweb32_bells-of-ireland/models"

type _ResponsePostList struct {
	Code ResCode                `json:"code"` //响应码
	Msg  string                 `json:"msg"`  //描述信息
	Data []models.ApiPostDetail `json:"data"` //响应数据
}
