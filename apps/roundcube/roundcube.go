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

package roundcube

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/app/controllers"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/app/loggers"
	"github.com/whrwsoftware/panelbase/appconf/roundcube"
	"github.com/whrwsoftware/panelbase/apps/roundcube/latest"
	"github.com/whrwsoftware/panelbase/appver"
)

var (
	name         = appver.Roundcube.Name
	curVer       = appver.RoundcubeVersion()
	checker      = &defaultChecker{}
	logger       = loggers.File(curVer.LogFile)
	configurator = configurators.DefaultConfigurator(roundcube.ConfBinds)
	controller   = controllers.NoController()
	installer    = map[string]app.VersionInstaller{curVer.Version: installers.BashInstallerDebug(latest.FS, logger)}
)

type defaultChecker struct{}

func (*defaultChecker) Check(manager app.Manager) (ok bool, err error) {
	// 1. nginx@1.8.1
	// 2. >=php@7.3
	// 3. >=php-fpm@7.3
	// 4. =roundcube@1.6.1
	// 5. =postfix@3.7.4
	// 6. =dovecot@2.3.17
	nginxVersion := appver.NginxMinVersion()
	php73Version := appver.Php73Version()
	err = manager.Required(
		app.NewRequired(appver.Nginx.Name, nginxVersion.Version, nginxVersion.VersionId), // nginx for min version
		app.NewRequired(appver.Php.Name, php73Version.Version, php73Version.VersionId),   // >=php7.3 >=php-fpm@7.3
	)
	ok = true
	return
}
