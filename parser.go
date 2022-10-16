package commandparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
)

type EnumValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Enum struct {
	Type   string            `json:"type"`
	Values map[string]string `json:"values"`
}

func collectEnums(node ast.Node, enums map[string]Enum) {
	collectEnums := func(n ast.Node) bool {
		var t ast.Expr
		var structName string

		switch n := n.(type) {
		case *ast.TypeSpec:
			typeSpec := n
			if typeSpec.Type == nil {
				return true
			}

			structName = typeSpec.Name.Name
			t = typeSpec.Type

			x2, ok := t.(*ast.Ident)
			if ok {
				enumVal := enums[structName]
				enumVal.Type = x2.Name
				enums[structName] = enumVal
			}
			return true
		case *ast.ValueSpec:
			valSpec := n
			if len(valSpec.Values) == 0 {
				return true
			}
			val := valSpec.Values[0]
			basicLit, ok := val.(*ast.BasicLit)
			if !ok {
				return true
			}
			if len(valSpec.Names) == 0 {
				return true
			}
			// we only process enum values with
			// explicit type specification
			// so  ABC SomeEnumType = "abc"
			// but not ABC = "abc"
			if valSpec.Type == nil {
				return true
			}
			valSpecType, ok := valSpec.Type.(*ast.Ident)
			if !ok {
				return true
			}
			enumVal := enums[valSpecType.Name]
			if enumVal.Values == nil {
				enumVal.Values = make(map[string]string, 0)
			}
			switch basicLit.Kind {
			case token.INT:
				enumVal.Values[basicLit.Value] = valSpec.Names[0].Name
			case token.STRING:
				enumVal.Values[strings.Trim(basicLit.Value, `"`)] = valSpec.Names[0].Name
			default:
				return true
			}
			enums[valSpecType.Name] = enumVal

		default:
			return true
		}

		return true
	}

	ast.Inspect(node, collectEnums)
}

func collectStructs(node ast.Node) map[string]*ast.StructType {
	structs := make(map[string]*ast.StructType, 0)
	collectStructs := func(n ast.Node) bool {
		var t ast.Expr
		var structName string

		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if typeSpec.Type == nil {
			return true
		}

		structName = typeSpec.Name.Name
		t = typeSpec.Type

		x, ok := t.(*ast.StructType)
		if !ok {
			return true
		}

		structs[structName] = x
		return true
	}

	ast.Inspect(node, collectStructs)
	return structs
}

type StructField struct {
	Name        string `json:"name"`
	IsStar      bool   `json:"is_star"`
	IsArray     bool   `json:"is_array"`
	Type        string `json:"type"`
	NameFromTag string `json:"name_from_tag"`
}

type Struct []StructField

type Structs map[string]Struct

func (structs Structs) processStruct(name string, s *ast.StructType) {
	fields := make(Struct, 0)
	for _, field := range s.Fields.List {
		if field.Names == nil {
			// anonymous field
			continue
		}
		if len(field.Names) == 0 {
			continue
		}

		// So we want to support something like this:
		// ABC *[]string
		// so pointer - array - string
		tp := field.Type
		starExpr, isStar := tp.(*ast.StarExpr)
		if isStar {
			tp = starExpr.X
		}

		arrayType, isArray := tp.(*ast.ArrayType)
		if isArray {
			tp = arrayType.Elt
		}

		fieldName := field.Names[0].Name

		tag := field.Tag
		rname := ""
		if tag != nil {
			rname = ExtractRnameFromTag(tag.Value)
		}

		typeString := types.ExprString(tp)
		fields = append(fields, StructField{fieldName, isStar, isArray, typeString, rname})
	}

	structs[name] = fields

	fmt.Println()
}

type ParserGenerator struct {
	// A list of struct names for each of the parsed files
	FileStructs map[string][]string

	// Enums. No need to care what file holds them
	Enums map[string]Enum

	// Structs
	Structs Structs

	//Package
	Package string
}

func (p *ParserGenerator) ParseFile(fname string, source []byte) (err error) {
	if p.FileStructs == nil {
		p.FileStructs = make(map[string][]string, 0)
	}

	if p.Enums == nil {
		p.Enums = make(map[string]Enum, 0)
	}

	if p.Structs == nil {
		p.Structs = make(map[string]Struct, 0)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fname, source, 0)
	if err != nil {
		return err
	}
	p.Package = file.Name.Name
	fileStructs := collectStructs(file)
	for name, s := range fileStructs {
		p.Structs.processStruct(name, s)
		p.FileStructs[fname] = append(p.FileStructs[fname], name)
	}

	collectEnums(file, p.Enums)

	return nil
}
