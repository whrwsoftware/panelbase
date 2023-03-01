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

package configurators

import (
	"errors"
	"github.com/whrwsoftware/panelbase/apptpl"
	"os"
)

var (
	ErrNotFoundConf = errors.New("configurator: not found conf")
)

type (
	file struct {
		ConfMapping map[string]conf
	}
	conf struct {
		Path string
		Perm os.FileMode
	}
)

func Conf(path string, perm os.FileMode) conf              { return conf{Path: path, Perm: perm} }
func File(confMapping map[string]conf) apptpl.Configurator { return &file{ConfMapping: confMapping} }

func (f *file) Configure(cfg apptpl.Cfg) (err error) {
	confV, ok := f.ConfMapping[cfg.Name]
	if !ok {
		return ErrNotFoundConf
	}
	return os.WriteFile(confV.Path, []byte(cfg.Data), confV.Perm)
}
