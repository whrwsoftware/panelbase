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

package appmanager

import "strings"

type Manager interface {
	Add(app *App) (err error)
	Update(app *App) (err error)
	Delete(pkg string) (err error)
	AddTag(pkg string, tag string) (err error)
	DeleteTag(pkg string, tag string) (err error)
	Installed(pkg string) (err error)
	Uninstalled(pkg string) (err error)
	Check(requiredPkg ...RequiredApp) (err error)
	FindAll() (apps []*App, err error)
	FindByPkg(pkg string) (app *App, err error)
	FindsByType(typ Type) (app []*App, err error)
	FindsByName(name string) (apps []*App, err error)
	FindsByTag(tag string) (apps []*App, err error)
}

type RequiredApp struct {
	Name             string
	OneOfThemVersion []string
}

var mgr Manager

func Use(manager Manager)                              { mgr = manager }
func GetManager() Manager                              { return mgr }
func Add(app *App) (err error)                         { return GetManager().Add(app) }
func Update(app *App) (err error)                      { return GetManager().Update(app) }
func Delete(pkg string) (err error)                    { return GetManager().Delete(pkg) }
func AddTag(pkg string, tag string) (err error)        { return GetManager().AddTag(pkg, tag) }
func DeleteTag(pkg string, tag string) (err error)     { return GetManager().DeleteTag(pkg, tag) }
func Installed(pkg string) (err error)                 { return GetManager().Installed(pkg) }
func Uninstalled(pkg string) (err error)               { return GetManager().Uninstalled(pkg) }
func Check(requiredPkg ...RequiredApp) (err error)     { return GetManager().Check(requiredPkg...) }
func FindAll() (apps []*App, err error)                { return GetManager().FindAll() }
func FindByPkg(pkg string) (app *App, err error)       { return GetManager().FindByPkg(pkg) }
func FindsByType(typ Type) (app []*App, err error)     { return GetManager().FindsByType(typ) }
func FindsByName(name string) (apps []*App, err error) { return GetManager().FindsByName(name) }
func FindsByTag(tag string) (apps []*App, err error)   { return GetManager().FindsByTag(tag) }

type CheckError struct {
	NotExistApp   []string
	NotInstallApp []string
}

func (c *CheckError) Error() string {
	errText := "check error: "
	if ss := c.NotExistApp; ss != nil && len(ss) > 0 {
		errText += "NotExistApp(" + strings.Join(ss, ",") + ") "
	}
	if ss := c.NotInstallApp; ss != nil && len(ss) > 0 {
		errText += "NotInstallApp(" + strings.Join(ss, ",") + ")"
	}
	return errText
}
