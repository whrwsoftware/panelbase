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
	"os"
	"reflect"
	"testing"
)

var testPath = "/path/to/app.json"

func TestFileManager(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want Manager
	}{
		{"test", args{testPath}, FileManager(testPath)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileManager(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseFileManager(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{testPath}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseFileManager(tt.args.file)
			if mgr == nil {
				t.Error("test failed!")
			}
			mgr = nil
		})
	}
}

func Test_fileManager_Add(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"","Version":"","Pkg":"","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":null,"Installed":false}]`
	if err := Add(&App{}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0"}); err != nil {
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

func Test_fileManager_CheckInstalled(t *testing.T) {
	UseFileManager(testPath)
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": "hello"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&App{Name: "app", Version: "1.0.1", Pkg: "app@1.0.1", Tag: map[string]string{"env": "world"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&App{Name: "app2", Version: "1.0.0", Pkg: "app2@1.0.0", Tag: map[string]string{"env": "coco"}}); err != nil {
		t.Error(err.Error())
	}
	if err := Add(&App{Name: "app3", Version: "1.0.0", Pkg: "app3@1.0.0", Tag: map[string]string{"env": "coco"}}); err != nil {
		t.Error(err.Error())
	}
	{
		if err := Check(RequiredApp{"app", []string{"1.0.0"}},
			RequiredApp{"app2", []string{"1.0.0"}},
			RequiredApp{"app3", nil},
			RequiredApp{"app4", nil}); err == nil {
			t.Error("test failed!")
		}
	}
	{
		_ = Installed("app@1.0.0")
		_ = Installed("app2@1.0.0")
		_ = Installed("app3@1.0.0")
		if err := Check(RequiredApp{"app", []string{"1.0.1"}},
			RequiredApp{"app2", []string{"1.0.0"}},
			RequiredApp{"app3", nil}); err == nil {
			t.Error("test failed!")
		}
	}
	{
		_ = Installed("app@1.0.0")
		_ = Installed("app@1.0.1")
		_ = Installed("app2@1.0.0")
		_ = Installed("app3@1.0.0")
		if err := Check(RequiredApp{"app", []string{"1.0.1"}},
			RequiredApp{"app2", []string{"1.0.0"}},
			RequiredApp{"app3", nil}); err != nil {
			t.Error("test failed!")
		}
	}
}

func Test_fileManager_Delete(t *testing.T) {
	UseFileManager(testPath)
	expect := `[]`
	if err := Add(&App{Pkg: "app@1.0.0"}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env0": "", "env": ""}}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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
	expect := `{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
		t.Error(err.Error())
	}
	if apps, err := FindsByType(Other); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":false}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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
	expect := `[{"AppName":"","Name":"app","Version":"1.0.0","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":{"env":""},"Installed":true}]`
	if err := Add(&App{Name: "app", Version: "1.0.0", Pkg: "app@1.0.0", Tag: map[string]string{"env": ""}}); err != nil {
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

func Test_fileManager_Update(t *testing.T) {
	UseFileManager(testPath)
	expect := `[{"AppName":"","Name":"app","Version":"","Pkg":"app@1.0.0","Icon":"","Provider":"","Description":"","Workdir":"","Type":0,"Tag":null,"Installed":false}]`
	if err := Add(&App{Pkg: "app@1.0.0"}); err != nil {
		t.Error(err.Error())
	}
	if err := Update(&App{Pkg: "app@1.0.0", Name: "app"}); err != nil {
		t.Error(err.Error())
	}
	if bytes, err := os.ReadFile(testPath); err != nil {
		t.Error(err.Error())
	} else if string(bytes) != expect {
		t.Error("test failed!")
	}
}
