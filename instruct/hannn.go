package instruct

import "github.com/vsm0/8chi-go/chip"

func Hannn(c *chip.Chip, o chip.Operation) error {
	c.Index = o.Addr()
	return nil
}
