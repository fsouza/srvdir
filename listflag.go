// Copyright 2018 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "strings"

type listflag struct {
	v []string
}

func (f *listflag) String() string {
	return strings.Join(f.v, "")
}

func (f *listflag) Set(v string) error {
	f.v = append(f.v, v)
	return nil
}
