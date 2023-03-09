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

type apk struct{ pkg string }

func Apk(pkg string) *apk                                       { return &apk{pkg} }
func (a *apk) run(v ...string) (out string, ok bool, err error) { return run("apk", v...) }
func (a *apk) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("apk", v, outC, errC)
}
func (a *apk) Version() (out string, ok bool, err error) { return a.run("--version") }
func (a *apk) Search() (out string, ok bool, err error)  { return a.run("search", a.pkg) }
func (a *apk) Install(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"add", a.pkg}, outC, errC)
}
func (a *apk) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return a.Install(outC, errC)
}
func (a *apk) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return a.start([]string{"del", a.pkg}, outC, errC)
}
