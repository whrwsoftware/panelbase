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

package app

type PreChecker interface{ Check() (ok bool, err error) }
type Installer interface{ Install() (ok bool, err error) }
type Uninstaller interface{ Uninstall() (ok bool, err error) }
type Configurator interface {
	Configure(c Cfg) (ok bool, err error)
}

type ServiceInstaller interface{ ServiceInstall() (ok bool, err error) }
type ServiceEnabler interface{ ServiceEnable() (ok bool, err error) }
type ServiceDisabler interface{ ServiceDisable() (ok bool, err error) }

type Starter interface{ Start() (ok bool, err error) }
type Stopper interface{ Stop() (ok bool, err error) }
type Restarter interface{ Restart() (ok bool, err error) }
type Status interface {
	Status() (s string, ok bool, err error)
}
type Version interface {
	Version() (v string, ok bool, err error)
}

type Installation interface {
	PreChecker
	Installer
	Uninstaller
	Configurator
}

type ServiceManager interface {
	ServiceInstaller
	ServiceEnabler
	ServiceDisabler
}

type Controller interface {
	Starter
	Stopper
	Restarter
	Status
	Version
}

type Application interface {
	Installation
	ServiceManager
	Controller
}
