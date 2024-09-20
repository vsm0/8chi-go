package chip

import (
	"errors"
	"fmt"
)

func (c *Chip) DecodeAndExecute(o Operation) error {
	switch cmd := o.Fun(); cmd {
	case 0x0:
		if b := o.Byte(); b == 0xe0 {
			w, h := c.Canvas.Size()
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					c.Canvas.Set(x, y, 0)
				}
			}
			return nil
		}
	case 0x1:
		c.Pc = o.Addr()
		return nil
	case 0x6:
		c.Reg[o.X()] = o.Byte()
		return nil
	case 0x7:
		c.Reg[o.X()] += o.Byte()
		return nil
	case 0xa:
		c.Index = o.Addr()
		return nil
	case 0xd:
		w, h := c.Canvas.Size()
		x := int(c.Reg[o.X()]) % w
		y := int(c.Reg[o.Y()]) % h
		n := int(o.Short())

		c.Reg[0xf] = 0
		idx := c.Index

		for row := 0; row < n && row + y < h; row++ { // line exists in scr
			line := c.Ram[idx]
			idx++

			for col := 0; col < 8 && col + x < w; col++ { // pix exists in scr
				bit := (line >> col) & 0x1
				old := c.Canvas.Get(x + col, y + row)
				cur := bit ^ old

				c.Canvas.Set(x + col, y + row, cur)

				if cur == 0 && old != 0 {
					c.Reg[0xf] = 1
				}
			}
		}
		
		return nil
	}

	return errors.New(fmt.Sprintf("invalid command: %x (%v)", o, o))
}
