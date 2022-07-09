package dbrloadcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	doc           = "dbr helper linter: Check loading on invalid pointer"
	targetPackage = "repositories"
)

var Analyzer = &analysis.Analyzer{
	Name: "dbrloadcheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	// repositoriesのフォルダのみ
	// Load, LoadContext, LoadOneに渡す変数がポインタ型かどうかをチェックする
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			if n.Name.Name != targetPackage {
				return
			}
			for _, decl := range n.Decls {
				if fnDecl, ok := decl.(*ast.FuncDecl); ok {
					analyze(pass, fnDecl)
				}
			}
		}
	})

	return nil, nil
}

func analyze(pass *analysis.Pass, fnDecl *ast.FuncDecl) {}
