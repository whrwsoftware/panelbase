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

type rpm struct{ pkg string }

func Rpm(pkg string) *rpm                                       { return &rpm{pkg} }
func (r *rpm) run(v ...string) (out string, ok bool, err error) { return run("rpm", v...) }
func (r *rpm) start(v []string, outC, errC chan<- string) (ok bool, err error) {
	return start("rpm", v, outC, errC)
}
func (r *rpm) Version() (out string, ok bool, err error) { return r.run("--version") }
func (r *rpm) Search() (out string, ok bool, err error)  { return r.run("search", r.pkg) }
func (r *rpm) Install(outC, errC chan<- string) (ok bool, err error) {
	return r.start([]string{"-ivh", r.pkg}, outC, errC)
}
func (r *rpm) Reinstall(outC, errC chan<- string) (ok bool, err error) {
	return r.Install(outC, errC)
}
func (r *rpm) Uninstall(outC, errC chan<- string) (ok bool, err error) {
	return r.start([]string{"-evh", r.pkg}, outC, errC)
}
