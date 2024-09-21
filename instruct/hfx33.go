package instruct

import "github.com/vsm0/8chi-go/chip"

func Hfx33(c *chip.Chip, o chip.Operation) error {
	vx := c.Reg[o.X()]
	c.Ram[c.Index + 0] = vx / 100
	c.Ram[c.Index + 1] = (vx / 10) % 10
	c.Ram[c.Index + 2] = vx % 10
	return nil
}
