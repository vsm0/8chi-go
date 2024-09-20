package render

import "image/color"

type Canvas struct {
	width, height uint
	buffer []uint8
}

func NewCanvas(w, h uint) *Canvas {
	return &Canvas{
		width: w,
		height: h,
		buffer: make([]uint8, w * h),
	}
}

func (s *Canvas) Size() (uint, uint) {
	return s.width, s.height
}

func (s *Canvas) SetPixel(x, y int, c color.Color) {

}
