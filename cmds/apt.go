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

type apt struct{ pkg string }

func Apt(pkg string) *apt                                       { return &apt{pkg} }
func (a *apt) run(v ...string) (out string, ok bool, err error) { return run("apt", v...) }
func (a *apt) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("apt", v, outC, errC)
}
func (a *apt) Version() (out string, ok bool, err error) { return a.run("--version") }
func (a *apt) Search() (out string, ok bool, err error)  { return a.run("search", a.pkg) }
func (a *apt) Install(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"install", "-y", a.pkg}, outC, errC)
}
func (a *apt) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"reinstall", "-y", a.pkg}, outC, errC)
}
func (a *apt) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"remove", "-y", a.pkg}, outC, errC)
}
