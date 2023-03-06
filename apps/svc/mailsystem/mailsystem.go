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

package mailsystem

import "github.com/whrwsoftware/panelbase/app"

type App struct {
	PostfixApp app.Application
	DovecotApp app.Application
}

func NewApp(postfixApp app.Application, dovecotApp app.Application) app.Applicable {
	return &App{PostfixApp: postfixApp, DovecotApp: dovecotApp}
}

func (a *App) Check() (ok bool, err error) {
	return
}

func (a *App) Install() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Uninstall() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Reinstall() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Start() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Stop() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Restart() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Version() (v string, ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Configure(m map[string]any) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Load(name string) (v string, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Clean() (err error) {
	//TODO implement me
	panic("implement me")
}
