package main

import (
	"SMOE/moe"
)

func main() {
	//todo ajax加载动画（loader）、后台重构(css嵌套)、sql语句权限管理、追番、评论回复样式//数据库重构
	s := moe.New()
	s.Bind()
	s.Init()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
