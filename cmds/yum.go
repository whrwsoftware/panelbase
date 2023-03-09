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

type yum struct{ pkg string }

func Yum(pkg string) *yum                                       { return &yum{pkg} }
func (y *yum) run(v ...string) (out string, ok bool, err error) { return run("yum", v...) }
func (y *yum) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("yum", v, outC, errC)
}
func (y *yum) Version() (out string, ok bool, err error) { return y.run("--version") }
func (y *yum) Search() (out string, ok bool, err error)  { return y.run("search", y.pkg) }
func (y *yum) Install(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"install", "-y", y.pkg}, outC, errC)
}
func (y *yum) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"reinstall", "-y", y.pkg}, outC, errC)
}
func (y *yum) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return y.start([]string{"remove", "-y", y.pkg}, outC, errC)
}
