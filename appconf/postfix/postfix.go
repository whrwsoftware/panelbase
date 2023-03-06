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

package postfix

import (
	"embed"
	_ "embed"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed main.cf master.cf
	fs embed.FS
)

const (
	NameMainCf   = "main.cf"
	NameMasterCf = "master.cf"
)

const (
	DistMainCf   = "/etc/postfix/main.cf"
	DistMasterCf = "/etc/postfix/master.cf"
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindMainCf   = appconf.NewConfBind[any](fs, NameMainCf, DistMainCf, perm)
	ConfBindMasterCf = appconf.NewConfBind[any](fs, NameMasterCf, DistMasterCf, perm)
)

var ConfBinds = []*appconf.ConfBind[any]{
	ConfBindMainCf,
	ConfBindMasterCf,
}

type (
	OptMainCf struct {
		MyHostname  string
		MyDomain    string
		TLS         bool
		TLSKeyFile  string
		TLSCertFile string
	}
	OptMasterCf struct{}
)
