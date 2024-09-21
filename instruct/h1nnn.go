package instruct

import "github.com/vsm0/8chi-go/chip"

func H1nnn(c *chip.Chip, o chip.Operation) error {
	c.Pc = o.Addr()
	return nil
}
