package commandparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractParserSingle(t *testing.T) {

	// Test ExtractParser memory
	rname := ExtractRnameFromTag("`rname:\"match\"`")

	assert.Equal(t, "match", rname)
}

func TestExtractParser(t *testing.T) {

	// Test ExtractParser memory
	rname := ExtractRnameFromTag("`rname:\"limit\" json:\"limit,omitempty\"`")

	assert.Equal(t, "limit", rname)
}

func TestExtractParserSlightlyWrong(t *testing.T) {

	// Test ExtractParser memory
	rname := ExtractRnameFromTag("`rname:limit json:\"limit, omitempty\"`")

	assert.Equal(t, "limit", rname)
}

func TestExtractParserWrong(t *testing.T) {

	// Test ExtractParser memory
	rname := ExtractRnameFromTag("`rname: limit json:\"limit, omitempty\"`")

	assert.Equal(t, "", rname)
}
