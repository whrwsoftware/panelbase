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

type Dnf struct{ pkg string }

func NewDnf(pkg string) *Dnf                                    { return &Dnf{pkg} }
func (d *Dnf) run(v ...string) (out string, ok bool, err error) { return run("dnf", v...) }
func (d *Dnf) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("dnf", v, outC, errC)
}
func (d *Dnf) Version() (out string, ok bool, err error) { return d.run("--version") }
func (d *Dnf) Search() (out string, ok bool, err error)  { return d.run("search", d.pkg) }
func (d *Dnf) Install(outC, errC chan<- string) (ok bool, err error) {
	return d.start([]string{"install", "-y", d.pkg}, outC, errC)
}
func (d *Dnf) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return d.start([]string{"reinstall", "-y", d.pkg}, outC, errC)
}
func (d *Dnf) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return d.start([]string{"remove", "-y", d.pkg}, outC, errC)
}
