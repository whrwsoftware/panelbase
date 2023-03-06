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

package appconf

import (
	"bytes"
	"embed"
	"os"
	"text/template"
)

type ConfBind[T any] struct {
	FS         embed.FS
	Name, Dist string
	Perm       os.FileMode
}

func NewConfBind[T any](FS embed.FS, name string, dist string, perm os.FileMode) *ConfBind[T] {
	return &ConfBind[T]{FS: FS, Name: name, Dist: dist, Perm: perm}
}

func (cb *ConfBind[T]) Configure(data T) (err error) {
	buf, bErr := cb.FS.ReadFile(cb.Name)
	if bErr != nil {
		err = bErr
		return
	}
	tpl, tErr := template.New("").Parse(string(buf[:]))
	if tErr != nil {
		err = tErr
		return
	}
	var w bytes.Buffer
	if err = tpl.Execute(&w, data); err != nil {
		return
	}

	if err = os.WriteFile(cb.Dist, w.Bytes(), cb.Perm); err != nil {
		return
	}

	return
}
