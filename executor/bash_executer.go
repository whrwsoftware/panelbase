// Copyright 2023 panelbase Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package executor

import (
	"os"
	"os/exec"
	"strings"
)

type BashExecutor struct {
	Cmd      string
	bashFile string
	logFile  string

	out []byte
	ok  bool
	err error
}

func NewBashExecutor(cmd string) *BashExecutor {
	tempFile, _ := os.CreateTemp("", "bash-executor-script-*.sh")
	if tempFile == nil {
		panic("create temp file error")
	}
	_, _ = tempFile.WriteString(cmd)
	_ = tempFile.Close()
	return &BashExecutor{Cmd: cmd, bashFile: tempFile.Name()}
}

func (b *BashExecutor) BindLog(logFile string) *BashExecutor { b.logFile = logFile; return b }

func (b *BashExecutor) Exec() *BashExecutor {
	b.ok = true
	// bash postfix-uninstall.sh  > aaa.log 2>&1
	arg0Builder := &strings.Builder{}
	arg0Builder.WriteString("/bin/bash ")
	arg0Builder.WriteString(b.bashFile)
	if b.logFile != "" {
		arg0Builder.WriteString(">>")
		arg0Builder.WriteString(b.logFile)
	}
	arg0Builder.WriteString(" 2>&1")
	arg0 := arg0Builder.String()
	arg0Builder.Reset()
	b.err = exec.Command("/bin/bash", "-c", arg0).Run()
	return b
}

func (b *BashExecutor) Run() *BashExecutor {
	b.ok = true
	arg0Builder := &strings.Builder{}
	arg0Builder.WriteString("/bin/bash ")
	arg0Builder.WriteString(b.bashFile)
	arg0Builder.WriteString(" 2>&1")
	arg0 := arg0Builder.String()
	arg0Builder.Reset()
	b.out, b.err = exec.Command("/bin/bash", "-c", arg0).CombinedOutput()
	return b
}

func (b *BashExecutor) Release() (ok bool, err error) {
	_ = os.Remove(b.bashFile)
	return b.ok, b.err
}

func (b *BashExecutor) OutRelease() (v string, ok bool, err error) {
	ok, err = b.Release()
	return string(b.out), ok, err
}
