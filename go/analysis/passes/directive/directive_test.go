// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package directive_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/bafko/tools/go/analysis"
	"github.com/bafko/tools/go/analysis/analysistest"
	"github.com/bafko/tools/go/analysis/passes/directive"
)

func Test(t *testing.T) {
	if strings.HasPrefix(runtime.Version(), "go1.") && runtime.Version() < "go1.16" {
		t.Skipf("skipping on %v", runtime.Version())
	}
	analyzer := *directive.Analyzer
	analyzer.Run = func(pass *analysis.Pass) (interface{}, error) {
		defer func() {
			// The directive pass is unusual in that it checks the IgnoredFiles.
			// After analysis, add IgnoredFiles to OtherFiles so that
			// the test harness checks for expected diagnostics in those.
			// (The test harness shouldn't do this by default because most
			// passes can't do anything with the IgnoredFiles without type
			// information, which is unavailable because they are ignored.)
			var files []string
			files = append(files, pass.OtherFiles...)
			files = append(files, pass.IgnoredFiles...)
			pass.OtherFiles = files
		}()

		return directive.Analyzer.Run(pass)
	}
	analysistest.Run(t, analysistest.TestData(), &analyzer, "a")
}
