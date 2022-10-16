package commandparser

type ZRange_ struct {
	Key        string
	Start      int
	Stop       int
	SortBy     *SortBy // Optional enum doesn't need a name
	Rev        *IsRev  // This is how we parse a boolean flag
	Limit      *Limit  `rname:"limit" json:"limit,omitempty"`
	WithScores *IsWithScores
}

func (cmd *ZRange) ParseCommand(args [][]byte) (err error) {
	// used in every command
	var present bool
	var s string
	var args2 [][]byte

	cmd.Key, args, err = parseString(args)
	if err != nil {
		return err
	}

	cmd.Start, args, err = parseInt(args)
	if err != nil {
		return err
	}

	cmd.Stop, args, err = parseInt(args)
	if err != nil {
		return err
	}

	s, args2, err = parseString(args)
	if err == nil {
		enumValue, ok := map[string]SortBy{
			"BYSCORE": SortByScore,
			"BYLEX":   SortByLex,
		}[s]
		if ok {
			cmd.SortBy = &enumValue
			// Found it. Let's progress args
			args = args2
		}
	}

	s, args2, err = parseString(args)
	if err == nil {
		enumValue, ok := map[string]IsRev{
			"rev": IsRevTrue,
		}[s]
		if ok {
			cmd.Rev = &enumValue
			// Found it. Let's progress args
			args = args2
		}
	}

	// always in the end
	_ = args
	_ = present
	_ = s
	_ = args2
	return nil
}
