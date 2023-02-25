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
	"errors"
	"sync"
	"time"
)

var (
	ErrStepsEmpty = errors.New("executor: the steps empty")
)

type (
	Executor struct {
		steps  []*Step
		killed bool
		async  bool

		outC, errC chan<- Message
	}
)

func NewExecutor(async bool) *Executor { return &Executor{steps: make([]*Step, 0), async: async} }

func (etr *Executor) Chan(outC, errC chan<- Message) {
	etr.outC = outC
	etr.errC = errC
}

func (etr *Executor) NextId() int { return len(etr.steps) + 1 }

func (etr *Executor) Create(cmd string, args ...string) { etr.Add(NewStep(etr.NextId(), cmd, args...)) }

func (etr *Executor) Add(step *Step) {
	if step != nil {
		etr.steps = append(etr.steps, step)
	}
}

func (etr *Executor) Clear() { etr.steps = make([]*Step, 0) }

func (etr *Executor) execAsync() (err error) {
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
			v, _ := <-errCh
			nextCh <- struct{}{}
			outErrCh <- v
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
				if sErr := step.Start(etr.outC, etr.errC); sErr != nil {
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

func (etr *Executor) execSync() (err error) {
	for _, step := range etr.steps {
		if err = step.Start(etr.outC, etr.errC); err != nil {
			return
		}
		if etr.killed {
			return
		}
	}
	return
}

func (etr *Executor) Exec() (err error) {
	if len(etr.steps) == 0 {
		return ErrStepsEmpty
	}

	if etr.async {
		return etr.execAsync()
	}

	err = etr.execSync()
	if etr.outC != nil {
		close(etr.outC)
	}
	if etr.errC != nil {
		close(etr.errC)
	}
	return
}

func (etr *Executor) Kill() {
	if !etr.killed {
		for _, step := range etr.steps {
			_ = step.Kill()
		}
		etr.killed = true
		if etr.outC != nil {
			close(etr.outC)
		}
		if etr.errC != nil {
			close(etr.errC)
		}
	}
}
