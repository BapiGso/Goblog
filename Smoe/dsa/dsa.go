package dsa

import (
	"fmt"
	_ "main/smoe"
)

type Smoe struct {
}

func (p *Smoe) SayHello() {
	fmt.Printf("Hello, my name is %v and I am %v years old.\n")
}
