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
	"github.com/whrwsoftware/panelbase/app"
	"os"
	"testing"
)

var testPath = "/path/to/app.json"

func UseFileManager(file string)                       { app.SetManager(FileManager(file)) }
func Add(info *app.Info) (err error)                   { return app.GetManager().Add(info) }
func Update(info *app.Info) (err error)                { return app.GetManager().Update(info) }
func Delete(pkg string) (err error)                    { return app.GetManager().Delete(pkg) }
func AddTag(pkg string, tag string) (err error)        { return app.GetManager().AddTag(pkg, tag) }
func DeleteTag(pkg string, tag string) (err error)     { return app.GetManager().DeleteTag(pkg, tag) }
func Installed(pkg string) (err error)                 { return app.GetManager().Installed(pkg) }
func Uninstalled(pkg string) (err error)               { return app.GetManager().Uninstalled(pkg) }
func Required(required ...*app.Required) (err error)   { return app.GetManager().Required(required...) }
func FindAll() (infos []*app.Info, err error)          { return app.GetManager().FindAll() }
func FindByPkg(pkg string) (info *app.Info, err error) { return app.GetManager().FindByPkg(pkg) }
func FindsByType(typ app.Type) (info []*app.Info, err error) {
	return app.GetManager().FindsByType(typ)
}
func FindsByName(name string) (infos []*app.Info, err error) {
	return app.GetManager().FindsByName(name)
}
func FindsByTag(tag string) (infos []*app.Info, err error) { return app.GetManager().FindsByTag(tag) }

func Test_fileManager_Add(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"","Version":"","VersionId":0,"Pkg":"","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":null,"Installed":false}]`
	if err := Add(&app.Info{}); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_Update(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":null,"Installed":false}]`
	if err := Add(&app.Info{Pkg: "app@1.0.0"}); err != nil {
		t.Error(err.Error())
	}
	if err := Update(&app.Info{Pkg: "app@1.0.0", Name: "app"}); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_AddTag(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0"}); err != nil {
		t.Error(err.Error())
	}
	if err := AddTag("app@1.0.0", "env"); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_Required(t *testing.T) {
	UseFileManager(testPath)
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", VersionId: 1, Pkg: "app@1.0.0", Tag: map[string]string{"env": "hello"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&app.Info{Name: "app", Version: "1.0.1", VersionId: 2, Pkg: "app@1.0.1", Tag: map[string]string{"env": "world"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&app.Info{Name: "app2", Version: "1.0.0", VersionId: 1, Pkg: "app2@1.0.0", Tag: map[string]string{"env": "coco"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&app.Info{Name: "app3", Version: "1.0.0", VersionId: 1, Pkg: "app3@1.0.0", Tag: map[string]string{"env": "coco"}}); err != nil {
		t.Error(err.Error())
	}
	{
		if err := Required(&app.Required{Name: "app", Version: "1.0.0", VersionId: 1},
			&app.Required{Name: "app2", Version: "1.0.0", VersionId: 1},
			&app.Required{Name: "app3", Version: "1.0.0", VersionId: 1},
			&app.Required{Name: "app3"},
			&app.Required{Name: "app4"}); err == nil {
			t.Error("test failed!")
		}
	}
	{
		_ = Installed("app@1.0.0")
		_ = Installed("app2@1.0.0")
		_ = Installed("app3@1.0.0")
		if err := Required(&app.Required{Name: "app", Version: "1.0.1", VersionId: 2},
			&app.Required{Name: "app2", Version: "1.0.0", VersionId: 1},
			&app.Required{Name: "app3", Version: "1.0.0", VersionId: 1}); err == nil {
			t.Error("test failed!")
		}
	}
	{
		_ = Installed("app@1.0.0")
		_ = Installed("app@1.0.1")
		_ = Installed("app2@1.0.0")
		_ = Installed("app3@1.0.0")
		if err := Required(&app.Required{Name: "app", Version: "1.0.1", VersionId: 2},
			&app.Required{Name: "app2", Version: "1.0.1", VersionId: 1},
			&app.Required{Name: "app3", Version: "1.0.0", VersionId: 1}); err != nil {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_Delete(t *testing.T) {
	UseFileManager(testPath)
	expect := `[]`
	if err := Add(&app.Info{Pkg: "app@1.0.0"}); err != nil {
		t.Error(err.Error())
	}
	if err := Delete("app@1.0.0"); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_DeleteTag(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env0": "", "env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if err := DeleteTag("app@1.0.0", "env0"); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_FindAll(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindAll(); err != nil {
		t.Error("test failed!")
	} else if apps != nil {
		if marshal, mErr := json.Marshal(apps); mErr != nil {
			t.Error("test failed!")
		} else if string(marshal) != expect {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_FindByPkg(t *testing.T) {
	UseFileManager(testPath)
	expect := `{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindByPkg("app@1.0.0"); err != nil {
		t.Error("test failed!")
	} else if apps != nil {
		if marshal, mErr := json.Marshal(apps); mErr != nil {
			t.Error("test failed!")
		} else if string(marshal) != expect {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_FindsByType(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindsByType(app.Other); err != nil {
		t.Error("test failed!")
	} else if apps != nil {
		if marshal, mErr := json.Marshal(apps); mErr != nil {
			t.Error("test failed!")
		} else if string(marshal) != expect {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_FindsByName(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindsByName("app"); err != nil {
		t.Error("test failed!")
	} else if apps != nil {
		if marshal, mErr := json.Marshal(apps); mErr != nil {
			t.Error("test failed!")
		} else if string(marshal) != expect {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_FindsByTag(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindsByTag("env"); err != nil {
		t.Error("test failed!")
	} else if apps != nil {
		if marshal, mErr := json.Marshal(apps); mErr != nil {
			t.Error("test failed!")
		} else if string(marshal) != expect {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_Installed(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if err := Installed("app@1.0.0"); err != nil {
		t.Error("test failed!")
	}
	if err := Uninstalled("app@1.0.0"); err != nil {
		t.Error("test failed!")
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}

func Test_fileManager_Uninstalled(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","VersionId":0,"Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":true}]`
	if err := Add(&app.Info{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if err := Installed("app@1.0.0"); err != nil {
		t.Error("test failed!")
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}
