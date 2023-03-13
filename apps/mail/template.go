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

package mail

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/appmanager"
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type template struct {
	PostfixAppTemplate, DovecotAppTemplate, RoundcubeAppTemplate apps.AppTemplate
}

func Template(postfixAppTemplate, dovecotAppTemplate, roundcubeAppTemplate apps.AppTemplate) *template {
	return &template{postfixAppTemplate, dovecotAppTemplate, roundcubeAppTemplate}
}

func (t *template) getApp(manager appmanager.Manager, os oss.OS) app.Applicable {
	return NewApp(
		apps.GetApp(t.PostfixAppTemplate, manager, os),
		apps.GetApp(t.DovecotAppTemplate, manager, os),
		apps.GetApp(t.RoundcubeAppTemplate, manager, os),
		appver.PostfixVersion().Version,
		appver.DovecotVersion().Version,
		appver.RoundcubeMaxVersion().Version,
	)
}

func (t *template) CentOS7(manager appmanager.Manager) app.Applicable {
	return t.getApp(manager, oss.CentOS7)
}

func (t *template) CentOS8(manager appmanager.Manager) app.Applicable {
	return t.getApp(manager, oss.CentOS8)
}

func (t *template) Ubuntu(manager appmanager.Manager) app.Applicable {
	return t.getApp(manager, oss.Ubuntu)
}

func (t *template) Debian(manager appmanager.Manager) app.Applicable {
	return t.getApp(manager, oss.Debian)
}

func (t *template) Arch(manager appmanager.Manager) app.Applicable {
	return t.getApp(manager, oss.Arch)
}
