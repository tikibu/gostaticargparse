package commandparser

type SortBy string

const (
	SortByScore SortBy = "BYSCORE"
	SortByLex   SortBy = "BYLEX"
)

type ABC int

const (
	ABC1 ABC = 1
	ABC2 ABC = 2
)

type IsRev string

const (
	IsRevTrue IsRev = "rev"
)

type IsWithScores string

const (
	IsWithScoresTrue IsWithScores = "withscores"
)

type Limit struct {
	Offset int
	Count  int
}

type MGET struct {
	Keys []string
}

//ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
type ZRange struct {
	Key        string
	Start      int
	Stop       int
	SortBy     *SortBy // Optional enum doesn't need a name
	Rev        *IsRev  // This is how we parse a boolean flag
	Limit      *Limit  `rname:"limit" json:"limit,omitempty"`
	WithScores *IsWithScores
}

type FieldValue struct {
	Field string
	Value string
}
