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

type Service struct{ name string }

func NewService(name string) *Service                               { return &Service{name} }
func (s *Service) run(v ...string) (out string, ok bool, err error) { return run("service", v...) }
func (s *Service) Start() (out string, ok bool, err error)          { return s.run("start", s.name) }
func (s *Service) Stop() (out string, ok bool, err error)           { return s.run("stop", s.name) }
func (s *Service) Restart() (out string, ok bool, err error)        { return s.run("restart", s.name) }
func (s *Service) Status() (out string, ok bool, err error)         { return s.run("status", s.name) }
