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

import "strings"

type VersionManager interface {
	GetAllVersion(appName string) (appVersions []*AppVersion, err error)
	GetVersion(appName, version string) (appVersion *AppVersion, err error)
	GetVersionById(appName string, versionId int) (appVersion *AppVersion, err error)
	CreateVersion(version *AppVersion) (err error)
	DeleteVersion(appName string, version string) (err error)

	RequireVersion(versionRequired ...*VersionRequired) (err error)
}

type AppVersion struct {
	AppName   string `db:"app_name"`
	Version   string
	VersionId int `db:"version_id"`
	Installed int
	LogPath   string `db:"log_path"` // logs
	LogFile   string `db:"log_file"` // $Name.log
}

func NewAppVersion(appName string, version string, versionId int, installed int, logPath string, logFile string) *AppVersion {
	return &AppVersion{AppName: appName, Version: version, VersionId: versionId, Installed: installed, LogPath: logPath, LogFile: logFile}
}

type VersionRequired struct {
	AppName   string
	Version   string
	VersionId int
}

func NewVersionRequired(appName string, version string, versionId int) *VersionRequired {
	return &VersionRequired{AppName: appName, Version: version, VersionId: versionId}
}

type RequiredError struct {
	NotExistApp   []string
	NotInstallApp []string
}

func (c *RequiredError) Error() string {
	errText := "required error: "
	if ss := c.NotExistApp; ss != nil && len(ss) > 0 {
		errText += "NotExistApp(" + strings.Join(ss, ",") + ") "
	}
	if ss := c.NotInstallApp; ss != nil && len(ss) > 0 {
		errText += "NotInstallApp(" + strings.Join(ss, ",") + ")"
	}
	return errText
}
