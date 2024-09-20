package chip

type Operation uint16

func (o Operation) Fun() uint8 {
	return uint8(o >> 12)
}

func (o Operation) X() uint8 {
	return uint8(o >> 8) & 0xf
}

func (o Operation) Y() uint8 {
	return uint8(o >> 4) & 0xf
}

func (o Operation) Short() uint8 {
	return uint8(o & 0xf)
}

func (o Operation) Byte() uint8 {
	return uint8(o & 0xff)
}

func (o Operation) Addr() uint16 {
	return uint16(o & 0xfff)
}
