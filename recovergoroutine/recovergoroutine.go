package recovergoroutine

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "recovergoroutine",
	Doc:  "finds goroutine code without recover",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			goStmt, ok := n.(*ast.GoStmt)
			if !ok {
				return true
			}

			if safeGoStmt(goStmt) {
				return true
			}

			pass.Report(analysis.Diagnostic{
				Pos:      goStmt.Pos(),
				End:      0,
				Category: "goroutine",
				Message:  "goroutine must have recover",
			})

			return false
		})
	}
	return nil, nil
}

func safeGoStmt(goStmt *ast.GoStmt) bool {
	fn := goStmt.Call
	safeGoStmt := false
	ast.Inspect(fn, func(n ast.Node) bool {
		deferStmt, ok := n.(*ast.DeferStmt)
		if !ok {
			return true
		}

		callExpr := deferStmt.Call
		if isRecover(callExpr) {
			safeGoStmt = true
			return false
		}

		ident, ok := callExpr.Fun.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Obj == nil {
			return true
		}

		funcDecl, ok := ident.Obj.Decl.(*ast.FuncDecl)
		if !ok {
			return true
		}

		ast.Inspect(funcDecl, func(node ast.Node) bool {
			if callExpr, ok := node.(*ast.CallExpr); ok && isRecover(callExpr) {
				safeGoStmt = true
				return false
			}
			return true
		})
		return true
	})

	return safeGoStmt
}

func isRecover(callExpr *ast.CallExpr) bool {
	ident, ok := callExpr.Fun.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == "recover"
}
