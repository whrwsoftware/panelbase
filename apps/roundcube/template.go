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

package roundcube

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/appver"
)

type template struct{}

func Template() *template { return &template{} }

func (t *template) defaultGetApp() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return app.NewApplication(checker, installers.Versioned(map[string]app.Installer{appver.RoundcubeVersion().Version: installers.Source(installCmd(), uninstallCmd(), reinstallCmd(), outC, errC)}), controller, configurator)
	}
}

func (t *template) CentOS7() apps.GetAppFunc { return t.defaultGetApp() }
func (t *template) CentOS8() apps.GetAppFunc { return t.defaultGetApp() }
func (t *template) Ubuntu() apps.GetAppFunc  { return t.defaultGetApp() }
func (t *template) Debian() apps.GetAppFunc  { return t.defaultGetApp() }
func (t *template) Arch() apps.GetAppFunc    { return t.defaultGetApp() }
