package main

import (
	"github.com/vsm0/8chi-go/chip"
	"github.com/vsm0/8chi-go/font"
//	"github.com/vsm0/8chi-go/rom"

	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := chip.New()
	font.Load(c.Ram, 0x50, font.Default)

	// rom.Load(c.Ram, 0x200, "game.rom")

	// fetch, load and execute
}
