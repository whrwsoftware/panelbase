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

package apptpl

import (
	"fmt"
	"strings"
)

type ErrPreCond struct {
	items []ErrPreCondItem
}

type ErrPreCondItem struct {
	Name      string
	Version   string
	Installed bool
}

func (e *ErrPreCond) Error() string {
	buf := &strings.Builder{}
	_, _ = buf.WriteString("Precondition list")
	_ = buf.WriteByte('\n')
	_, _ = buf.WriteString("---")
	installed := make([]ErrPreCondItem, 0)
	notInstalled := make([]ErrPreCondItem, 0)
	for _, ii := range e.items {
		if ii.Installed {
			installed = append(installed, ii)
		} else {
			notInstalled = append(notInstalled, ii)
		}
	}
	for _, ii := range installed {
		_, _ = buf.WriteString(fmt.Sprintf("%s %s", ii.Name, ii.Version))
	}
	for _, ii := range notInstalled {
		_, _ = buf.WriteString(fmt.Sprintf("%s not installed", ii.Name))
	}
	return buf.String()
}
