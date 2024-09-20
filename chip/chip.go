package chip

import (
	"github.com/vsm0/8chi-go/render"

	"github.com/vsm0/adt"
	"github.com/vsm0/adt/stack"
)

type Chip struct {
	Ram []uint8
	Reg []uint8
	Pc, Index uint8
	Stack adt.Stack[uint16]
	Canvas render.Surface
}

func New() *Chip {
	return &Chip{
		Ram: make([]uint8, 4096),
		Reg: make([]uint8, 16),
		Stack: stack.New[uint16](),
		Canvas: render.NewCanvas(64, 32),
	}
}
