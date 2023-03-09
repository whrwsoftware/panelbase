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

package oss

import (
	"os"
	"testing"
)

func Test_getOSId(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name          string
		args          args
		wantId        string
		wantVersionId string
	}{
		{"centos7", args{centos7Release.release}, centos7Release.id, centos7Release.versionId},
		{"centos8", args{centos8Release.release}, centos8Release.id, centos8Release.versionId},
		{"ubuntu", args{ubuntuRelease.release}, ubuntuRelease.id, ubuntuRelease.versionId},
		{"debian", args{debianRelease.release}, debianRelease.id, debianRelease.versionId},
		{"arch", args{archRelease.release}, archRelease.id, archRelease.versionId},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotVersionId := getOSId(tt.args.str)
			if gotId != tt.wantId {
				t.Errorf("getOSId() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotVersionId != tt.wantVersionId {
				t.Errorf("getOSId() gotVersionId = %v, want %v", gotVersionId, tt.wantVersionId)
			}
		})
	}
}

func TestCurrentOS(t *testing.T) {
	_ = os.MkdirAll("/etc", 0600)
	defer func() { _ = os.Remove("/etc") }()
	var (
		createOsReleaseFile = func(text string) { _ = os.WriteFile(osReleasePath, []byte(text), 0600) }
		removeOsReleaseFile = func() { _ = os.Remove(osReleasePath) }
	)
	tests := []struct {
		name      string
		wantOs    OS
		preFunc   func()
		deferFunc func()
	}{
		{"unknown", Unknown, func() {}, func() {}},
		{"centos7", CentOS7, func() { createOsReleaseFile(centos7Release.release) }, func() { removeOsReleaseFile() }},
		{"centos8", CentOS8, func() { createOsReleaseFile(centos8Release.release) }, func() { removeOsReleaseFile() }},
		{"ubuntu", Ubuntu, func() { createOsReleaseFile(ubuntuRelease.release) }, func() { removeOsReleaseFile() }},
		{"debian", Debian, func() { createOsReleaseFile(debianRelease.release) }, func() { removeOsReleaseFile() }},
		{"arch", Arch, func() { createOsReleaseFile(archRelease.release) }, func() { removeOsReleaseFile() }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.deferFunc()
			tt.preFunc()
			if gotOs := CurrentOS(); gotOs != tt.wantOs {
				t.Errorf("CurrentOS() = %v, want %v", gotOs, tt.wantOs)
			}
		})
	}
}
