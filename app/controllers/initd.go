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
	"github.com/whrwsoftware/panelbase/executor"
)

type initD struct {
	BashFile string

	app.Logger
}

func InitD(bashFile string, logger app.Logger) *initD {
	return &initD{BashFile: bashFile, Logger: logger}
}

func (i *initD) run(cmd string) (ok bool, err error) {
	return executor.NewBashExecutor(fmt.Sprintf("/bin/bash /etc/init.d/%s %s", i.BashFile, cmd), i.File()).Exec().Release()
}

func (i *initD) Enable() (ok bool, err error)  { return i.run("enable") }
func (i *initD) Disable() (ok bool, err error) { return i.run("disable") }
func (i *initD) Start() (ok bool, err error)   { return i.run("start") }
func (i *initD) Stop() (ok bool, err error)    { return i.run("stop") }
func (i *initD) Restart() (ok bool, err error) { return i.run("restart") }
func (i *initD) Status() (st string, ok bool, err error) {
	return executor.NewBashExecutor(fmt.Sprintf("/bin/bash /etc/init.d/%s %s", i.BashFile, "status"), i.File()).Run().OutRelease()
}
