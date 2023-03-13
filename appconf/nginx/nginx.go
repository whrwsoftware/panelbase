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
	"embed"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed nginx.conf conf.d/default.conf
	fs embed.FS
)

const (
	NameNginxConf        = "nginx.conf"
	NameConfDDefaultConf = "conf.d/default.conf"
)

const (
	DistNginxConf        = "/etc/nginx/nginx.conf"
	DistConfDDefaultConf = "/etc/nginx/conf.d/default.conf"
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindNginxConf        = appconf.NewConfBind(fs, NameNginxConf, DistNginxConf, perm)
	ConfBindConfDDefaultConf = appconf.NewConfBind(fs, NameConfDDefaultConf, DistConfDDefaultConf, perm)
)

var ConfBinds = []*appconf.ConfBind{
	ConfBindNginxConf,
	ConfBindConfDDefaultConf,
}

func ConfBindMap(optNginxConf OptNginxConf, optConfDDefaultConf OptConfDDefaultConf) map[string]any {
	return map[string]any{
		NameNginxConf:        optNginxConf,
		NameConfDDefaultConf: optConfDDefaultConf,
	}
}

type (
	OptNginxConf        struct{}
	OptConfDDefaultConf struct{}
)
