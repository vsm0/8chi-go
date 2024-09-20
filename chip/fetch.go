package chip

import "errors"

func (c *Chip) Fetch() (uint16, error) {
	pc := c.Pc

	if l := uint16(len(c.Ram)); l < pc || l < pc + 1 {
		return 0, errors.New("pc exceed ram during fetch")
	}

	c.Pc += 2

	return uint16(c.Ram[pc]) << 8 | uint16(c.Ram[pc + 1]), nil
}
