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

package app

import "strings"

var manager Manager

func SetManager(mgr Manager) { manager = mgr }
func GetManager() Manager    { return manager }

type Manager interface {
	Add(info *Info) (err error)
	Update(info *Info) (err error)
	Delete(pkg string) (err error)
	AddTag(pkg string, tag string) (err error)
	DeleteTag(pkg string, tag string) (err error)
	Installed(pkg string) (err error)
	Uninstalled(pkg string) (err error)
	Required(required ...*Required) (err error)
	FindAll() (apps []*Info, err error)
	FindByPkg(pkg string) (info *Info, err error)
	FindsByType(typ Type) (infos []*Info, err error)
	FindsByName(name string) (infos []*Info, err error)
	FindsByTag(tag string) (infos []*Info, err error)
}

type Info struct {
	AppName     string            // Nginx 1.22.1
	Name        string            // nginx
	Version     string            // 1.22.1
	VersionId   int               // 100
	Pkg         string            // nginx@1.22.1
	Icon        string            // base64;
	Provider    string            // 官方
	Description string            // 轻量级，占有内存少，并发能力强
	Workdir     string            // /duckcp/apps/$Name/$Version
	Type        Type              // 分类
	Tag         map[string]string // Tag

	Installed bool
}

type Required struct {
	Name      string
	Version   string
	VersionId int
}

func NewRequired(name string, version string, versionId int) *Required {
	return &Required{Name: name, Version: version, VersionId: versionId}
}

type Type uint

const (
	Other       Type = iota // 其他
	Environment             // 运行环境
	SystemTool              // 系统工具
	Plugin                  // 插件
	ThirdParty              // 第三方
	Suite                   // 套件
)

type CheckError struct {
	NotExistApp   []string
	NotInstallApp []string
}

func (c *CheckError) Error() string {
	errText := "check error: "
	if ss := c.NotExistApp; ss != nil && len(ss) > 0 {
		errText += "NotExistApp(" + strings.Join(ss, ",") + ") "
	}
	if ss := c.NotInstallApp; ss != nil && len(ss) > 0 {
		errText += "NotInstallApp(" + strings.Join(ss, ",") + ")"
	}
	return errText
}
