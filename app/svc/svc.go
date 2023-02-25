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

package svc

import (
	"github.com/whrwsoftware/panelbase/app"
)

type Service interface {
	app.Precondition
	app.ServiceEnabler
	app.ServiceDisabler
	app.Installer
	app.Uninstaller
	app.Configurator
	app.Starter
	app.Stopper
	app.Restarter
	app.Status
	app.Version
}

type Template struct {
}

func (t *Template) Check() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Enable() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Disable() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Install() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Uninstall() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Configure(data any) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Start() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Stop() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Restart() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Ver() (v string, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template) Status() (s string, err error) {
	//TODO implement me
	panic("implement me")
}
