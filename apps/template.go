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

package apps

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/appmanager"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type AppTemplate interface {
	CentOS7(manager appmanager.Manager) app.Applicable
	CentOS8(manager appmanager.Manager) app.Applicable
	Ubuntu(manager appmanager.Manager) app.Applicable
	Debian(manager appmanager.Manager) app.Applicable
	Arch(manager appmanager.Manager) app.Applicable
}

func GetApp(appTemplate AppTemplate, manager appmanager.Manager, os oss.OS) (app app.Applicable) {
	switch os {
	case oss.CentOS7:
		return appTemplate.CentOS7(manager)
	case oss.CentOS8:
		return appTemplate.CentOS8(manager)
	case oss.Ubuntu:
		return appTemplate.Ubuntu(manager)
	case oss.Debian:
		return appTemplate.Debian(manager)
	case oss.Arch:
		return appTemplate.Arch(manager)
	}
	return
}
