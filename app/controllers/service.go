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

package controllers

import (
	"fmt"
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/cmds"
	"github.com/whrwsoftware/panelbase/executor"
)

type service struct {
	Name string
	*cmds.Service
	app.Logger
}

func Service(name string, logger app.Logger) *service {
	return &service{Name: name, Service: cmds.NewService(name), Logger: logger}
}

func (s *service) run(cmd string) (ok bool, err error) {
	return executor.NewBashExecutor(fmt.Sprintf("service %s %s", s.Name, cmd), s.File()).Exec().Release()
}

func (s *service) Enable() (ok bool, err error)  { return s.run("enable") }
func (s *service) Disable() (ok bool, err error) { return s.run("disable") }
func (s *service) Start() (ok bool, err error)   { return s.run("start") }
func (s *service) Stop() (ok bool, err error)    { return s.run("stop") }
func (s *service) Restart() (ok bool, err error) { return s.run("restart") }
func (s *service) Status() (st string, ok bool, err error) {
	return executor.NewBashExecutor(fmt.Sprintf("service %s %s", s.Name, "status"), s.File()).Run().OutRelease()
}
