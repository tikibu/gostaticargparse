package commandparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseScan(t *testing.T) {
	cmd_strings := []string{"123", "match", "test*", "count", "10", "type", "string"}
	var cmd_bytes [][]byte
	for _, s := range cmd_strings {
		cmd_bytes = append(cmd_bytes, []byte(s))
	}
	cmd := Scan{}
	err := cmd.ParseCommand(cmd_bytes)

	assert.NoError(t, err)
	assert.Equal(t, int(123), cmd.Cursor)

	assert.NotNil(t, cmd.Match)
	assert.Equal(t, "test*", *cmd.Match)

	assert.NotNil(t, cmd.Count)
	assert.Equal(t, int(10), *cmd.Count)

	assert.NotNil(t, cmd.Type)
	assert.Equal(t, "string", *cmd.Type)

}

func TestParseScan2(t *testing.T) {
	cmd_strings := []string{"123", "match", "test*", "count", "10"}
	var cmd_bytes [][]byte
	for _, s := range cmd_strings {
		cmd_bytes = append(cmd_bytes, []byte(s))
	}
	cmd := Scan{}
	err := cmd.ParseCommand(cmd_bytes)

	assert.NoError(t, err)
	assert.Equal(t, int(123), cmd.Cursor)

	assert.NotNil(t, cmd.Match)
	assert.Equal(t, "test*", *cmd.Match)

	assert.NotNil(t, cmd.Count)
	assert.Equal(t, int(10), *cmd.Count)

}

func TestParseScan3(t *testing.T) {
	cmd_strings := []string{"some_string", "match", "test*", "count", "10"}
	var cmd_bytes [][]byte
	for _, s := range cmd_strings {
		cmd_bytes = append(cmd_bytes, []byte(s))
	}
	cmd := Scan{}
	err := cmd.ParseCommand(cmd_bytes)

	assert.Error(t, ErrParsingInt, err)

}
