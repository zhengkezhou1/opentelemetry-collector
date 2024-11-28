// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
)

func TestLine_MoveTo(t *testing.T) {
	ms := generateTestLine()
	dest := NewLine()
	ms.MoveTo(dest)
	assert.Equal(t, NewLine(), ms)
	assert.Equal(t, generateTestLine(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newLine(&otlpprofiles.Line{}, &sharedState)) })
	assert.Panics(t, func() { newLine(&otlpprofiles.Line{}, &sharedState).MoveTo(dest) })
}

func TestLine_CopyTo(t *testing.T) {
	ms := NewLine()
	orig := NewLine()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestLine()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newLine(&otlpprofiles.Line{}, &sharedState)) })
}

func TestLine_FunctionIndex(t *testing.T) {
	ms := NewLine()
	assert.Equal(t, int32(0), ms.FunctionIndex())
	ms.SetFunctionIndex(int32(1))
	assert.Equal(t, int32(1), ms.FunctionIndex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLine(&otlpprofiles.Line{}, &sharedState).SetFunctionIndex(int32(1)) })
}

func TestLine_Line(t *testing.T) {
	ms := NewLine()
	assert.Equal(t, int64(0), ms.Line())
	ms.SetLine(int64(1))
	assert.Equal(t, int64(1), ms.Line())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLine(&otlpprofiles.Line{}, &sharedState).SetLine(int64(1)) })
}

func TestLine_Column(t *testing.T) {
	ms := NewLine()
	assert.Equal(t, int64(0), ms.Column())
	ms.SetColumn(int64(1))
	assert.Equal(t, int64(1), ms.Column())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newLine(&otlpprofiles.Line{}, &sharedState).SetColumn(int64(1)) })
}

func generateTestLine() Line {
	tv := NewLine()
	fillTestLine(tv)
	return tv
}

func fillTestLine(tv Line) {
	tv.orig.FunctionIndex = int32(1)
	tv.orig.Line = int64(1)
	tv.orig.Column = int64(1)
}
