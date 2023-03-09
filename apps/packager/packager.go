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

package packager

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/checkers"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/app/controllers"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/app/installers/packagers"
	"github.com/whrwsoftware/panelbase/app/loggers"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

var (
	checker      = checkers.NoChecker()
	logger       = loggers.NoLogger()
	configurator = configurators.NoConfigurator()
	controller   = controllers.NoController()
	installer    = func(pkg string, os oss.OS, logger app.Logger) app.Installer {
		switch os {
		default:
			return installers.NoInstaller()
		case oss.Unknown:
			return installers.NoInstaller()
		case oss.CentOS7, oss.CentOS8, oss.CentOS9:
			return installers.Versioned(map[string]app.VersionInstaller{"@": installers.PackagerInstaller(pkg, packagers.YumPackager(), logger)})
		case oss.Ubuntu, oss.Debian:
			return installers.Versioned(map[string]app.VersionInstaller{"@": installers.PackagerInstaller(pkg, packagers.AptPackager(), logger)})
		case oss.Arch:
			return installers.Versioned(map[string]app.VersionInstaller{"@": installers.PackagerInstaller(pkg, packagers.PacmanPackager(), logger)})
		}
	}
)
