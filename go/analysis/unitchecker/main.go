// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This file provides an example command for static checkers
// conforming to the golang.org/x/tools/go/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//	$ go build -o unitchecker main.go
//	$ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"github.com/bafko/tools/go/analysis/unitchecker"

	"github.com/bafko/tools/go/analysis/passes/asmdecl"
	"github.com/bafko/tools/go/analysis/passes/assign"
	"github.com/bafko/tools/go/analysis/passes/atomic"
	"github.com/bafko/tools/go/analysis/passes/bools"
	"github.com/bafko/tools/go/analysis/passes/buildtag"
	"github.com/bafko/tools/go/analysis/passes/cgocall"
	"github.com/bafko/tools/go/analysis/passes/composite"
	"github.com/bafko/tools/go/analysis/passes/copylock"
	"github.com/bafko/tools/go/analysis/passes/directive"
	"github.com/bafko/tools/go/analysis/passes/errorsas"
	"github.com/bafko/tools/go/analysis/passes/framepointer"
	"github.com/bafko/tools/go/analysis/passes/httpresponse"
	"github.com/bafko/tools/go/analysis/passes/ifaceassert"
	"github.com/bafko/tools/go/analysis/passes/loopclosure"
	"github.com/bafko/tools/go/analysis/passes/lostcancel"
	"github.com/bafko/tools/go/analysis/passes/nilfunc"
	"github.com/bafko/tools/go/analysis/passes/printf"
	"github.com/bafko/tools/go/analysis/passes/shift"
	"github.com/bafko/tools/go/analysis/passes/sigchanyzer"
	"github.com/bafko/tools/go/analysis/passes/stdmethods"
	"github.com/bafko/tools/go/analysis/passes/stringintconv"
	"github.com/bafko/tools/go/analysis/passes/structtag"
	"github.com/bafko/tools/go/analysis/passes/testinggoroutine"
	"github.com/bafko/tools/go/analysis/passes/tests"
	"github.com/bafko/tools/go/analysis/passes/timeformat"
	"github.com/bafko/tools/go/analysis/passes/unmarshal"
	"github.com/bafko/tools/go/analysis/passes/unreachable"
	"github.com/bafko/tools/go/analysis/passes/unsafeptr"
	"github.com/bafko/tools/go/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		testinggoroutine.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
