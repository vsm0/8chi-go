package instruct

import "github.com/vsm0/8chi-go/chip"

func Hfx65(c *chip.Chip, o chip.Operation) error {
	max := int(o.X())
	for i := 0; i <= max; i++ {
		c.Reg[i] = c.Ram[int(c.Index) + i]
	}

	return nil
}
