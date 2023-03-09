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

package installers

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/executor"
)

type bashInstaller struct {
	app.BashFS
	app.Logger
	Debug bool
}

func BashInstaller(bashFS app.BashFS, logger app.Logger) *bashInstaller {
	return &bashInstaller{BashFS: bashFS, Logger: logger}
}

func BashInstallerWithDebug(bashFS app.BashFS, logger app.Logger, debug bool) *bashInstaller {
	return &bashInstaller{BashFS: bashFS, Logger: logger, Debug: debug}
}

func BashInstallerDebug(bashFS app.BashFS, logger app.Logger) *bashInstaller {
	return BashInstallerWithDebug(bashFS, logger, true)
}

func (b *bashInstaller) Install() (ok bool, err error) {
	return executor.NewBashExecutor(b.BashFS.Install()).BindLog(b.File()).SetDebug(b.Debug).Exec().Release()
}

func (b *bashInstaller) Uninstall() (ok bool, err error) {
	return executor.NewBashExecutor(b.BashFS.Uninstall()).BindLog(b.File()).SetDebug(b.Debug).Exec().Release()
}

func (b *bashInstaller) Reinstall() (ok bool, err error) {
	return executor.NewBashExecutor(b.BashFS.Reinstall()).BindLog(b.File()).SetDebug(b.Debug).Exec().Release()
}
