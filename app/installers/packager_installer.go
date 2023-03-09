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
)

type packagerInstaller struct {
	app.VersionInstaller
	Debug bool
}

func PackagerInstaller(pkg string, packager Packager, logger app.Logger) *packagerInstaller {
	return &packagerInstaller{VersionInstaller: BashInstaller(PackagerBashFS(pkg, packager), logger)}
}

func PackagerInstallerWithDebug(pkg string, packager Packager, logger app.Logger, debug bool) *packagerInstaller {
	return &packagerInstaller{VersionInstaller: BashInstallerWithDebug(PackagerBashFS(pkg, packager), logger, debug)}
}

func PackagerInstallerDebug(pkg string, packager Packager, logger app.Logger) *packagerInstaller {
	return PackagerInstallerWithDebug(pkg, packager, logger, true)
}

type Packager interface {
	InstallCmd(pkg string) string
	UninstallCmd(pkg string) string
	Reinstall(pkg string) string
}

type packagerBashFS struct {
	Pkg string
	Packager
}

func PackagerBashFS(pkg string, packager Packager) *packagerBashFS {
	return &packagerBashFS{Pkg: pkg, Packager: packager}
}

func (p *packagerBashFS) Install() string   { return p.Packager.InstallCmd(p.Pkg) }
func (p *packagerBashFS) Uninstall() string { return p.Packager.UninstallCmd(p.Pkg) }
func (p *packagerBashFS) Reinstall() string { return p.Packager.Reinstall(p.Pkg) }
