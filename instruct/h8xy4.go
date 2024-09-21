package instruct

import "github.com/vsm0/8chi-go/chip"

func H8xy4(c *chip.Chip, o chip.Operation) error {
	res := uint16(c.Reg[o.X()]) + uint16(c.Reg[o.Y()])
	c.Reg[o.X()] = uint8(res)

	if res > 0xff {
		c.Reg[0xf] = 1
	} else {
		c.Reg[0xf] = 0
	}

	return nil
}
