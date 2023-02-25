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

package centos

type Template6 struct {
	Name, ServiceName string
	RegisterService   bool
	PreCheck          bool
	PreCheckFunc      func() (ok bool, err error)
}

func (t *Template6) Check() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Enable() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Disable() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Install() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Uninstall() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Configure(data any) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Start() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Stop() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Restart() (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Status() (s string, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *Template6) Ver() (v string, err error) {
	//TODO implement me
	panic("implement me")
}
