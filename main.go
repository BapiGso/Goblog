package main

import (
	"SMOE/moe"
)

func main() {
	//todo ajax加载动画、音乐播放进度条、js重置、后台重构
	s := moe.New()
	s.Bind()
	s.Init()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
