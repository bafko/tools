// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testinggoroutine_test

import (
	"testing"

	"github.com/bafko/tools/go/analysis/analysistest"
	"github.com/bafko/tools/go/analysis/passes/testinggoroutine"
	"github.com/bafko/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	pkgs := []string{"a"}
	if typeparams.Enabled {
		pkgs = append(pkgs, "typeparams")
	}
	analysistest.Run(t, testdata, testinggoroutine.Analyzer, pkgs...)
}
