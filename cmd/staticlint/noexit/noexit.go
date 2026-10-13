// Package noexit implements a custom analyzer for the multichecker that
// forbids a direct call to os.Exit inside the main function of the main
// package.
package noexit

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

// Analyzer reports direct calls to os.Exit in the main function of the main
// package.
var Analyzer = &analysis.Analyzer{
	Name:     "noexit",
	Doc:      "forbids calling os.Exit in the main function of the main package",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	if pass.Pkg.Name() != "main" {
		return nil, nil
	}

	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		fn := n.(*ast.FuncDecl)

		if !isProjectFile(pass, fn.Name.Pos()) {
			return
		}

		if fn.Name.Name != "main" || fn.Body == nil {
			return
		}

		ast.Inspect(fn.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			callee := typeutil.Callee(pass.TypesInfo, call)
			if callee == nil {
				return true
			}

			sig, ok := callee.Type().(*types.Signature)
			if !ok {
				return true
			}

			recv := sig.Recv()
			if recv != nil {
				return true
			}

			obj := callee
			if objName := obj.Name(); objName != "Exit" {
				return true
			}

			if pkg := obj.Pkg(); pkg == nil || pkg.Name() != "os" {
				return true
			}

			pass.Reportf(call.Pos(), "direct call to os.Exit in main")
			return true
		})
	})

	return nil, nil
}

func isProjectFile(pass *analysis.Pass, pos token.Pos) bool {
	path := pass.Fset.Position(pos).Filename
	if path == "" {
		return false
	}

	// Synthetic main packages generated for the test build have no .go
	// extension and live in the build cache. Real project sources always
	// end with .go.
	return strings.HasSuffix(path, ".go")
}
