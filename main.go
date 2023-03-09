package main

import (
	"main/smoe"
)

func main() {
	s := Smoe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
