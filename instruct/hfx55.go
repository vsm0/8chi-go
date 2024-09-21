package instruct

import "github.com/vsm0/8chi-go/chip"

func Hfx55(c *chip.Chip, o chip.Operation) error {
	max := int(o.X())
	for i := 0; i <= max; i++ {
		c.Ram[int(c.Index) + i] = c.Reg[i]
	}

	return nil
}
