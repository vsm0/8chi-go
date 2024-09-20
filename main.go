package main

import (
	"github.com/vsm0/8chi-go/chip"
	"github.com/vsm0/8chi-go/font"
	"github.com/vsm0/8chi-go/rom"

	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	Chip *chip.Chip
}

func main() {
	if len(os.Args) < 2 {
		panic("missing rom name")
	}

	rand.Seed(time.Now().UnixNano())

	c := chip.New()
	
	if err := font.Load(c.Ram, 0x50, font.Default); err != nil {
		panic(err)
	}

	c.Pc = 0x200

	if err := rom.Load(c.Ram, c.Pc, os.Args[1]); err != nil {
		panic(err)
	}

	ebiten.SetWindowTitle("Untitled")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizable(true)

	g := &game{c}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}


func (g *game) Layout(w, h int) (int, int) {
	return g.Chip.Canvas.Size()
}

func (g *game) Update() error {
	op, err := g.Chip.Fetch()
	if err != nil {
		return err
	}

	return g.Chip.DecodeAndExecute(chip.Operation(op))
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Chip.Canvas.Draw(screen)
}
