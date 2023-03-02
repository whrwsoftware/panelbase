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

package tpl

import (
	"github.com/whrwsoftware/panelbase/apptpl"
	"github.com/whrwsoftware/panelbase/apptpl/controllers"
	"github.com/whrwsoftware/panelbase/apptpl/installers"
)

func YumApp(name string, ver string, pkg string, outC chan string, errC chan string, checker apptpl.Checker, configurator apptpl.Configurator) apptpl.Applicable {
	return apptpl.NewApplication(checker, installers.Yum(pkg, outC, errC), controllers.Systemctl(name, "echo "+ver), configurator)
}
