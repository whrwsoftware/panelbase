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

type aptGet struct {
	Name       string
	OutC, ErrC chan string
	*cmds.AptGet
}

func AptGet(name string, outC chan string, errC chan string) *aptGet {
	return &aptGet{name, outC, errC, cmds.NewAptGet(name)}
}

func (a *aptGet) Install() (ok bool, err error)   { return a.AptGet.Install(a.OutC, a.ErrC) }
func (a *aptGet) Reinstall() (ok bool, err error) { return a.AptGet.Reinstall(a.OutC, a.ErrC) }
func (a *aptGet) Uninstall() (ok bool, err error) { return a.AptGet.Uninstall(a.OutC, a.ErrC) }
