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

package nginx

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/installers"
	"github.com/whrwsoftware/panelbase/appmanager"
)

type template struct{}

func Template() *template { return &template{} }

func (t *template) CentOS7(manager appmanager.Manager) app.Applicable {
	return app.NewApplication(checker, installers.Versioned(installer), controller, configurator, logger, manager)
}

func (t *template) CentOS8(manager appmanager.Manager) app.Applicable { return t.CentOS7(manager) }

func (t *template) Ubuntu(manager appmanager.Manager) app.Applicable {
	//TODO implement me
	panic("implement me")
}

func (t *template) Debian(manager appmanager.Manager) app.Applicable {
	//TODO implement me
	panic("implement me")
}

func (t *template) Arch(manager appmanager.Manager) app.Applicable {
	//TODO implement me
	panic("implement me")
}
