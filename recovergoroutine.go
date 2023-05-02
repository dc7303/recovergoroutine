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

			if !checkGoStmt(pass, goStmt) {
				pass.Report(analysis.Diagnostic{
					Pos:      goStmt.Pos(),
					End:      0,
					Category: "goroutine",
					Message:  "goroutine must has recover",
				})
			}

			return false
		})
	}
	return nil, nil
}

func checkGoStmt(pass *analysis.Pass, goStmt *ast.GoStmt) bool {
	fn := goStmt.Call
	safeGoStmt := false
	ast.Inspect(fn, func(n ast.Node) bool {
		if deferStmt, ok := n.(*ast.DeferStmt); ok {
			safeGoStmt = checkHasRecover(deferStmt)
			return false
		}

		if c, ok := n.(*ast.CallExpr); ok {
			if ident, ok := c.Fun.(*ast.Ident); ok {
				if ident.Name == "recover" {
					safeGoStmt = true
					return false
				}
			}
		}
		return true
	})

	return safeGoStmt
}

func checkHasRecover(deferStmt *ast.DeferStmt) bool {
	fn := deferStmt.Call

	hasRecover := false
	ast.Inspect(fn, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if ident, ok := c.Fun.(*ast.Ident); ok && ident.Name == "recover" {
				hasRecover = true
				return false
			}
		}

		return true
	})

	return hasRecover
}
