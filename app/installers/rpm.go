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
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/cmds"
)

type rpm struct {
	Name       string
	OutC, ErrC chan string
	*cmds.Rpm
}

func Rpm(name string, outC chan string, errC chan string) app.Installer {
	return &rpm{name, outC, errC, cmds.NewRpm(name)}
}

func (r *rpm) Install() (ok bool, err error)   { return r.Rpm.Install(r.OutC, r.ErrC) }
func (r *rpm) Reinstall() (ok bool, err error) { return r.Rpm.Reinstall(r.OutC, r.ErrC) }
func (r *rpm) Uninstall() (ok bool, err error) { return r.Rpm.Uninstall(r.OutC, r.ErrC) }
