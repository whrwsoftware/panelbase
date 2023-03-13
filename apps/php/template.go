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

package php

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/appmanager"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type template struct {
	phpV int
	*appver.PhpVer
}

func Template(phpV int) *template { return &template{phpV, appver.FindPhpVerByV(phpV)} }

func (t *template) getApp(manager appmanager.Manager, os oss.OS) app.Applicable {
	return app.NewApplication(checker, getInstaller(t.PhpVer, os), getController(t.PhpVer), getConfigurator(t.PhpVer), getLogger(t.PhpVer), manager)
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
