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

package latest

import (
	"bytes"
	_ "embed"
	"github.com/whrwsoftware/panelbase/zvars/oss"
	"text/template"
)

var (
	//go:embed install.sh
	install string
	//go:embed uninstall.sh
	uninstall string
	//go:embed reinstall.sh
	reinstall string
)

type fs struct{ oss.OS }

func FS(OS oss.OS) *fs          { return &fs{OS: OS} }
func (f *fs) Install() string   { return f.exec(install) }
func (f *fs) Uninstall() string { return f.exec(uninstall) }
func (f *fs) Reinstall() string { return f.exec(reinstall) }
func (f *fs) exec(tpl string) string {
	tt, _ := template.New("").Parse(tpl)
	var w bytes.Buffer
	_ = tt.Execute(&w, f)
	return w.String()
}
