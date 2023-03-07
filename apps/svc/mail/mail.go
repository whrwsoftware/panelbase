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

package mail

import (
	"github.com/whrwsoftware/panelbase/app"
)

//Apache, Lighttpd, Nginx, IIS, Cherokee or Hiawatha web server
//PHP Version 7.3 or greater1
//MySQL/MariaDB, PostgreSQL, SQLite, MSSQL or Oracle database
//SMTP server and IMAP server with IMAP4 rev1 support

type App struct {
	PostfixApp app.Application
	DovecotApp app.Application

	//Apache, Lighttpd, Nginx, IIS, Cherokee or Hiawatha web server
	//PHP Version 7.3 or greater1
	//MySQL/MariaDB, PostgreSQL, SQLite, MSSQL or Oracle database
	//SMTP server and IMAP server with IMAP4 rev1 support

	RoundcubeApp app.Application
}

func NewApp(postfixApp, dovecotApp, roundcubeApp app.Application) app.Applicable {
	return &App{PostfixApp: postfixApp, DovecotApp: dovecotApp, RoundcubeApp: roundcubeApp}
}

func (a *App) Check() (ok bool, err error) {
	// 1. nginx@any
	// 2. >=php@7.3
	// 3. >=php-fpm@7.3
	// 4. =roundcube@1.6.1
	// 5. =postfix@3.4.7
	// 6. =dovecot@2.2.36
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
