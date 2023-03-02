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

type Yum struct{ pkg string }

func NewYum(pkg string) *Yum                                    { return &Yum{pkg} }
func (y *Yum) run(v ...string) (out string, ok bool, err error) { return run("yum", v...) }
func (y *Yum) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("yum", v, outC, errC)
}
func (y *Yum) Version() (out string, ok bool, err error) { return y.run("--version") }
func (y *Yum) Search() (out string, ok bool, err error)  { return y.run("search", y.pkg) }
func (y *Yum) Install(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"install", "-y", y.pkg}, outC, errC)
}
func (y *Yum) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"reinstall", "-y", y.pkg}, outC, errC)
}
func (y *Yum) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"remove", "-y", y.pkg}, outC, errC)
}
