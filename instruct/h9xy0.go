package instruct

import "github.com/vsm0/8chi-go/chip"

func H9xy0(c *chip.Chip, o chip.Operation) error {
	if c.Reg[o.X()] != c.Reg[o.Y()] {
		c.Pc += 2
	}

	return nil
}
