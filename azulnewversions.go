// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"go/ast"
	"strconv"
)

func init() {
	register(azulnewversionsFix)
}

var azulnewversionsFix = fix{
	"azulnewversions",
	"2014-7-10",
	azulnewversions,
	`Updates import paths from old Google Code paths to new GitHub ones.

http://azul3d.org/news/2014/important-import-paths-have-changed.html
`,
}

// rewriteImportAndName rewrites any import of path oldPath to path newPath and
// aliases the import as name. For example given the import statement:
//
//  import "azul3d.org/v1/math"
//
// rewriteImport("azul3d.org/v1/math", "azul3d.org/lmath.v1", "math") would
// write:
//
//  import math "azul3d.org/v1/lmath"
//
// Or with an import statement already aliased:
//
//  import sheep "azul3d.org/v1/math"
//
// rewriteImport("azul3d.org/v1/math", "azul3d.org/lmath.v1", "math") would
// write:
//
//  import sheep "azul3d.org/v1/lmath"
//
func rewriteImportAndName(f *ast.File, oldPath, newPath, name string) (rewrote bool) {
	for _, imp := range f.Imports {
		if importPath(imp) == oldPath {
			rewrote = true
			// record old End, because the default is to compute
			// it using the length of imp.Path.Value.
			imp.EndPos = imp.End()
			if imp.Name == nil {
				imp.Name = ast.NewIdent(name)
			}
			imp.Path.Value = strconv.Quote(newPath)
		}
	}
	return
}

func azulnewversions(f *ast.File) bool {
	var fixed bool

	// audio package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/audio", "azul3d.org/audio.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/audio", "azul3d.org/audio.v1")

	// audio/wav package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/audio/wav", "azul3d.org/audio/wav.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/audio/wav", "azul3d.org/audio/wav.v1")

	// chippy package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/chippy", "azul3d.org/chippy.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/chippy", "azul3d.org/chippy.v1")

	// chippy win32 internal package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/chippy/wrappers/win32", "azul3d.org/chippy.v1/internal/win32")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/chippy/wrappers/win32", "azul3d.org/chippy.v1/internal/win32")

	// chippy x11 internal package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/chippy/wrappers/x11", "azul3d.org/chippy.v1/internal/x11")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/chippy/wrappers/x11", "azul3d.org/chippy.v1/internal/x11")

	// clock package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/clock", "azul3d.org/clock.v0")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/clock", "azul3d.org/clock.v1")

	// gfx package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/gfx", "azul3d.org/gfx.v1")

	// gfx/gl2 package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/gfx/gl2", "azul3d.org/gfx/gl2.v1")

	// gfx/window package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/gfx/window", "azul3d.org/gfx/window.v1")

	// keyboard package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/chippy/keyboard", "azul3d.org/keyboard.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/keyboard", "azul3d.org/keyboard.v1")

	// mouse package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/chippy/mouse", "azul3d.org/mouse.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/mouse", "azul3d.org/mouse.v1")

	// math package.
	fixed = fixed || rewriteImportAndName(f, "azul3d.org/v0/math", "azul3d.org/lmath.v1", "math")
	fixed = fixed || rewriteImportAndName(f, "azul3d.org/v1/math", "azul3d.org/lmath.v1", "math")

	// native/gl package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/native/gl", "azul3d.org/native/gl.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/native/gl", "azul3d.org/native/gl.v1")

	// native/gles1 package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/native/gles1", "azul3d.org/native/gles1.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/native/gles1", "azul3d.org/native/gles1.v1")

	// native/gles2 package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/native/gles2", "azul3d.org/native/gles2.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/native/gles2", "azul3d.org/native/gles2.v1")

	// native/al package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/native/al", "azul3d.org/native/al.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/native/al", "azul3d.org/native/al.v1")

	// native/freetype package.
	fixed = fixed || rewriteImport(f, "azul3d.org/v0/native/freetype", "azul3d.org/native/freetype.v1")
	fixed = fixed || rewriteImport(f, "azul3d.org/v1/native/freetype", "azul3d.org/native/freetype.v1")

	if fixed {
		for fixed {
			fixed = azulnewversions(f)
		}
		return true
	}
	return fixed
}
