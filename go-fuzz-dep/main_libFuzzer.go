// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

//go:build gofuzz && gofuzz_libfuzzer
// +build gofuzz,gofuzz_libfuzzer

package gofuzzdep

import (
	"unsafe"

	. "github.com/trailofbits/go-fuzz/go-fuzz-defs"
)

func Initialize(coverTabPtr unsafe.Pointer, coverTabSize uint64) {
	if coverTabSize != CoverSize {
		panic("Incorrect cover tab size")
	}
	CoverTab = (*[CoverSize]byte)(coverTabPtr)
}
