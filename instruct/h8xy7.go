package instruct

import "github.com/vsm0/8chi-go/chip"

func H8xy7(c *chip.Chip, o chip.Operation) error {
	vx := c.Reg[o.X()]
	vy := c.Reg[o.Y()]
	
	if vy > vx {
		c.Reg[0xf] = 1
	} else {
		c.Reg[0xf] = 0
	}

	c.Reg[o.X()] = vy - vx

	return nil
}
