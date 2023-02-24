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

package cmdexecutor

import (
	"fmt"
	"os/exec"
)

type Step struct {
	cmd           *exec.Cmd
	id            int
	name          string
	args          []string
	started       bool        // The step started?
	finished      bool        // The step executed finished?
	killed        bool        // The step killed?
	success       bool        //The step executed and not return an error?
	passed        bool        // The step execute passed ok? go next step
	passCheckFunc func() bool // The step pass check function
}

func NewStep(id int, name string, args []string, passCheckFunc func() bool) *Step {
	return &Step{id: id, name: name, args: args, passCheckFunc: passCheckFunc}
}
func (s *Step) logPrefix() string { return fmt.Sprintf("[%d,%s \"%s\"]", s.id, s.cmd, s.args) }

func (s *Step) kill() (err error) {
	if c := s.cmd; c != nil {
		if p := c.Process; p != nil {
			err = p.Kill()
		}
	}
	return
}
