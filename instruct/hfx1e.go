package instruct

import "github.com/vsm0/8chi-go/chip"

func Hfx1e(c *chip.Chip, o chip.Operation) error {
	c.Index += uint16(c.Reg[o.X()])
	return nil
}
