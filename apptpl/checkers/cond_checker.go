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

import "errors"

var (
	ErrCondChecker = errors.New("cond_checker: cond check not pass")
)

type (
	ValueFunc   func() (data any)
	CondFunc    func(expect, data any) bool
	condChecker struct {
		ValueFunc
		CondFunc
		Expect any
	}
)

func CondChecker(valueFunc ValueFunc, condFunc CondFunc, expect any) *condChecker {
	return &condChecker{ValueFunc: valueFunc, CondFunc: condFunc, Expect: expect}
}

func (c *condChecker) Check() (ok bool, err error) {
	if fn := c.ValueFunc; fn != nil {
		if cFn := c.CondFunc; cFn != nil {
			if ok = cFn(c.Expect, fn()); !ok {
				err = ErrCondChecker
				return
			}
		}
	}
	return true, nil
}
