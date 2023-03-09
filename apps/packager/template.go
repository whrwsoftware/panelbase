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

package packager

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type template struct {
	Pkg string
}

func Template(pkg string) *template { return &template{Pkg: pkg} }

func (t *template) CentOS7() app.Applicable {
	return app.NewApplication(checker, installer(t.Pkg, oss.CentOS7, logger), controller, configurator, logger)
}

func (t *template) CentOS8() app.Applicable {
	return app.NewApplication(checker, installer(t.Pkg, oss.CentOS8, logger), controller, configurator, logger)
}

func (t *template) Ubuntu() app.Applicable {
	return app.NewApplication(checker, installer(t.Pkg, oss.Ubuntu, logger), controller, configurator, logger)
}

func (t *template) Debian() app.Applicable {
	return app.NewApplication(checker, installer(t.Pkg, oss.Debian, logger), controller, configurator, logger)
}

func (t *template) Arch() app.Applicable {
	return app.NewApplication(checker, installer(t.Pkg, oss.Arch, logger), controller, configurator, logger)
}
