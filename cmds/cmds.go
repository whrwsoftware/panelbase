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

package cmds

import "github.com/whrwsoftware/panelbase/cmd"

func run(n string, v ...string) (out string, ok bool, err error) {
	out, _, ok, err = cmd.Run(n, v...)
	return
}
func start(n string, v []string, outC, errC chan<- string) (ok bool, err error) {
	return cmd.Start(n, v, outC, errC, nil, nil)
}
