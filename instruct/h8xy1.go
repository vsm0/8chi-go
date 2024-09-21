package instruct

import "github.com/vsm0/8chi-go/chip"

func H8xy1(c *chip.Chip, o chip.Operation) error {
	c.Reg[o.X()] |= c.Reg[o.Y()]
	return nil
}
