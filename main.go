package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var src string
	if len(args) != 0 {
		srcFile := args[0]
		srcB, err := os.ReadFile(srcFile)
		if err != nil {
			log.Fatal(err)
		}
		src = string(srcB)
	} else {
		src = `package main

func MerryXMas() {
    var hoge = 2
}`
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.Mode(0))

	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, f)

}
