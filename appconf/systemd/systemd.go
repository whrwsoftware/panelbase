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

package systemd

import (
	"embed"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed system.service
	fs embed.FS
)

const (
	RootService = "/etc/systemd/system/multi-user.target.wants"
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindSystemService = func(name string) *appconf.ConfBind[any] {
		return appconf.NewConfBind[any](fs, name+".service", RootService+name+".service", perm)
	}
)

type (
	OptSystemService struct {
		Unit    OptSystemServiceUint
		Service OptSystemServiceService
	}
	OptSystemServiceUint struct {
		Description string
	}
	OptSystemServiceService struct {
		Type               string
		PIDFile            bool
		PIDFileVal         string
		EnvironmentFile    bool
		EnvironmentFileVal string
		ExecStart          bool
		ExecStartVal       string
		ExecReload         bool
		ExecReloadVal      string
		ExecStop           bool
		ExecStopVal        string
	}
)
