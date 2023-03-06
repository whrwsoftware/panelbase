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

package dovecot

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/checkers"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/appconf/dovecot"
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/apps/svc"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

const (
	name = "dovecot"
)

var (
	checker      = checkers.NoChecker()
	configurator = configurators.DefaultConfigurator(dovecot.ConfBinds)
)

func GetApp(outC, errC chan string) (app app.Applicable) {
	return apps.GetApp(&svc.StdTemplate{Name: name, Ver: "", Pkg: name, Checker: checker, Configurator: configurator}, oss.CurrentOS(), outC, errC)
}
