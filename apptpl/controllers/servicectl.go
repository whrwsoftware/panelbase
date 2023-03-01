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
	"github.com/whrwsoftware/panelbase/apptpl"
	"github.com/whrwsoftware/panelbase/cmd"
)

type serviceCtl struct {
	Name   string
	VerCmd string
}

func ServiceCtl(name string, verCmd string) apptpl.Controller {
	return &serviceCtl{Name: name, VerCmd: verCmd}
}

func (s *serviceCtl) run(v string) (ok bool, err error) { _, ok, err = s.runWithOut(v); return }

func (s *serviceCtl) runWithOut(v string) (str string, ok bool, err error) {
	str, _, ok, err = cmd.Run("service", s.Name, v)
	return
}

func (s *serviceCtl) Start() (ok bool, err error)   { return s.run("start") }
func (s *serviceCtl) Stop() (ok bool, err error)    { return s.run("stop") }
func (s *serviceCtl) Restart() (ok bool, err error) { return s.run("restart") }
func (s *serviceCtl) Version() (v string, ok bool, err error) {
	v, _, ok, err = cmd.RunFullCmd(s.VerCmd)
	return
}
