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

type systemctl struct{ unit string }

func Systemctl(unit string) *systemctl                { return &systemctl{unit} }
func systemctl0(v ...string) (out string, err error)  { return run("systemctl", v...) }
func (s *systemctl) Version() (out string, err error) { return systemctl0("--version") }
func (s *systemctl) Enable() (out string, err error)  { return systemctl0("enable", s.unit) }
func (s *systemctl) Disable() (out string, err error) { return systemctl0("disable", s.unit) }
func (s *systemctl) Start() (out string, err error)   { return systemctl0("start", s.unit) }
func (s *systemctl) Stop() (out string, err error)    { return systemctl0("stop", s.unit) }
func (s *systemctl) Restart() (out string, err error) { return systemctl0("restart", s.unit) }
func (s *systemctl) Status() (out string, err error)  { return systemctl0("status", s.unit) }
