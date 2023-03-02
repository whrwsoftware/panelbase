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

type Apt struct{ pkg string }

func NewApt(pkg string) *Apt                                    { return &Apt{pkg} }
func (a *Apt) run(v ...string) (out string, ok bool, err error) { return run("apt", v...) }
func (a *Apt) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("apt", v, outC, errC)
}
func (a *Apt) Version() (out string, ok bool, err error) { return a.run("--version") }
func (a *Apt) Search() (out string, ok bool, err error)  { return a.run("search", a.pkg) }
func (a *Apt) Install(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"install", "-y", a.pkg}, outC, errC)
}
func (a *Apt) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"reinstall", "-y", a.pkg}, outC, errC)
}
func (a *Apt) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"remove", "-y", a.pkg}, outC, errC)
}
