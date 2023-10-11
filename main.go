package main

import (
	"SMOE/moe"
)

func main() {
	s := moe.New()
	s.Bind()
	s.Init()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
