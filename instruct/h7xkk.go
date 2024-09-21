package instruct

import "github.com/vsm0/8chi-go/chip"

func H7xkk(c *chip.Chip, o chip.Operation) error {
	c.Reg[o.X()] += o.Byte()
	return nil
}
