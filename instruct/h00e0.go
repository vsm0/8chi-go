package instruct

import "github.com/vsm0/8chi-go/chip"

func H00e0(c *chip.Chip, o chip.Operation) error {
	w, h := c.Canvas.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c.Canvas.Set(x, y, 0)
		}
	}

	return nil
}
