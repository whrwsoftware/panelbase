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
	"errors"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"github.com/go-the-way/streams"
	"github.com/jmoiron/sqlx"
	"github.com/whrwsoftware/panelbase/appmanager"
)

func (s *sqliteManager) getDB() (db *sqlx.DB, err error) { return sqlx.Open("sqlite", s.DbFile) }

type sqliteManager struct{ DbFile string }

func SqliteManager(dbFile string) *sqliteManager {
	m := &sqliteManager{DbFile: dbFile}
	if db, err := m.getDB(); err != nil {
		fmt.Println(err)
	} else if db != nil {
		defer func() { _ = db.Close() }()
		if _, err = db.Exec(sqlFS); err != nil {
			fmt.Println(err)
		}
	}
	return m
}

func (s *sqliteManager) Add(appInfo *appmanager.AppInfo) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.NamedExec(`insert into app_info(name, icon, provider, workdir, app_type) values (:name, :icon, :provider, :workdir, :app_type)`, appInfo)
	}
	return
}

func (s *sqliteManager) Update(appInfo *appmanager.AppInfo) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.NamedExec(`update app_info set icon=:icon, provider=:provider, workdir=:workdir, app_type=:app_type where name=:name`, appInfo)
	}
	return
}

func (s *sqliteManager) Delete(name string) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.Exec(`delete from app_info where name=$1`, name)
		_, err = db.Exec(`delete from app_version where app_name=$1`, name)
	}
	return
}

func (s *sqliteManager) FindAll() (appInfos []*appmanager.AppInfo, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		err = db.Select(&appInfos, `select * from app_info`)
	}
	return
}

func (s *sqliteManager) FindByType(appType byte) (appInfos []*appmanager.AppInfo, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		err = db.Select(&appInfos, `select * from app_info where app_type=$1`, appType)
	}
	return
}

func (s *sqliteManager) FindByName(name string) (appInfo *appmanager.AppInfo, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		var appInfos []*appmanager.AppInfo
		err = db.Select(&appInfos, `select * from app_info where name=$1`, name)
		if len(appInfos) > 0 {
			appInfo = appInfos[0]
		}
	}
	return
}

func (s *sqliteManager) CreateVersion(version *appmanager.AppVersion) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		var cc int
		if err = db.Get(&cc, "select count(*) from app_info where name=$1", version.AppName); err != nil {
			return
		}
		if cc <= 0 {
			return errors.New("app is not exists")
		}
		_, err = db.NamedExec(`insert into app_version(app_name, version, version_id, installed, log_path, log_file) values (:app_name, :version, :version_id, :installed, :log_path, :log_file)`, version)
	}
	return
}

func (s *sqliteManager) GetAllVersion(appName string) (appVersions []*appmanager.AppVersion, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		err = db.Select(&appVersions, `select * from app_version where app_name=$1`, appName)
	}
	return
}

func (s *sqliteManager) GetVersion(appName, version string) (appVersion *appmanager.AppVersion, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		var appVersions []*appmanager.AppVersion
		err = db.Select(&appVersions, `select * from app_version where app_name=$1 and version=$2`, appName, version)
		if len(appVersions) > 0 {
			appVersion = appVersions[0]
		}
	}
	return
}

func (s *sqliteManager) GetVersionById(appName string, versionId int) (appVersion *appmanager.AppVersion, err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		var appVersions []*appmanager.AppVersion
		err = db.Select(&appVersions, `select * from app_version where app_name=$1 and version_id=$2`, appName, versionId)
		if len(appVersions) > 0 {
			appVersion = appVersions[0]
		}
	}
	return
}

func (s *sqliteManager) DeleteVersion(appName string, version string) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.Exec(`delete from app_version where app_name=$1 and version=$2`, appName, version)
	}
	return
}

func (s *sqliteManager) Installed(appName, version string) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.Exec(`update app_version set installed=1 where app_name=$1 and version=$2`, appName, version)
	}
	return
}

func (s *sqliteManager) Uninstalled(appName, version string) (err error) {
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		_, err = db.Exec(`update app_version set installed=0 where app_name=$1 and version=$2`, appName, version)
	}
	return
}

func (s *sqliteManager) RequireVersion(versionRequired ...*appmanager.VersionRequired) (err error) {
	if len(versionRequired) <= 0 {
		return nil
	}
	var db *sqlx.DB
	if db, err = s.getDB(); err != nil {
		return
	} else if db != nil {
		defer func() { _ = db.Close() }()
		var appInfos []*appmanager.AppInfo
		if err = db.Select(&appInfos, `select * from app_info`); err != nil {
			return
		}
		var appVersions []*appmanager.AppVersion
		if err = db.Select(&appVersions, `select * from app_version`); err != nil {
			return
		}
		var appInfoMap = streams.ToMap[*appmanager.AppInfo, string, *appmanager.AppInfo](appInfos,
			func(t *appmanager.AppInfo) (string, *appmanager.AppInfo) {
				return t.Name, t
			})
		var appVersionMap = streams.GroupBy[*appmanager.AppVersion, string, *appmanager.AppVersion](appVersions,
			func(t *appmanager.AppVersion) (string, *appmanager.AppVersion) {
				return t.AppName, t
			})

		var (
			notExistApp   []string
			notInstallApp []string
		)

		for _, rd := range versionRequired {
			if _, ok := appInfoMap[rd.AppName]; !ok {
				notExistApp = append(notExistApp, rd.AppName)
			} else {
				if appVer, okk := appVersionMap[rd.AppName]; !okk {
					notInstallApp = append(notInstallApp, rd.AppName+"@"+rd.Version)
				} else if streams.FilterThenCount[*appmanager.AppVersion](appVer, func(t *appmanager.AppVersion) bool {
					return t.VersionId >= rd.VersionId && t.Installed == 1
				}) <= 0 {
					notInstallApp = append(notInstallApp, rd.AppName+"@"+rd.Version)
				}
			}
		}

		if len(notExistApp) > 0 || len(notInstallApp) > 0 {
			err = &appmanager.RequiredError{NotExistApp: notInstallApp, NotInstallApp: notInstallApp}
		}

	}
	return
}
