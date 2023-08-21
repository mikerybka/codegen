package codegen

import (
	"fmt"
	"go/ast"
	"io"
)

type LanguageSpec struct {
	ImportMap map[string]string
	Syntax    SyntaxSpec
	Build     BuildSpec
}

func (lang *LanguageSpec) WriteFile(w io.Writer, f *ast.File) error {
	return fmt.Errorf("not implemented: codegen.LanguageSpec.WriteFile")
}

type SyntaxSpec struct {
	BlockStyle        BlockStyle
	Casing            CasingSpec
	Typed             bool
	CommentLinePrefix string
	Indent            string
}

type BuildSpec struct {
	Install          []string
	Compile          []string
	Run              []string
	InstallPackage   []string
	InitializeModule []string
}

type CasingSpec struct {
	Variable Casing
	Constant Casing
	Function Casing
	Type     Casing
}

type Casing int

const (
	CasingPascal = iota
	CasingCamel
	CasingSnake
	CasingKebab
	CasingScreamingSnake
)

type BlockStyle int

const (
	BlockStyleBraces BlockStyle = iota
	BlockStylePython
	BlockStyleRuby
)
