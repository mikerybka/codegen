package codegen

var Java = LanguageSpec{
	ImportMap: map[string]string{},
	Syntax:    JavaSyntax,
}

var JavaSyntax = SyntaxSpec{
	BlockStyle: BlockStyleBraces,
	Casing: CasingSpec{
		Variable: CasingCamel,
		Constant: CasingScreamingSnake,
		Function: CasingCamel,
		Type:     CasingPascal,
	},
	Typed: true,
}
