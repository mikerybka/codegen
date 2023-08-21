package codegen

var Go = LanguageSpec{
	ImportMap: map[string]string{},
	Syntax:    GoSyntax,
	Build:     GoBuild,
}

var GoSyntax = SyntaxSpec{
	BlockStyle: BlockStyleBraces,
	Casing: CasingSpec{
		Variable: CasingCamel,
		Constant: CasingScreamingSnake,
		Function: CasingCamel,
		Type:     CasingPascal,
	},
	Typed:             true,
	CommentLinePrefix: "//",
	Indent:            "\t",
}

var GoBuild = BuildSpec{
	Install: []string{
		"wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz",
		"rm -rf /usr/local/go",
		"tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz",
		"rm go1.21.0.linux-amd64.tar.gz",
	},
	Compile: []string{
		"go build -o {{.Output}} {{.Input}}",
	},
	Run: []string{
		"go run {{.MainFile}}",
	},
	InstallPackage: []string{
		"go get {{.Package}}",
	},
}
