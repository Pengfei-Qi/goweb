package mysql

import "errors"

var (
	ErrorAccountExit = errors.New("邮箱已被注册")
	ErrorPramValid   = errors.New("参数值无效")
)
