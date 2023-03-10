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

package apps

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/app/configurators"
	"github.com/whrwsoftware/panelbase/appmanager"
	"strings"
)

type GroupedApp struct {
	bundles []*Bundle

	app.ConfigurationLoader
}

type Bundle struct {
	app.Applicable
	Name, Ver string
}

func NewBundle(applicable app.Applicable, ver string) *Bundle {
	return &Bundle{Applicable: applicable, Ver: ver}
}

func NewGroupApp(bundle ...*Bundle) *GroupedApp {
	return &GroupedApp{bundle, configurators.DefaultConfigurationLoader()}
}

func (a *GroupedApp) Check(manager appmanager.Manager) (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Check(manager); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Install(string) (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Install(ba.Ver); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Uninstall(string) (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Uninstall(ba.Ver); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Reinstall(string) (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Reinstall(ba.Ver); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Enable() (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Enable(); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Disable() (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Disable(); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Start() (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Start(); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Stop() (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Stop(); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Restart() (ok bool, err error) {
	for _, ba := range a.bundles {
		if ok, err = ba.Restart(); err != nil {
			return
		}
	}
	return true, nil
}

func (a *GroupedApp) Status() (st string, ok bool, err error) {
	for _, ba := range a.bundles {
		if st, ok, err = ba.Status(); err != nil {
			return
		}
	}
	return "", true, nil
}
func (a *GroupedApp) Configure(m map[string]any) (err error) {
	if m != nil || len(m) <= 0 {
		return nil
	}
	for _, bApp := range a.bundles {
		if bOpt, ok := m[bApp.Name]; ok {
			if bOptAny, okk := bOpt.(map[string]any); okk {
				if err = bApp.Configure(bOptAny); err != nil {
					return
				}
			}
		}
	}
	return
}

func (a *GroupedApp) File() string {
	var files []string
	for _, ba := range a.bundles {
		files = append(files, ba.File())
	}
	return strings.Join(files, ",")
}

func (a *GroupedApp) Truncate() (err error) {
	for _, ba := range a.bundles {
		if err = ba.Truncate(); err != nil {
			return
		}
	}
	return
}

func (a *GroupedApp) Clean() (err error) {
	for _, ba := range a.bundles {
		if err = ba.Clean(); err != nil {
			return
		}
	}
	return
}
