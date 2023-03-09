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
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/executor"
)

type commandLine struct {
	EnableCmd  string
	DisableCmd string
	StartCmd   string
	StopCmd    string
	RestartCmd string
	StatusCmd  string

	app.Logger
}

func CommandLine(enableCmd, disableCmd, startCmd, stopCmd, restartCmd, statusCmd string, logger app.Logger) *commandLine {
	return &commandLine{
		EnableCmd:  enableCmd,
		DisableCmd: disableCmd,
		StartCmd:   startCmd,
		StopCmd:    stopCmd,
		RestartCmd: restartCmd,
		StatusCmd:  statusCmd,
		Logger:     logger,
	}
}

func (c *commandLine) run(cmd string) (ok bool, err error) {
	return executor.NewBashExecutor(cmd).BindLog(c.File()).Exec().Release()
}

func (c *commandLine) Enable() (ok bool, err error)  { return c.run(c.EnableCmd) }
func (c *commandLine) Disable() (ok bool, err error) { return c.run(c.DisableCmd) }
func (c *commandLine) Start() (ok bool, err error)   { return c.run("start") }
func (c *commandLine) Stop() (ok bool, err error)    { return c.run("stop") }
func (c *commandLine) Restart() (ok bool, err error) { return c.run("restart") }
func (c *commandLine) Status() (st string, ok bool, err error) {
	return executor.NewBashExecutor(c.StatusCmd).Run().OutRelease()
}
