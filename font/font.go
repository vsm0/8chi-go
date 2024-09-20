package font

import (
	"errors"
	"fmt"
)

func Load(dst []uint8, addr uint16, src []uint8) error {
	if ls, ld := len(src), len(dst); int(addr) + ls > ld {
		return errors.New(fmt.Sprintf("src length (%d) exceeds dst (%d) at index %d (%x)", ls, ld, addr, addr))
	}

	copy(dst[addr:], src)

	return nil
}
