package codegen

var Python = LanguageSpec{
	ImportMap: map[string]string{},
	Syntax:    PythonSyntax,
}

var PythonSyntax = SyntaxSpec{
	BlockStyle: BlockStylePython,
	Casing: CasingSpec{
		Variable: CasingSnake,
		Constant: CasingScreamingSnake,
		Function: CasingSnake,
		Type:     CasingPascal,
	},
	Typed: false,
}
