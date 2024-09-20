package chip

func (c *Chip) Cycle() error {
	op, err := c.Fetch()
	if err != nil {
		return err
	}

	return c.DecodeAndExecute(Operation(op))
}
