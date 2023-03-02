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

type Systemctl struct{ unit string }

func NewSystemctl(unit string) *Systemctl                             { return &Systemctl{unit} }
func (s *Systemctl) run(v ...string) (out string, ok bool, err error) { return run("systemctl", v...) }
func (s *Systemctl) Version() (out string, ok bool, err error)        { return s.run("--version") }
func (s *Systemctl) Enable() (out string, ok bool, err error)         { return s.run("enable", s.unit) }
func (s *Systemctl) Disable() (out string, ok bool, err error)        { return s.run("disable", s.unit) }
func (s *Systemctl) Start() (out string, ok bool, err error)          { return s.run("start", s.unit) }
func (s *Systemctl) Stop() (out string, ok bool, err error)           { return s.run("stop", s.unit) }
func (s *Systemctl) Restart() (out string, ok bool, err error)        { return s.run("restart", s.unit) }
func (s *Systemctl) Status() (out string, ok bool, err error)         { return s.run("status", s.unit) }
