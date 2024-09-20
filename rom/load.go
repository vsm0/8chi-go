package rom

import (
	"io"
	"os"
)

func Load(src []uint8, addr uint16, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if _, err = file.Read(src[addr:]); err != nil && err != io.EOF {
		return err
	}

	return nil
}
