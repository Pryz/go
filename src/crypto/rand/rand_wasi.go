// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package rand

import "syscall"

func init() {
	Reader = &reader{}
}

type reader struct{}

func (r *reader) Read(b []byte) (int, error) {
	err := syscall.RandomGet(b)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
