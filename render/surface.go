package render

import "image/color"

type Surface interface {
	Size() (uint, uint)
	SetPixel(int, int, color.Color)
}
