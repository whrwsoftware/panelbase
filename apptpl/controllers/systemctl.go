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
	"github.com/whrwsoftware/panelbase/cmds"
)

type systemctl struct {
	Name   string
	VerCmd string
	*cmds.Systemctl
}

func Systemctl(name string, verCmd string) apptpl.Controller {
	return &systemctl{name, verCmd, cmds.NewSystemctl(name)}
}

func (s *systemctl) Start() (ok bool, err error)   { _, ok, err = s.Systemctl.Start(); return }
func (s *systemctl) Stop() (ok bool, err error)    { _, ok, err = s.Systemctl.Stop(); return }
func (s *systemctl) Restart() (ok bool, err error) { _, ok, err = s.Systemctl.Restart(); return }
func (s *systemctl) Version() (v string, ok bool, err error) {
	v, _, ok, err = cmd.RunFullCmd(s.VerCmd)
	return
}
