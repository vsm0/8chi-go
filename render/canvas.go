package render

import (
	"image/color"
	"math"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type Canvas struct {
	buffer []uint8
	image *pixel.PictureData
	sprite *pixel.Sprite
}

func NewCanvas(w, h int) *Canvas {
	r := pixel.R(0, 0, float64(w), float64(h))
	pd := pixel.MakePictureData(r)
	spr := pixel.NewSprite(pd, r)

	c := &Canvas{
		buffer: make([]uint8, w * h),
		image: pd,
		sprite: spr,
	}

	for x := 0; x < int(r.W()); x++ {
		for y := 0; y < int(r.H()); y++ {
			c.Set(x, y, 0)
		}
	}

	return c
}

func (s *Canvas) Size() (int, int) {
	b := s.image.Bounds()
	return int(b.W()), int(b.H())
}

func (s *Canvas) Set(x, y int, b uint8) {
	w, _ := s.Size()
	s.buffer[x + y * w] = b
}

func (s *Canvas) Get(x, y int) uint8 {
	w, _ := s.Size()
	return s.buffer[x + y * w]
}

func (s *Canvas) Draw(window *opengl.Window) {
	w, h := s.Size()

	colors := []color.RGBA{
		color.RGBA{A:0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			i := x + y * w
			b := s.buffer[i]
			s.image.Pix[i] = colors[b]
		}
	}

	sw := window.Bounds().W()
	sh := window.Bounds().H()
	scale := math.Min(sw / float64(w), sh / float64(h))

	pos := window.Bounds().Center()
	mat := pixel.IM.Moved(pos)
	mat = mat.ScaledXY(pos, pixel.V(scale, scale))

	s.sprite.Draw(window, mat)
}
