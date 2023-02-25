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

package cmd

import (
	"bufio"
	"bytes"
	"os/exec"
)

func Run(name string, args ...string) (outStr, errStr string, err error) {
	c := exec.Command(name, args...)
	var (
		ob bytes.Buffer
		eb bytes.Buffer
	)
	c.Stdout = &ob
	c.Stderr = &eb
	if err = c.Run(); err != nil {
		return
	}
	outStr = ob.String()
	errStr = eb.String()
	return
}

func Start(name string, args []string, outC, errC chan<- string, inC <-chan string, startedFunc func(outCmd *exec.Cmd)) (err error) {
	outCmd := exec.Command(name, args...)
	{
		stdin, wErr := outCmd.StdinPipe()
		if wErr != nil {
			err = wErr
			return
		}
		if c := inC; c != nil {
			go func() {
				for {
					iv := <-c
					if _, err2 := stdin.Write([]byte(iv)); err2 != nil {
						break
					}
					if _, err2 := stdin.Write([]byte("\n")); err2 != nil {
						break
					}
				}
			}()
		}
	}
	{
		stdout, wErr := outCmd.StdoutPipe()
		if wErr != nil {
			err = wErr
			return
		}
		if c := outC; c != nil {
			rr := bufio.NewReader(stdout)
			go func() {
				defer close(c)
				for {
					readLine, _, err2 := rr.ReadLine()
					if err2 != nil {
						break
					}
					c <- string(readLine)
				}
			}()
		}
	}
	{
		stderr, wErr := outCmd.StderrPipe()
		if wErr != nil {
			err = wErr
			return
		}
		if c := errC; c != nil {
			rr := bufio.NewReader(stderr)
			go func() {
				defer close(c)
				for {
					readLine, _, err2 := rr.ReadLine()
					if err2 != nil {
						break
					}
					c <- string(readLine)
				}
			}()
		}
	}
	if err = outCmd.Start(); err != nil {
		return
	}
	if fn := startedFunc; fn != nil {
		fn(outCmd)
	}
	err = outCmd.Wait()
	return
}
