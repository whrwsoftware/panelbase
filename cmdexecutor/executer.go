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
	"errors"
	"github.com/whrwsoftware/panelbase/cmd"
	"os/exec"
	"sync"
	"time"
)

var (
	ErrStepsEmpty = errors.New("executor: the steps empty")
)

type (
	Executor struct {
		steps      []*Step
		killed     bool
		async      bool
		outC, errC chan<- string
	}
)

func NewExecutor(async bool) *Executor { return &Executor{steps: make([]*Step, 0), async: async} }

func (etr *Executor) NextId() int { return len(etr.steps) + 1 }

func (etr *Executor) Create(cmd string, args []string, passCheckFunc func() bool) {
	etr.Add(NewStep(etr.NextId(), cmd, args, passCheckFunc))
}

func (etr *Executor) kill(step *Step) (err error) { return step.kill() }

func (etr *Executor) start(step *Step) (err error) {
	var (
		outC = make(chan string, 1)
		errC = make(chan string, 1)
	)
	go func() {
		for {
			if outStr, ok := <-outC; !ok {
				break
			} else {
				etr.outC <- step.logPrefix() + " " + outStr
			}
		}
	}()
	go func() {
		for {
			if errStr, ok := <-errC; !ok {
				break
			} else {
				etr.errC <- step.logPrefix() + " " + errStr
			}
		}
	}()
	err = cmd.Start(step.name, step.args, outC, errC, nil, func(outCmd *exec.Cmd) {
		step.cmd = outCmd
		step.started = true
	})
	step.finished = true
	step.success = err == nil
	step.passed = step.success
	if fn := step.passCheckFunc; fn != nil {
		step.passed = fn()
	}
	return
}

func (etr *Executor) Add(step *Step) {
	if step != nil {
		etr.steps = append(etr.steps, step)
	}
}

func (etr *Executor) Clear() { etr.steps = make([]*Step, 0) }

func (etr *Executor) startAsync() (err error) {
	var (
		nextCh   = make(chan struct{}, 1)
		outErrCh = make(chan error, 1)
	)
	go func() {
		var (
			wg    = &sync.WaitGroup{}
			errCh = make(chan error, 1)
		)
		wg.Add(len(etr.steps))
		go func() {
			if v, ok := <-errCh; ok {
				nextCh <- struct{}{}
				outErrCh <- v
			}
		}()
		go func() {
			for {
				if etr.killed {
					nextCh <- struct{}{}
					outErrCh <- nil
				}
				time.Sleep(time.Millisecond)
			}
		}()
		for _, step := range etr.steps {
			go func(step *Step, errCh chan<- error) {
				defer wg.Done()
				if sErr := etr.start(step); sErr != nil {
					errCh <- sErr
					close(errCh)
				}
			}(step, errCh)
		}
		wg.Wait()
		nextCh <- struct{}{}
		outErrCh <- nil
	}()
	<-nextCh
	err = <-outErrCh
	close(nextCh)
	close(outErrCh)
	return
}

func (etr *Executor) startSync() (err error) {
	for _, step := range etr.steps {
		if err = etr.start(step); err != nil {
			return
		}
		if etr.killed {
			return
		}
	}
	return
}

func (etr *Executor) Start() (err error) {
	if len(etr.steps) == 0 {
		return ErrStepsEmpty
	}

	if etr.async {
		return etr.startAsync()
	}

	return etr.startSync()
}

func (etr *Executor) Kill() {
	etr.killed = true
	for _, step := range etr.steps {
		_ = etr.kill(step)
	}
}
