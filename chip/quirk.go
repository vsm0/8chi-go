package chip

type Quirk uint8

const (
	ShiftQuirk Quirk = 1 << iota
	LoadQuirk
	JumpQuirk
)

func (q *Quirk) Set(v Quirk) {
	*q |= v
}

func (q *Quirk) Unset(v Quirk) {
	*q &= ^v
}

func (q Quirk) Check(v Quirk) bool {
	return q & v != 0
}
