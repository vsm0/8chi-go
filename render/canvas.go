package render

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Canvas struct {
	buffer []uint8
	image *ebiten.Image
}

func NewCanvas(w, h int) *Canvas {
	return &Canvas{
		buffer: make([]uint8, w * h),
		image: ebiten.NewImage(w, h),
	}
}

func (s *Canvas) Size() (int, int) {
	b := s.image.Bounds()
	return b.Dx(), b.Dy()
}

func (s *Canvas) Set(x, y int, b uint8) {
	w, _ := s.Size()
	s.buffer[x + y * w] = b
}

func (s *Canvas) Get(x, y int) uint8 {
	w, _ := s.Size()
	return s.buffer[x + y * w]
}

func (s *Canvas) Draw(screen *ebiten.Image) {
	w, h := s.Size()

	colors := []color.Color{
		color.Black,
		color.White,
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			b := s.buffer[x + y * w]
			s.image.Set(x, y, colors[b])
		}
	}

	sw := float64(screen.Bounds().Dx())
	sh := float64(screen.Bounds().Dy())
	scale := math.Min(sw / float64(w), sh / float64(h))
	scw := float64(w) * scale
	sch := float64(h) * scale

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate((sw - scw) / 2, (sh - sch) / 2)

	screen.DrawImage(s.image, opts)
}
