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

type rpm struct{}

func Rpm() *rpm                                         { return &rpm{} }
func rpm0(v ...string) (out string, ok bool, err error) { return run("rpm", v...) }
func rpm1(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("rpm", v, outC, errC)
}
func (r *rpm) Version() (out string, ok bool, err error)        { return rpm0("--version") }
func (r *rpm) Search(v string) (out string, ok bool, err error) { return rpm0("search", v) }
func (r *rpm) Install(v string, outC, errC chan<- string) (ok bool, err error) {
	return rpm1([]string{"install", "-ivh", v}, outC, errC)
}
func (r *rpm) Uninstall(v string, outC, errC chan<- string) (ok bool, err error) {
	return rpm1([]string{"remove", "-evh", v}, outC, errC)
}
