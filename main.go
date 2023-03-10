package main

import "github.com/BapiGso/SMOE/smoe"

func main() {
	s := smoe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
