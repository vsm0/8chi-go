package instruct

import (
	"github.com/vsm0/8chi-go/chip"

	"errors"
	"fmt"
)

func Cycle(c *chip.Chip) error {
	//fetch
	pc := c.Pc

	if l := uint16(len(c.Ram)); l < pc || l < pc + 1 {
		return errors.New("pc exceed ram during fetch")
	}

	c.Pc += 2

	op := uint16(c.Ram[pc]) << 8 | uint16(c.Ram[pc + 1])

	// decode and execute
	o := chip.Operation(op)

	switch cmd := o.Fun(); cmd {
	case 0x0:
		if b := o.Byte(); b == 0xe0 {
			return H00e0(c, o)
		}
	case 0x1:
		return H1nnn(c, o)
	case 0x6:
		return H6xkk(c, o)
	case 0x7:
		return H7xkk(c, o)
	case 0xa:
		return Hannn(c, o)
	case 0xd:
		return Hdxyn(c, o)
	}

	return errors.New(fmt.Sprintf("invalid command: %x (%v)", o, o))
}
