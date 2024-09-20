package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	cfg := NewConfig()
	g, err := NewGame(cfg)
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowTitle("8Chi-Go Emulator")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizable(true)
	ebiten.SetMaxTPS(cfg.Tps)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
