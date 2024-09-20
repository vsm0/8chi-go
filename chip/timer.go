package chip

func (c *Chip) Countdown() {
	if c.DelayTimer > 0 {
		c.DelayTimer--
	}

	if c.SoundTimer > 0 {
		c.SoundTimer--
	}
}
