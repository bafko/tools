// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package checker_test

import (
	"go/ast"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/bafko/tools/go/analysis"
	"github.com/bafko/tools/go/analysis/analysistest"
	"github.com/bafko/tools/go/analysis/internal/checker"
	"github.com/bafko/tools/go/analysis/passes/inspect"
	"github.com/bafko/tools/go/ast/inspector"
	"github.com/bafko/tools/internal/testenv"
)

// TestStartFixes make sure modifying the first character
// of the file takes effect.
func TestStartFixes(t *testing.T) {
	testenv.NeedsGoPackages(t)

	files := map[string]string{
		"comment/doc.go": `/* Package comment */
package comment
`}

	want := `// Package comment
package comment
`

	testdata, cleanup, err := analysistest.WriteFiles(files)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(testdata, "src/comment/doc.go")
	checker.Fix = true
	checker.Run([]string{"file=" + path}, []*analysis.Analyzer{commentAnalyzer})

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	got := string(contents)
	if got != want {
		t.Errorf("contents of rewritten file\ngot: %s\nwant: %s", got, want)
	}

	defer cleanup()
}

var commentAnalyzer = &analysis.Analyzer{
	Name:     "comment",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      commentRun,
}

func commentRun(pass *analysis.Pass) (interface{}, error) {
	const (
		from = "/* Package comment */"
		to   = "// Package comment"
	)
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(nil, func(n ast.Node) {
		if n, ok := n.(*ast.Comment); ok && n.Text == from {
			pass.Report(analysis.Diagnostic{
				Pos: n.Pos(),
				End: n.End(),
				SuggestedFixes: []analysis.SuggestedFix{{
					TextEdits: []analysis.TextEdit{{
						Pos:     n.Pos(),
						End:     n.End(),
						NewText: []byte(to),
					}},
				}},
			})
		}
	})

	return nil, nil
}
