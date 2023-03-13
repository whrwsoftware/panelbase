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

package configurators

type noConfigurator struct{}

func NoConfigurator() *noConfigurator                                 { return &noConfigurator{} }
func (n *noConfigurator) Configure(map[string]any) (err error)        { return nil }
func (n *noConfigurator) ConfigureFile(map[string]string) (err error) { return nil }
func (n *noConfigurator) Load(string) (v string, err error)           { return "", nil }
func (n *noConfigurator) Clean() (err error)                          { return nil }
