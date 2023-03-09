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

type Checker interface{ Check() (ok bool, err error) }

type Installer interface {
	Install(ver string) (ok bool, err error)
	Uninstall(ver string) (ok bool, err error)
	Reinstall(ver string) (ok bool, err error)
}

type VersionInstaller interface {
	Install() (ok bool, err error)
	Uninstall() (ok bool, err error)
	Reinstall() (ok bool, err error)
}

type Controller interface {
	Enable() (ok bool, err error)
	Disable() (ok bool, err error)
	Start() (ok bool, err error)
	Stop() (ok bool, err error)
	Restart() (ok bool, err error)
	Status() (st string, ok bool, err error)
}

type Configurator interface {
	Configure(m map[string]any) (err error)
	ConfigurationLoader
	ConfigurationCleaner
}

type ConfigurationLoader interface {
	Load(name string) (v string, err error)
}

type ConfigurationCleaner interface {
	Clean() (err error)
}

type Logger interface {
	File() string
	Truncate() (err error)
}

type Applicable interface {
	Checker
	Installer
	Controller
	Configurator
	Logger
}

type Application struct {
	Checker
	Installer
	Controller
	Configurator
	Logger
}

type BashFS interface {
	Install() string
	Uninstall() string
	Reinstall() string
}

func NewApplication(checker Checker, installer Installer, controller Controller, configurator Configurator, logger Logger) Applicable {
	return &Application{Checker: checker, Installer: installer, Controller: controller, Configurator: configurator, Logger: logger}
}
