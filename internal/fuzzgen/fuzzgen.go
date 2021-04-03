package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	err := generateFuzzers()
	if err != nil {
		log.Fatal(err)
	}
}

func generateFuzzers() error {
	rdrs, err := rdrFuncs()
	if err != nil {
		return err
	}
	out, err := os.Create("fuzzers_gen.go")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(out, `// Code generated by gen/fuzzgen. DO NOT EDIT.

package rjson

import (
	"encoding/json"
	"bytes"
)

var generatedFuzzers = []fuzzer{`)

	if err != nil {
		return err
	}

	for _, rdr := range rdrs {
		_, err = fmt.Fprintf(out, `	{name: "fuzz%sGen", fn: fuzz%sGen},
`, rdr.FuncName, rdr.FuncName)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(out, `}`)
	if err != nil {
		return err
	}
	for _, rdr := range rdrs {
		err = tmpl.Execute(out, rdr)
		if err != nil {
			return err
		}
	}
	return nil
}

var tmpl = template.Must(template.New("").Parse(`
func fuzz{{ .FuncName }}Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want {{ .WantType }}
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := {{ .FuncName }}(data{{ .ExtraParams }})
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal := stdLibCompatibleValue(got)
	err = fuzzCompare(want, gotVal)
	return 0, err
}
`))

type rdrFunc struct {
	FuncName    string
	ValType     string
	ExtraParams string
	GotValCheck string
	WantType    string
}

func rdrFuncs() ([]rdrFunc, error) {
	var result []rdrFunc
	node, err := parser.ParseFile(token.NewFileSet(), "readers.go", nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok ||
			!strings.HasPrefix(fn.Name.Name, "Read") ||
			fn.Recv != nil {
			continue
		}
		name := fn.Name.Name
		if len(fn.Type.Results.List) != 3 && name != "ReadInt" {
			continue
		}
		res0Type, err := tpString(fn.Type.Results.List[0].Type)
		if err != nil {
			return nil, err
		}
		rf := rdrFunc{
			FuncName:    name,
			ValType:     res0Type,
			GotValCheck: "got",
			WantType:    res0Type,
		}
		if fn.Type.Params.NumFields() == 2 {
			rf.ExtraParams = ", nil"
		}
		if rf.FuncName == "ReadStringBytes" {
			rf.WantType = "string"
			rf.GotValCheck = "string(got)"
		}
		result = append(result, rf)
	}
	return result, nil
}

func tpString(expr ast.Expr) (string, error) {
	switch tp := expr.(type) {
	case *ast.Ident:
		return tp.String(), nil
	case *ast.MapType:
		ks, err := tpString(tp.Key)
		if err != nil {
			return "", err
		}
		vs, err := tpString(tp.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`map[%s]%s`, ks, vs), nil
	case *ast.ArrayType:
		elt, err := tpString(tp.Elt)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`[]%s`, elt), nil
	case *ast.InterfaceType:
		return "interface{}", nil
	}
	return "", fmt.Errorf("unexpected type: %T", expr)
}
