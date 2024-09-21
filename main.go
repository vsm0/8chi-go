package main

import (
	"time"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

func main() {
	opengl.Run(run)
}

func run() {
	winCfg := opengl.WindowConfig{
		Title: "8Chi-Go Emulator",
		Bounds: pixel.R(0, 0, 480, 360),
		Resizable: true,
	//	VSync: true,
	}

	win, err := opengl.NewWindow(winCfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	gameCfg := NewConfig()
	g, err := NewGame(gameCfg)
	if err != nil {
		panic(err)
	}

	dur := time.Second / time.Duration(g.Tps)
	ticker := time.NewTicker(dur)

	for range ticker.C {
		if win.Closed() {
			break
		}

		win.Update()

		// tick
		err := g.Update(win)
		if err != nil {
			panic(err)
		}

		g.Draw(win)
	}
}
