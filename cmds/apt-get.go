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

type AptGet struct{ pkg string }

func NewAptGet(pkg string) *AptGet                                 { return &AptGet{pkg} }
func (a *AptGet) run(v ...string) (out string, ok bool, err error) { return run("apt-get", v...) }
func (a *AptGet) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("apt-get", v, outC, errC)
}
func (a *AptGet) Version() (out string, ok bool, err error) { return a.run("--version") }
func (a *AptGet) Search() (out string, ok bool, err error)  { return run("apt-cache", a.pkg) }
func (a *AptGet) Install(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"install -y", a.pkg}, outC, errC)
}
func (a *AptGet) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"reinstall -y", a.pkg}, outC, errC)
}
func (a *AptGet) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"remove -y", a.pkg}, outC, errC)
}