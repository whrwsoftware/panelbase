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

type Dpkg struct{ pkg string }

func NewDpkg(pkg string) *Dpkg                                   { return &Dpkg{pkg} }
func (d *Dpkg) run(v ...string) (out string, ok bool, err error) { return run("dpkg", v...) }
func (d *Dpkg) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("dpkg", v, outC, errC)
}
func (d *Dpkg) Version() (out string, ok bool, err error) { return d.run("--version") }
func (d *Dpkg) Search() (out string, ok bool, err error)  { return d.run("search", d.pkg) }
func (d *Dpkg) Install(outC, errC chan<- string) (ok bool, err error) {
	return d.start([]string{"-i", d.pkg}, outC, errC)
}
func (d *Dpkg) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return d.Install(outC, errC)
}
func (d *Dpkg) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return d.start([]string{"-r", d.pkg}, outC, errC)
}
