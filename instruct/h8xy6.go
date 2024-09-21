package instruct

import "github.com/vsm0/8chi-go/chip"

func H8xy6(c *chip.Chip, o chip.Operation) error {
	vx := c.Reg[o.X()]

	if c.Quirks.Check(chip.ShiftQuirk) {
		vx = c.Reg[o.Y()]
	}

	c.Reg[0xf] = vx & 0x1

	c.Reg[o.X()] = vx >> 1

	return nil
}
