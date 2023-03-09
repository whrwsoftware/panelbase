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
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/appver"
)

type App struct {
	*apps.GroupedApp

	PostfixApp, DovecotApp, RoundcubeApp             app.Applicable
	PostfixVersion, DovecotVersion, RoundcubeVersion string
}

func NewApp(postfixApp, dovecotApp, roundcubeApp app.Applicable, postfixVersion, dovecotVersion, roundcubeVersion string) app.Applicable {
	return &App{
		GroupedApp: apps.NewGroupApp(
			apps.NewBundle(postfixApp, postfixVersion),
			apps.NewBundle(dovecotApp, dovecotVersion),
			apps.NewBundle(roundcubeApp, roundcubeVersion),
		),
		PostfixApp:       postfixApp,
		DovecotApp:       dovecotApp,
		RoundcubeApp:     roundcubeApp,
		PostfixVersion:   postfixVersion,
		DovecotVersion:   dovecotVersion,
		RoundcubeVersion: roundcubeVersion,
	}
}

func (a *App) Check() (ok bool, err error) {
	// 1. nginx@1.8.1
	// 2. >=php@7.3
	// 3. >=php-fpm@7.3
	// 4. =roundcube@1.6.1
	// 5. =postfix@3.7.4
	// 6. =dovecot@2.3.17
	err = app.Manager.Required(
		app.NewRequired(appver.Nginx.Name, appver.NginxMinVersion().Version, appver.NginxMinVersion().VersionId), // nginx for min version
		app.NewRequired(appver.Php.Name, appver.Php73Version().Version, appver.Php73Version().VersionId),         // >=php7.3 >=php-fpm@7.3
	)
	ok = true
	return
}
