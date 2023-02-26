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

type apt struct{}

func Apt(unit string) *apt                                  { return &apt{} }
func apt0(v ...string) (out string, err error)              { return run("apt", v...) }
func apt1(v []string, outC, errC chan<- string) (err error) { return start("apt", v, outC, errC) }
func (a *apt) Version() (out string, err error)             { return apt0("--version") }
func (a *apt) Search(v string) (out string, err error)      { return apt0("search", v) }
func (a *apt) Install(v string, outC, errC chan<- string) (err error) {
	return apt1([]string{"install -y", v}, outC, errC)
}
func (a *apt) Uninstall(v string, outC, errC chan<- string) (err error) {
	return apt1([]string{"uninstall -y", v}, outC, errC)
}
