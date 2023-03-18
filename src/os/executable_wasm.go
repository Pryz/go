// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js || wasip1

package os

import (
	"errors"
	"runtime"
)

func executable() (string, error) {
	return "", errors.New("Executable not implemented for " + runtime.GOOS)
}
