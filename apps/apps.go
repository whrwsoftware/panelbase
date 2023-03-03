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
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type GetAppFunc func(outC, errC chan string) app.Applicable

type GetAppTemplate interface {
	CentOS6() GetAppFunc
	CentOS7() GetAppFunc
	CentOS8() GetAppFunc
	Ubuntu() GetAppFunc
	Debian() GetAppFunc
	Arch() GetAppFunc
}

func GetApp(get GetAppTemplate, os oss.OS, outC, errC chan string) (app app.Applicable) {
	switch os {
	case oss.CentOS6:
		if fn := get.CentOS6(); fn != nil {
			return fn(outC, errC)
		}
	case oss.CentOS7:
		if fn := get.CentOS7(); fn != nil {
			return fn(outC, errC)
		}
	case oss.CentOS8:
		if fn := get.CentOS8(); fn != nil {
			return fn(outC, errC)
		}
	case oss.Ubuntu:
		if fn := get.Ubuntu(); fn != nil {
			return fn(outC, errC)
		}
	case oss.Debian:
		if fn := get.Debian(); fn != nil {
			return fn(outC, errC)
		}
	case oss.Arch:
		if fn := get.Arch(); fn != nil {
			return fn(outC, errC)
		}
	}
	return
}
