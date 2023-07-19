// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unusedresult_test

import (
	"testing"

	"github.com/bafko/tools/go/analysis/analysistest"
	"github.com/bafko/tools/go/analysis/passes/unusedresult"
	"github.com/bafko/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	funcs := "typeparams/userdefs.MustUse,errors.New,fmt.Errorf,fmt.Sprintf,fmt.Sprint"
	unusedresult.Analyzer.Flags.Set("funcs", funcs)
	tests := []string{"a"}
	if typeparams.Enabled {
		tests = append(tests, "typeparams")
	}
	analysistest.Run(t, testdata, unusedresult.Analyzer, tests...)
}
