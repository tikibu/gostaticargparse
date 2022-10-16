package commandparser

import "errors"

func GenerateCommands(p *ParserGenerator) (generated_code map[string][]byte, err error) {
	if p == nil {
		return nil, errors.New("nil parser generator")
	}
	generated_code = make(map[string][]byte)
	/*for _, struct := range p.Structs {

	}*/
	return generated_code, nil
}
