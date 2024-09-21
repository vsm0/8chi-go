package instruct

import "github.com/vsm0/8chi-go/chip"

func Hdxyn(c *chip.Chip, o chip.Operation) error {
	w, h := c.Canvas.Size()
	x := int(c.Reg[o.X()]) % w
	y := int(c.Reg[o.Y()]) % h
	n := int(o.Short())

	c.Reg[0xf] = 0
	idx := c.Index

	for row := 0; row < n && row + y < h; row++ {
		line := c.Ram[idx]
		idx++

		for col := 0; col < 8 && col + x < w; col++ {
			bit := (line >> (7 - col)) & 0x1
			old := c.Canvas.Get(x + col, y + row)
			cur := bit ^ old

			c.Canvas.Set(x + col, y + row, cur)

			//if cur == 0 && old != 0 {
			//	c.Reg[0xf] = 1
			//}
			c.Reg[0xf] |= (cur & 1) ^ (old & 1)
		}
	}
		
	return nil
}
