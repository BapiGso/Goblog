package main

import "SMOE/smoe"

func main() {
	s := smoe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
