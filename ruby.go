package codegen

var Ruby = LanguageSpec{
	ImportMap: map[string]string{},
	Syntax:    RubySyntax,
}

var RubySyntax = SyntaxSpec{
	BlockStyle: BlockStyleRuby,
	Casing: CasingSpec{
		Variable: CasingSnake,
		Constant: CasingScreamingSnake,
		Function: CasingSnake,
		Type:     CasingPascal,
	},
	Typed: false,
}
