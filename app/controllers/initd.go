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
	"github.com/whrwsoftware/panelbase/appconf/initd"
	"github.com/whrwsoftware/panelbase/cmd"
	"path/filepath"
)

type initD struct {
	File       string
	IsFullPath bool
}

func InitD(file string, isFullPath bool) *initD { return &initD{File: file, IsFullPath: isFullPath} }

func (i *initD) run(v string) (ok bool, err error) { _, ok, err = i.runWithOut(v); return }

func (i *initD) runWithOut(v string) (s string, ok bool, err error) {
	if i.IsFullPath {
		s, _, ok, err = cmd.Run(i.File, v)
		return
	}
	s, _, ok, err = cmd.Run(filepath.Join(initd.RootInitD, i.File), v)
	return
}

func (i *initD) Start() (ok bool, err error)             { return i.run("start") }
func (i *initD) Stop() (ok bool, err error)              { return i.run("stop") }
func (i *initD) Restart() (ok bool, err error)           { return i.run("restart") }
func (i *initD) Version() (v string, ok bool, err error) { return i.runWithOut("version") }
