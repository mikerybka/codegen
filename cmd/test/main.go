package main

import (
	"go/parser"
	"go/token"
	"os"

	"gihtub.com/mikerybka/codegen"
)

func main() {
	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, "cmd/test/data/input/codegentest.go", "", parser.ParseComments)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("cmd/test/data/output/codegentest.rb")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = codegen.Ruby.WriteFile(f, fileAST)

}
