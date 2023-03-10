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
	"github.com/whrwsoftware/panelbase/cmd"
	"os/exec"
	"time"
)

type Step struct {
	cmd        *exec.Cmd
	id         int
	name       string
	args       []string
	Started    bool          // The step started?
	Exited     bool          // The step exited?
	ExitCode   int           // The step exit code
	Success    bool          //The step executed and not return an error?
	State      string        // The step state
	Pid        int           // The step process id
	UserTime   time.Duration // The step user time
	SystemTime time.Duration // The step system time
}

func NewStep(id int, name string, args ...string) *Step {
	return &Step{id: id, name: name, args: args}
}

func (s *Step) Id() int           { return s.id }
func (s *Step) CmdName() string   { return s.name }
func (s *Step) CmdArgs() []string { return s.args }

func (s *Step) Start(outC, errC chan<- string) (ok bool, err error) {
	ok, err = cmd.Start(s.name, s.args, outC, errC, nil, func(outCmd *exec.Cmd) {
		if outCmd != nil {
			s.cmd = outCmd
			s.Started = true
		}
	})
	if c := s.cmd; c != nil {
		if ps := c.ProcessState; ps != nil {
			s.Exited = ps.Exited()
			s.ExitCode = ps.ExitCode()
			s.Success = ps.Success()
			s.State = ps.String()
			s.UserTime = ps.UserTime()
			s.SystemTime = ps.SystemTime()
		}
	}
	return
}

func (s *Step) Kill() (err error) {
	if c := s.cmd; c != nil {
		if p := c.Process; p != nil {
			err = p.Kill()
		}
	}
	return
}
