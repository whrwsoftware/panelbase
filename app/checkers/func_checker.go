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

package checkers

import "github.com/whrwsoftware/panelbase/app"

type CheckFunc func() (ok bool, err error)

type funcChecker struct {
	CheckFunc []CheckFunc
}

func FuncChecker(CheckFunc ...CheckFunc) *funcChecker { return &funcChecker{CheckFunc} }

func (c *funcChecker) Check(app.Manager) (ok bool, err error) {
	if fns := c.CheckFunc; fns != nil && len(fns) > 0 {
		for _, fn := range fns {
			if ok, err = fn(); err != nil || !ok {
				return false, err
			}
		}
	}
	return true, nil
}
