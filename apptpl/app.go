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

package apptpl

type Checker interface{ Check() (ok bool, err error) }

type Installer interface {
	Install() (ok bool, err error)
	Uninstall() (ok bool, err error)
}

type Controller interface {
	Start() (ok bool, err error)
	Stop() (ok bool, err error)
	Restart() (ok bool, err error)
	Version() (v string, ok bool, err error)
}

type Configurator interface{ Configure(cfg Cfg) (err error) }

type Cfg struct {
	Name string // The configuration name
	Data string // The configuration data
}

type Applicable interface {
	Checker
	Installer
	Controller
	Configurator
}

type Application struct {
	Checker
	Installer
	Controller
	Configurator
}

func NewApplication(checker Checker, installer Installer, controller Controller, configurator Configurator) *Application {
	return &Application{Checker: checker, Installer: installer, Controller: controller, Configurator: configurator}
}
