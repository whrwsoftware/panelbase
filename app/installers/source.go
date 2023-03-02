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

package installers

import (
	"errors"
	"fmt"
	"github.com/whrwsoftware/panelbase/executor"
	"sync"
)

var (
	ErrInstallRunning = errors.New("source_installer: the installation is running, please wait while finished")
)

type source struct {
	mu *sync.Mutex

	Name string

	InstallCmd   []string
	ReinstallCmd []string
	UninstallCmd []string

	OutC, ErrC chan string
}

func Source(name string, installCmd []string, reinstallCmd []string, uninstallCmd []string, outC chan string, errC chan string) *source {
	return &source{&sync.Mutex{}, name, installCmd, reinstallCmd, uninstallCmd, outC, errC}
}

func (s *source) cc(C chan string, mC <-chan executor.Message) {
	defer func() { _ = recover() }()
	for {
		if msg, okk := <-mC; !okk {
			break
		} else if c := C; c != nil {
			c <- fmt.Sprintf("[%d: %s] %s", msg.Step.Id(), msg.Step.CmdName(), msg.Text)
		}
	}
}

func (s *source) outCC(mC <-chan executor.Message) { s.cc(s.OutC, mC) }
func (s *source) errCC(mC <-chan executor.Message) { s.cc(s.ErrC, mC) }

func (s *source) start(cmd []string) (ok bool, err error) {
	if !s.mu.TryLock() {
		return false, ErrInstallRunning
	}
	defer s.mu.Unlock()
	etr := executor.NewExecutorSteps(false, cmd...)
	var (
		outC = make(chan executor.Message, 1)
		errC = make(chan executor.Message, 1)
	)
	etr.Chan(outC, errC)
	return etr.Exec()
}

func (s *source) Install() (ok bool, err error)   { return s.start(s.InstallCmd) }
func (s *source) Reinstall() (ok bool, err error) { return s.start(s.ReinstallCmd) }
func (s *source) Uninstall() (ok bool, err error) { return s.start(s.UninstallCmd) }
