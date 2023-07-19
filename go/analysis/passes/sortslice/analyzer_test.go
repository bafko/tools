// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sortslice_test

import (
	"testing"

	"github.com/bafko/tools/go/analysis/analysistest"
	"github.com/bafko/tools/go/analysis/passes/sortslice"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, sortslice.Analyzer, "a")
}
