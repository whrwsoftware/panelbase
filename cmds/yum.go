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

type yum struct{}

func Yum() *yum                                         { return &yum{} }
func yum0(v ...string) (out string, ok bool, err error) { return run("yum", v...) }
func yum1(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("yum", v, outC, errC)
}
func (y *yum) Version() (out string, ok bool, err error)        { return yum0("--version") }
func (y *yum) Search(v string) (out string, ok bool, err error) { return yum0("search", v) }
func (y *yum) Install(v string, outC, errC chan<- string) (ok bool, err error) {
	return yum1([]string{"install", "-y", v}, outC, errC)
}
func (y *yum) LocalInstall(v string, outC, errC chan<- string) (ok bool, err error) {
	return yum1([]string{"localinstall", "-y", v}, outC, errC)
}
func (y *yum) Uninstall(v string, outC, errC chan<- string) (ok bool, err error) {
	return yum1([]string{"remove", "-y", v}, outC, errC)
}
