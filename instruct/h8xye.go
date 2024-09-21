package instruct

import "github.com/vsm0/8chi-go/chip"

func H8xye(c *chip.Chip, o chip.Operation) error {
	vx := c.Reg[o.X()]

	if c.Quirks.Check(chip.ShiftQuirk) {
		vx = c.Reg[o.Y()]
	}

	c.Reg[0xf] = vx >> 7

	c.Reg[o.X()] <<= 1

	return nil
}
