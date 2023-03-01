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

type commandLine struct {
	StartCmd   string
	StopCmd    string
	RestartCmd string
	VersionCmd string
}

func CommandLine(startCmd, stopCmd, restartCmd, versionCmd string) apptpl.Controller {
	return &commandLine{StartCmd: startCmd, StopCmd: stopCmd, RestartCmd: restartCmd, VersionCmd: versionCmd}
}

func (c *commandLine) run(v string) (ok bool, err error) { _, ok, err = c.runWithOut(v); return }

func (c *commandLine) runWithOut(v string) (s string, ok bool, err error) {
	s, _, ok, err = cmd.RunFullCmd(v)
	return
}

func (c *commandLine) Start() (ok bool, err error)             { return c.run(c.StartCmd) }
func (c *commandLine) Stop() (ok bool, err error)              { return c.run(c.StopCmd) }
func (c *commandLine) Restart() (ok bool, err error)           { return c.run(c.RestartCmd) }
func (c *commandLine) Version() (v string, ok bool, err error) { return c.runWithOut(c.VersionCmd) }
