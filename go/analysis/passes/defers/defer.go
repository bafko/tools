// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package defers

import (
	_ "embed"
	"go/ast"
	"go/types"

	"github.com/bafko/tools/go/analysis"
	"github.com/bafko/tools/go/analysis/passes/inspect"
	"github.com/bafko/tools/go/analysis/passes/internal/analysisutil"
	"github.com/bafko/tools/go/ast/inspector"
	"github.com/bafko/tools/go/types/typeutil"
)

//go:embed doc.go
var doc string

// Analyzer is the defer analyzer.
var Analyzer = &analysis.Analyzer{
	Name:     "defer",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Doc:      analysisutil.MustExtractDoc(doc, "defer"),
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if !analysisutil.Imports(pass.Pkg, "time") {
		return nil, nil
	}

	checkDeferCall := func(node ast.Node) bool {
		switch v := node.(type) {
		case *ast.CallExpr:
			fn, ok := typeutil.Callee(pass.TypesInfo, v).(*types.Func)
			if ok && fn.Name() == "Since" && fn.Pkg().Path() == "time" {
				pass.Reportf(v.Pos(), "call to time.Since is not deferred")
			}
		case *ast.FuncLit:
			return false // prune
		}
		return true
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.DeferStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		d := n.(*ast.DeferStmt)
		ast.Inspect(d.Call, checkDeferCall)
	})

	return nil, nil
}
