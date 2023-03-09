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

package managers

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"sync"

	appPkg "github.com/whrwsoftware/panelbase/app"
)

type fileManager struct {
	mu *sync.Mutex

	File string
}

func FileManager(file string) *fileManager { return FileManagerLoad(file, false) }

func FileManagerLoad(file string, load bool) *fileManager {
	_ = os.MkdirAll(path.Dir(file), 0644)
	if !load {
		_ = os.Remove(file)
	}
	f, _ := os.Open(file)
	if f != nil {
		f.Close()
	}
	if f == nil {
		_ = os.WriteFile(file, []byte("[]"), 0644)
	}
	return &fileManager{&sync.Mutex{}, file}
}

func (f *fileManager) Add(app *appPkg.Info) (err error) {
	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	} else if apps != nil {
		apps = append(apps, app)
	} else {
		apps = []*appPkg.Info{app}
	}
	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Update(app *appPkg.Info) (err error) {
	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findApp *appPkg.Info

	for _, a := range apps {
		if a.Pkg == app.Pkg {
			findApp = a
			break
		}
	}

	if findApp == nil {
		return errors.New("the app is not exists")
	}

	findApp.Name = app.Name
	findApp.AppName = app.AppName
	findApp.Icon = app.Icon
	findApp.Provider = app.Provider
	findApp.Description = app.Description
	findApp.Workdir = app.Workdir
	findApp.Type = app.Type

	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Delete(pkg string) (err error) {
	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findIndex = -1

	for i, a := range apps {
		if a.Pkg == pkg {
			findIndex = i
			break
		}
	}

	if findIndex == -1 {
		return errors.New("the app is not exists")
	}

	newApps := make([]*appPkg.Info, 0)

	if findIndex == 0 {
		newApps = append(newApps, apps[1:]...)
	} else if findIndex == len(apps)-1 {
		newApps = append(newApps, apps[:len(apps)-1]...)
	} else {
		newApps = append(newApps, apps[:findIndex]...)
		newApps = append(newApps, apps[findIndex+1:]...)
	}

	if buf, bErr := json.Marshal(newApps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) AddTag(pkg string, tag string) (err error) {
	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findApp *appPkg.Info

	for _, a := range apps {
		if a.Pkg == pkg {
			findApp = a
			break
		}
	}

	if findApp == nil {
		return errors.New("the app is not exists")
	}

	if findApp.Tag == nil {
		findApp.Tag = map[string]string{tag: ""}
	} else {
		findApp.Tag[tag] = ""
	}

	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) DeleteTag(pkg string, tag string) (err error) {
	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findApp *appPkg.Info

	for _, a := range apps {
		if a.Pkg == pkg {
			findApp = a
			break
		}
	}

	if findApp == nil {
		return errors.New("the app is not exists")
	}

	if findApp.Tag != nil {
		delete(findApp.Tag, tag)
	}

	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Installed(pkg string) (err error) {
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findApp *appPkg.Info

	for _, a := range apps {
		if a.Pkg == pkg {
			findApp = a
			break
		}
	}

	if findApp == nil {
		return errors.New("the app is not exists")
	}

	findApp.Installed = true

	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()

	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Uninstalled(pkg string) (err error) {
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var findApp *appPkg.Info

	for _, a := range apps {
		if a.Pkg == pkg {
			findApp = a
			break
		}
	}

	if findApp == nil {
		return errors.New("the app is not exists")
	}

	findApp.Installed = false

	if !f.mu.TryLock() {
		return errors.New("the file is busy now")
	}
	defer f.mu.Unlock()

	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Required(required ...*appPkg.Required) (err error) {
	if required == nil || len(required) <= 0 {
		return nil
	}
	apps, aErr := f.FindAll()
	if aErr != nil {
		err = aErr
		return
	}
	if apps == nil || len(apps) <= 0 {
		return errors.New("the app is not exists")
	}

	var (
		pkgMap  = map[string]*appPkg.Info{}
		nameMap = map[string]struct{}{}
	)

	for _, p := range apps {
		pkgMap[p.Pkg] = p
		nameMap[p.Name] = struct{}{}
	}

	var (
		notExistApp     []string
		notInstalledApp []string
	)

loop:
	for _, reqPkg := range required {
		if _, ok := nameMap[reqPkg.Name]; !ok {
			// not exists
			notExistApp = append(notExistApp, reqPkg.Name)
			continue loop
		}
		for _, app := range apps {
			if app.Name == reqPkg.Name && app.Installed && app.VersionId >= reqPkg.VersionId {
				continue loop
			}
		}
		notInstalledApp = append(notInstalledApp, reqPkg.Name+"@"+reqPkg.Version)
	}

	if len(notExistApp) > 0 || len(notInstalledApp) > 0 {
		return &appPkg.CheckError{NotExistApp: notExistApp, NotInstallApp: notInstalledApp}
	}

	return
}

func (f *fileManager) FindAll() (apps []*appPkg.Info, err error) {
	if readFile, rErr := os.ReadFile(f.File); rErr != nil {
		err = rErr
		return
	} else {
		err = json.Unmarshal(readFile, &apps)
	}
	return
}

func (f *fileManager) FindByPkg(pkg string) (app *appPkg.Info, err error) {
	if apps, aErr := f.FindAll(); aErr != nil {
		err = aErr
		return
	} else if apps != nil && len(apps) > 0 {
		for _, a := range apps {
			if a.Pkg == pkg {
				app = a
				return
			}
		}
	}
	return
}

func (f *fileManager) FindsByType(typ appPkg.Type) (app []*appPkg.Info, err error) {
	if apps, aErr := f.FindAll(); aErr != nil {
		err = aErr
		return
	} else if apps != nil && len(apps) > 0 {
		for _, a := range apps {
			if a.Type == typ {
				app = append(app, a)
				return
			}
		}
	}
	return
}

func (f *fileManager) FindsByName(name string) (apps []*appPkg.Info, err error) {
	if fApps, aErr := f.FindAll(); aErr != nil {
		err = aErr
		return
	} else if fApps != nil && len(fApps) > 0 {
		for _, a := range fApps {
			if a.Name == name {
				apps = append(apps, a)
			}
		}
	}
	return
}

func (f *fileManager) FindsByTag(tag string) (apps []*appPkg.Info, err error) {
	if fApps, aErr := f.FindAll(); aErr != nil {
		err = aErr
		return
	} else if fApps != nil && len(fApps) > 0 {
	loop:
		for _, a := range fApps {
			if a.Tag != nil && len(a.Tag) > 0 {
				if _, ok := a.Tag[tag]; ok {
					apps = append(apps, a)
					continue loop
				}
			}
		}
	}
	return
}
