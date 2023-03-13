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

package php

import (
	"embed"
	"fmt"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed conf/php.ini@82 conf/www.conf@82 conf/php.ini@81 conf/www.conf@81 conf/php.ini@80 conf/www.conf@80 conf/php.ini@74 conf/www.conf@74 conf/php.ini@73 conf/www.conf@73 conf/php.ini@72 conf/www.conf@72 conf/php.ini@71 conf/www.conf@71 conf/php.ini@70 conf/www.conf@70 conf/php.ini@56 conf/www.conf@56 conf/php.ini@55 conf/www.conf@55 conf/php.ini@54 conf/www.conf@54
	fs embed.FS
)

var (
	NamePhpIni  = func(v int) string { return fmt.Sprintf("conf/php.ini@%d", v) }
	NameWwwConf = func(v int) string { return fmt.Sprintf("conf/www.conf@%d", v) }
)

var (
	DistPhpIni = func(v int) string {
		if v < 70 {
			return fmt.Sprintf("/opt/remi/php%d/root/etc/php.ini", v)
		}
		return fmt.Sprintf("/etc/opt/remi/php%d/php.ini", v)
	}
	DistWwwConf = func(v int) string {
		if v < 70 {
			return fmt.Sprintf("/opt/remi/php%d/root/etc/php-fpm.d/www.conf", v)
		}
		return fmt.Sprintf("/etc/opt/remi/php%d/php-fpm.d/www.conf", v)
	}
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindPhpIni  = func(v int) *appconf.ConfBind { return appconf.NewConfBind(fs, NamePhpIni(v), DistPhpIni(v), perm) }
	ConfBindWwwConf = func(v int) *appconf.ConfBind { return appconf.NewConfBind(fs, NameWwwConf(v), DistWwwConf(v), perm) }
)

var ConfBinds = func(v int) []*appconf.ConfBind {
	return []*appconf.ConfBind{
		ConfBindPhpIni(v),
		ConfBindWwwConf(v),
	}
}

func ConfBindMap(v int, optPhpIni OptPhpIni, optWwwConf OptWwwConf) map[string]any {
	return map[string]any{
		NamePhpIni(v):  optPhpIni,
		NameWwwConf(v): optWwwConf,
	}
}

type (
	OptPhpIni  struct{}
	OptWwwConf struct{}
)
