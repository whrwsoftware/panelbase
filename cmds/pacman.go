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

type pacman struct{ pkg string }

func Pacman(pkg string) *pacman                                    { return &pacman{pkg} }
func (a *pacman) run(v ...string) (out string, ok bool, err error) { return run("pacman", v...) }
func (a *pacman) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("pacman", v, outC, errC)
}
func (a *pacman) Version() (out string, ok bool, err error) { return a.run("--version") }
func (a *pacman) Search() (out string, ok bool, err error)  { return a.run("-S", "-s", a.pkg) }
func (a *pacman) Install(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"-S", "--noconfirm", a.pkg}, outC, errC)
}
func (a *pacman) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return a.Install(outC, errC)
}
func (a *pacman) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"-R", "-noconfirm", a.pkg}, outC, errC)
}
