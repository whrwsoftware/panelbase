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
	"github.com/whrwsoftware/panelbase/cmds"
)

type apt struct {
	Name       string
	OutC, ErrC chan string
	*cmds.Apt
}

func Apt(name string, outC chan string, errC chan string) *apt {
	return &apt{name, outC, errC, cmds.NewApt(name)}
}

func (a *apt) Install() (ok bool, err error)   { return a.Apt.Install(a.OutC, a.ErrC) }
func (a *apt) Reinstall() (ok bool, err error) { return a.Apt.Reinstall(a.OutC, a.ErrC) }
func (a *apt) Uninstall() (ok bool, err error) { return a.Apt.Uninstall(a.OutC, a.ErrC) }
