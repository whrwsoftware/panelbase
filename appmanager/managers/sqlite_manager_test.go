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
	"github.com/whrwsoftware/panelbase/appmanager"
	"reflect"
	"testing"
)

var dbFile = "test.db"

func TestManager(t *testing.T) {
	type args struct {
		dbFile string
	}
	tests := []struct {
		name string
		args args
		want *sqliteManager
	}{
		{"test", args{dbFile}, SqliteManager(dbFile)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqliteManager(tt.args.dbFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_Add(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appInfo *appmanager.AppInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{&appmanager.AppInfo{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Add(tt.args.appInfo); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_CreateVersion(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		version *appmanager.AppVersion
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{&appmanager.AppVersion{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.CreateVersion(tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("CreateVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_Delete(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Delete(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_DeleteVersion(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName string
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{"", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.DeleteVersion(tt.args.appName, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("DeleteVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_FindAll(t *testing.T) {
	type fields struct {
		DbFile string
	}
	tests := []struct {
		name         string
		fields       fields
		wantAppInfos []*appmanager.AppInfo
		wantErr      bool
	}{
		{"test", fields{dbFile}, []*appmanager.AppInfo{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppInfos, err := m.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppInfos, tt.wantAppInfos) {
				t.Errorf("FindAll() gotAppInfos = %v, want %v", gotAppInfos, tt.wantAppInfos)
			}
		})
	}
}

func Test_manager_FindByName(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAppInfo *appmanager.AppInfo
		wantErr     bool
	}{
		{"test", fields{dbFile}, args{""}, &appmanager.AppInfo{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppInfo, err := m.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppInfo, tt.wantAppInfo) {
				t.Errorf("FindByName() gotAppInfo = %v, want %v", gotAppInfo, tt.wantAppInfo)
			}
		})
	}
}

func Test_manager_FindByType(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appType byte
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantAppInfos []*appmanager.AppInfo
		wantErr      bool
	}{
		{"test", fields{dbFile}, args{0}, []*appmanager.AppInfo{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppInfos, err := m.FindByType(tt.args.appType)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppInfos, tt.wantAppInfos) {
				t.Errorf("FindByType() gotAppInfos = %v, want %v", gotAppInfos, tt.wantAppInfos)
			}
		})
	}
}

func Test_manager_GetAllVersion(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantAppVersions []*appmanager.AppVersion
		wantErr         bool
	}{
		{"test", fields{dbFile}, args{""}, []*appmanager.AppVersion{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppVersions, err := m.GetAllVersion(tt.args.appName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppVersions, tt.wantAppVersions) {
				t.Errorf("GetAllVersion() gotAppVersions = %v, want %v", gotAppVersions, tt.wantAppVersions)
			}
		})
	}
}

func Test_manager_GetVersion(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName string
		version string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantAppVersion *appmanager.AppVersion
		wantErr        bool
	}{
		{"test", fields{dbFile}, args{"", ""}, &appmanager.AppVersion{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppVersion, err := m.GetVersion(tt.args.appName, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppVersion, tt.wantAppVersion) {
				t.Errorf("GetVersion() gotAppVersion = %v, want %v", gotAppVersion, tt.wantAppVersion)
			}
		})
	}
}

func Test_manager_GetVersionById(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName   string
		versionId int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantAppVersion *appmanager.AppVersion
		wantErr        bool
	}{
		{"test", fields{dbFile}, args{"", 0}, &appmanager.AppVersion{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			gotAppVersion, err := m.GetVersionById(tt.args.appName, tt.args.versionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAppVersion, tt.wantAppVersion) {
				t.Errorf("GetVersionById() gotAppVersion = %v, want %v", gotAppVersion, tt.wantAppVersion)
			}
		})
	}
}

func Test_manager_Installed(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName string
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{"", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Installed(tt.args.appName, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("Installed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_Require(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		required []*appmanager.VersionRequired
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{[]*appmanager.VersionRequired{appmanager.NewRequired("", "", 0)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Require(tt.args.required...); (err != nil) != tt.wantErr {
				t.Errorf("Require() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_Uninstalled(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appName string
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{"", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Uninstalled(tt.args.appName, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("Uninstalled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_Update(t *testing.T) {
	type fields struct {
		DbFile string
	}
	type args struct {
		appInfo *appmanager.AppInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test", fields{dbFile}, args{&appmanager.AppInfo{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SqliteManager(tt.fields.DbFile)
			if err := m.Update(tt.args.appInfo); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
