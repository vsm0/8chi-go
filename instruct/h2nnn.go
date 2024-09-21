package instruct

import "github.com/vsm0/8chi-go/chip"

func H2nnn(c *chip.Chip, o chip.Operation) error {
	c.Stack.Push(c.Pc)
	c.Pc = o.Addr()
	return nil
}
