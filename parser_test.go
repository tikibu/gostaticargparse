package commandparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Unmarshaller interface {
	Unmarshal(cmd []string) (rest []string, err error)
}

func TestParseFiles(t *testing.T) {
	path, err := os.Getwd()
	assert.NoError(t, err)

	// join paths path and "parser.go"
	testilepath := filepath.Join(path, "p1.go")
	source, err := ioutil.ReadFile(testilepath)
	assert.NoError(t, err)

	pg := ParserGenerator{}
	err = pg.ParseFile(testilepath, source)
	assert.NoError(t, err)

	testilepath = filepath.Join(path, "p2.go")
	source, err = ioutil.ReadFile(testilepath)
	assert.NoError(t, err)

	err = pg.ParseFile(testilepath, source)
	assert.NoError(t, err)

	b, err := json.MarshalIndent(pg, "", "  ")
	assert.NoError(t, err)
	fmt.Println("ALL\n=================================\n", string(b))
	//	fmt.Println(fileStructs}
}
