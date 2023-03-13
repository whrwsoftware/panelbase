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

package php

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/checkers"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/app/controllers"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/app/loggers"
	"github.com/whrwsoftware/panelbase/appconf/php"
	"github.com/whrwsoftware/panelbase/apps/php/latest"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

var (
	checker         = checkers.NoChecker()
	getLogger       = func(v *appver.PhpVer) app.Logger { return loggers.File(v.LogFile) }
	getConfigurator = func(v *appver.PhpVer) app.Configurator {
		return configurators.DefaultConfigurator(php.ConfBinds(v.VersionV))
	}
	getController = func(v *appver.PhpVer) app.Controller { return controllers.Systemctl(v.Service, getLogger(v)) }
	getInstaller  = func(v *appver.PhpVer, os oss.OS) app.Installer {
		return installers.Versioned(map[string]app.VersionInstaller{v.Version: installers.BashInstaller(latest.FS(os, v.Pkg), getLogger(v))})
	}
)
