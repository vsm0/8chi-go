package instruct

import "github.com/vsm0/8chi-go/chip"

func H00ee(c *chip.Chip, o chip.Operation) error {
	addr, err := c.Stack.Pop()
	if err != nil {
		return err
	}

	c.Pc = addr
	return nil
}
