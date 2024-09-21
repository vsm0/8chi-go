package main

import (
	"github.com/vsm0/8chi-go/chip"
	"github.com/vsm0/8chi-go/font"
	"github.com/vsm0/8chi-go/instruct"
	"github.com/vsm0/8chi-go/rom"

	"image/color"
	"math/rand"

	"github.com/gopxl/pixel/v2/backends/opengl"
)

var backgroundColor = color.RGBA{0x1f, 0x1f, 0x1f, 0xff}

type Game struct {
	Machine *chip.Chip
	Tps, Frequency, Cycles int
	Limited bool
}

func NewGame(cfg *Config) (*Game, error) {
	c := chip.New()
	c.Pc = cfg.Pc

	if err := rom.Load(c.Ram, c.Pc, cfg.File); err != nil {
		return nil, err
	}

	if err := font.Load(c.Ram, cfg.FontAddr, font.Default); err != nil {
		return nil, err
	}

	rand.Seed(cfg.Seed)

	if cfg.QuirkShift {
		c.Quirks.Set(chip.ShiftQuirk)
	}
	if cfg.QuirkJump {
		c.Quirks.Set(chip.JumpQuirk)
	}
	if cfg.QuirkLoad {
		c.Quirks.Set(chip.LoadQuirk)
	}

	return &Game{
		Machine: c,
		Tps: cfg.Tps,
		Frequency: cfg.Frequency / cfg.Tps,
		Cycles: cfg.Cycles,
		Limited: cfg.Limited,
	}, nil
}

func (g *Game) Update(window *opengl.Window) error {
	g.Machine.Countdown()

	for i := 0; i < g.Frequency; i++ {
		if g.Limited {
			if g.Cycles > 0 {
				g.Cycles--
			} else {
				break
			}
		}

		if err := instruct.Cycle(g.Machine); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(window *opengl.Window) {
	window.Clear(backgroundColor)

	g.Machine.Canvas.Draw(window)
}
