package commandparser

type Scan struct {
	Cursor int
	Match  *string `rname:"match"`
	Count  *int    `rname:"count"`
	Type   *string `rname:"type"`
}

type HSet struct {
	Key        string
	FieldValue []FieldValue
}
