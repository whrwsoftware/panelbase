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

package svc

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/tpl"
	"github.com/whrwsoftware/panelbase/apps"
)

type StdSvcTemplate struct {
	Name string
	Ver  string
	Pkg  string
	app.Checker
	app.Configurator
}

func (s *StdSvcTemplate) CentOS6() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return tpl.YumApp(s.Name, s.Ver, s.Pkg, outC, errC, s.Checker, s.Configurator)
	}
}

func (s *StdSvcTemplate) CentOS7() apps.GetAppFunc { return s.CentOS6() }

func (s *StdSvcTemplate) CentOS8() apps.GetAppFunc { return s.CentOS7() }

func (s *StdSvcTemplate) Ubuntu() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return tpl.AptApp(s.Name, s.Ver, s.Pkg, outC, errC, s.Checker, s.Configurator)
	}
}

func (s *StdSvcTemplate) Debian() apps.GetAppFunc { return s.Ubuntu() }

func (s *StdSvcTemplate) Arch() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return tpl.PacmanApp(s.Name, s.Ver, s.Pkg, outC, errC, s.Checker, s.Configurator)
	}
}
