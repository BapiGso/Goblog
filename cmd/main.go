package main

import (
	"smoe/moe"
)

func main() {
	s := moe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
