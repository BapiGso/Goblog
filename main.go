package main

import (
	"github.com/BapiGso/SMOE/smoe"
)

func main() {
	s := Smoe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.LoadMiddlewareRoutes()
	s.Listen()
}
