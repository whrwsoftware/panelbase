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

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

type defaultConfigurator struct {
	ConfBind []*appconf.ConfBind[any]

	Loader  app.ConfigurationLoader
	Cleaner app.ConfigurationCleaner
}

func DefaultConfigurator(ConfBind []*appconf.ConfBind[any]) *defaultConfigurator {
	var cleanFile []string
	if cbs := ConfBind; cbs != nil {
		for _, cb := range cbs {
			cleanFile = append(cleanFile, cb.Dist)
		}
	}
	return &defaultConfigurator{ConfBind, DefaultConfigurationLoader(), DefaultConfigurationCleaner(cleanFile)}
}

func DefaultConfiguratorParam(ConfBind []*appconf.ConfBind[any], configurationLoader app.ConfigurationLoader, configurationCleaner app.ConfigurationCleaner) app.Configurator {
	return &defaultConfigurator{ConfBind, configurationLoader, configurationCleaner}
}

func (d *defaultConfigurator) Configure(m map[string]any) (err error) {
	cbs := d.ConfBind
	if cbs != nil && len(cbs) > 0 {
		for _, cb := range cbs {
			if data, ok := m[cb.Name]; ok {
				err = cb.Configure(data)
			} else {
				err = cb.Configure(map[string]any{})
			}
			if err != nil {
				return
			}
		}
	}
	return
}

func (d *defaultConfigurator) Load(name string) (v string, err error) { return d.Loader.Load(name) }

func (d *defaultConfigurator) Clean() (err error) { return d.Cleaner.Clean() }

type defaultConfigurationLoader struct{}

func DefaultConfigurationLoader() app.ConfigurationLoader { return &defaultConfigurationLoader{} }

func (d *defaultConfigurationLoader) Load(name string) (v string, err error) {
	readFile, rErr := os.ReadFile(name)
	if rErr != nil {
		err = rErr
		return
	}
	if readFile != nil {
		v = string(readFile[:])
	}
	return
}

type defaultConfigurationCleaner struct{ CleanFile []string }

func DefaultConfigurationCleaner(cleanFile []string) app.ConfigurationCleaner {
	return &defaultConfigurationCleaner{cleanFile}
}

func (d *defaultConfigurationCleaner) Clean() (err error) {
	if d.CleanFile != nil {
		for _, f := range d.CleanFile {
			err = os.Remove(f)
		}
	}
	return
}
