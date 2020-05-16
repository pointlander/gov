// Copyright 2020 The GoV Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// https://stackoverflow.com/questions/52242077/go-modules-finding-out-right-pseudo-version-vx-y-z-timestamp-commit-of-re
func main() {
	command := exec.Command("git", "--no-pager", "show", "--quiet",
		"--abbrev=12", "--date=format-local:%Y%m%d%H%M%S", "--format=%cd-%h")
	command.Env = append(command.Env, "TZ=UTC")
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		panic(err)
	}
	version := strings.TrimSpace(string(output))
	fmt.Printf("v0.0.0-%s\n", version)
}
