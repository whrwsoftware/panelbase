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

package appmanager

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
	"sync"
)

type fileManager struct {
	mu *sync.Mutex

	File string
}

func UseFileManager(file string)                { mgr = FileManager(file) }
func UseFileManagerLoad(file string, load bool) { mgr = FileManagerLoad(file, load) }
func FileManager(file string) Manager           { return FileManagerLoad(file, false) }
func FileManagerLoad(file string, load bool) Manager {
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

func (f *fileManager) Add(app *App) (err error) {
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
		apps = []*App{app}
	}
	if buf, bErr := json.Marshal(apps); bErr != nil {
		err = bErr
		return bErr
	} else {
		err = os.WriteFile(f.File, buf, 0600)
	}
	return
}

func (f *fileManager) Update(app *App) (err error) {
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

	var findApp *App

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

	newApps := make([]*App, 0)

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

	var findApp *App

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

	var findApp *App

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

	var findApp *App

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

	var findApp *App

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

func (f *fileManager) Check(requiredApp ...RequiredApp) (err error) {
	if requiredApp == nil || len(requiredApp) <= 0 {
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
		pkgMap  = map[string]*App{}
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
	for _, reqPkg := range requiredApp {
		ootVer := reqPkg.OneOfThemVersion
		if _, ok := nameMap[reqPkg.Name]; !ok {
			// not exists
			notExistApp = append(notExistApp, reqPkg.Name)
			continue loop
		}
		if ootVer == nil {
			ootVer = []string{}
		}
		if len(ootVer) <= 0 {
			for _, app := range apps {
				if app.Name == reqPkg.Name && app.Installed {
					continue loop
				}
			}
		} else {
			for _, ver := range ootVer {
				for _, app := range apps {
					if app.Name == reqPkg.Name && app.Version == ver && app.Installed {
						continue loop
					}
				}
			}
		}
		// not installed
		if len(ootVer) == 0 {
			notInstalledApp = append(notInstalledApp, reqPkg.Name)
		} else {
			notInstalledApp = append(notInstalledApp, reqPkg.Name+"("+strings.Join(ootVer, ",")+")")
		}
	}

	if len(notExistApp) > 0 || len(notInstalledApp) > 0 {
		return &CheckError{notExistApp, notInstalledApp}
	}

	return
}

func (f *fileManager) FindAll() (apps []*App, err error) {
	if readFile, rErr := os.ReadFile(f.File); rErr != nil {
		err = rErr
		return
	} else {
		err = json.Unmarshal(readFile, &apps)
	}
	return
}

func (f *fileManager) FindByPkg(pkg string) (app *App, err error) {
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

func (f *fileManager) FindsByType(typ Type) (app []*App, err error) {
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

func (f *fileManager) FindsByName(name string) (apps []*App, err error) {
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

func (f *fileManager) FindsByTag(tag string) (apps []*App, err error) {
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
