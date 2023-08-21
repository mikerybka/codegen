package codegen

var Typescript = LanguageSpec{
	ImportMap: map[string]string{},
	Syntax:    TypescriptSyntax,
}

var TypescriptSyntax = SyntaxSpec{
	BlockStyle: BlockStyleBraces,
	Casing: CasingSpec{
		Variable: CasingCamel,
		Constant: CasingScreamingSnake,
		Function: CasingCamel,
		Type:     CasingPascal,
	},
	Typed: true,
}
