package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
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

func reti(i int) {
	return i
}`
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.Mode(0))

	if err != nil {
		log.Fatal(err)
	}
	insertHello(f)

	prer := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	prer.Fprint(os.Stdout, fset, f)

	// ast.Print(fset, f)

}

func insertHello(f *ast.File) {

	hello := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("fmt"),
				Sel: ast.NewIdent("Println"),
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("Hello, world"),
				},
			},
		}}

	decls := f.Decls
	for _, dec := range decls {
		switch dec := interface{}(dec).(type) {
		case *ast.FuncDecl:
			pp.Printf("%v is function.\n", dec.Name.Name)
			pp.Println(dec)
			dec.Body.List = append([]ast.Stmt{hello}, dec.Body.List...)
			pp.Println(dec)
		case *ast.GenDecl:
			// pp.Printf("%v is not function.\n", dec.Tok)
			// pp.Printf("specs: %v\n", dec.Specs)
		default:
			fmt.Printf("%T\n", dec)
		}
	}
}
