package commandparser

func (cmd *Scan) ParseCommand(args [][]byte) (err error) {
	// used in every command
	var present bool

	cmd.Cursor, args, err = parseInt(args)
	if err != nil {
		return err
	}

	present, args, err = optionalPresent(args, "match")
	if err != nil {
		return err
	}

	if present {
		var s string
		s, args, err = parseString(args)
		if err != nil {
			return err
		}
		cmd.Match = &s
	}

	present, args, err = optionalPresent(args, "count")
	if err != nil {
		return err
	}
	if present {
		var i int
		i, args, err = parseInt(args)
		if err != nil {
			return err
		}
		cmd.Count = &i
	}

	present, args, err = optionalPresent(args, "type")
	if err != nil {
		return err
	}
	if present {
		var s string
		s, args, err = parseString(args)
		if err != nil {
			return err
		}
		cmd.Type = &s
	}

	// always in the end
	_ = args
	_ = present
	return nil
}

func (cmd *FieldValue) ParseCommand(args [][]byte) (args_reminder [][]byte, err error) {
	// used in every command
	var present bool

	cmd.Field, args, err = parseString(args)
	if err != nil {
		return args, err
	}

	cmd.Value, args, err = parseString(args)
	if err != nil {
		return args, err
	}

	// always in the end
	_ = args
	_ = present
	return args, nil
}

func (cmd *HSet) ParseCommand(args [][]byte) (err error) {
	// used in every command
	var present bool

	cmd.Key, args, err = parseString(args)
	if err != nil {
		return err
	}

	for {
		var fieldValue FieldValue
		args, err = fieldValue.ParseCommand(args)
		if err == ErrNotEnoughArgs {
			break
		}
		if err != nil {
			return err
		}
		cmd.FieldValue = append(cmd.FieldValue, fieldValue)
	}

	if len(cmd.FieldValue) == 0 {
		return ErrWrongNumberOfArgs
	}

	// always in the end
	_ = args
	_ = present
	return nil
}
