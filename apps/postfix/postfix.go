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

package postfix

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/checkers"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/app/controllers"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/app/loggers"
	"github.com/whrwsoftware/panelbase/appconf/postfix"
	"github.com/whrwsoftware/panelbase/apps/postfix/latest"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

var (
	name         = appver.Postfix.Name
	curVer       = appver.PostfixVersion()
	checker      = checkers.NoChecker()
	logger       = loggers.File(curVer.LogFile)
	configurator = configurators.DefaultConfigurator(postfix.ConfBinds)
	controller   = controllers.Systemctl(name, logger)
	getInstaller = func(os oss.OS) app.Installer {
		return installers.Versioned(map[string]app.VersionInstaller{curVer.Version: installers.BashInstaller(latest.FS(os), logger)})
	}
)
