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
		switch o.Byte() {
		case 0xe0:
			return H00e0(c, o)
		case 0xee:
			return H00ee(c, o)
		}
	case 0x1:
		return H1nnn(c, o)
	case 0x2:
		return H2nnn(c, o)
	case 0x3:
		return H3xnn(c, o)
	case 0x4:
		return H4xnn(c, o)
	case 0x5:
		return H5xy0(c, o)
	case 0x6:
		return H6xnn(c, o)
	case 0x7:
		return H7xnn(c, o)
	case 0x8:
		switch o.Short() {
		case 0x0:
			return H8xy0(c, o)
		case 0x1:
			return H8xy1(c, o)
		case 0x2:
			return H8xy2(c, o)
		case 0x3:
			return H8xy3(c, o)
		case 0x4:
			return H8xy4(c, o)
		case 0x5:
			return H8xy5(c, o)
		case 0x6:
			return H8xy6(c, o)
		case 0x7:
			return H8xy7(c, o)
		case 0xe:
			return H8xye(c, o)
		}
	case 0x9:
		return H9xy0(c, o)
	case 0xa:
		return Hannn(c, o)
	case 0xd:
		return Hdxyn(c, o)
	case 0xf:
		switch o.Byte() {
		case 0x1e:
			return Hfx1e(c, o)
		case 0x33:
			return Hfx33(c, o)
		case 0x55:
			return Hfx55(c, o)
		case 0x65:
			return Hfx65(c, o)
		}
	}

	return errors.New(fmt.Sprintf("invalid command: %x (%v)", o, o))
}
