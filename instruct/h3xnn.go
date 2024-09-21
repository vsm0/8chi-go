package instruct

import "github.com/vsm0/8chi-go/chip"

func H3xnn(c *chip.Chip, o chip.Operation) error {
	if c.Reg[o.X()] == o.Byte() {
		c.Pc += 2
	}

	return nil
}
