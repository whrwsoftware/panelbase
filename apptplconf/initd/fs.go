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

package initd

import (
	"embed"
	"github.com/whrwsoftware/panelbase/apptplconf"
)

//go:embed bash.sh
var fs embed.FS

const initDRoot = "/etc/init.d/"

type Args struct {
	App     string
	Command string
	Option  string
	PidFile string
	LogFile string
}

func InitD(a Args) (err error) { return apptplconf.Gen(fs, "bash.sh", initDRoot+a.App, a, 0700) }
